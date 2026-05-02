# PCG/UPF Platform — Design & Observability Guide

## Quick Start
| Role                  | Start Here                                         |
|-----------------------|----------------------------------------------------|
| **SRE**               | Section 3.2 (Alerts) → 3.4 (Runbooks) → Appendix  |
| **Developer**         | Section 1 → Section 2 → Appendix                   |
| **Platform Engineer** | Section 2 → Section 3.5 (Thresholds) → Appendix    |

**Prerequisites:** Kubernetes basics, Prometheus/Grafana familiarity, cluster monitoring access.

## SRE Quick Reference

### Day 1
- [ ] Get access to Grafana, Prometheus, Elasticsearch, kubectl
- [ ] Import dashboards (Section 3.3), configure alert routes (Section 3.2)

### On-Call Priorities
| Priority | Watch                                                | Why                    |
|----------|------------------------------------------------------|------------------------|
| P0       | `non_mtls_*_conn_count > 0`                          | Unencrypted traffic    |
| P0       | `redis_cluster_state != 1`                           | Data layer down        |
| P0       | `elasticsearch_cluster_health_status{color="red"}`   | Log pipeline broken    |
| P1       | gRPC error rate > 5%                                 | Service degradation    |
| P1       | `*_max_attempts_exceeded_total > 0`                  | ID exhaustion          |
| P2       | Redis memory > 80%, ES heap > 75%                    | Approaching saturation |

### Incident Flow
1. Dashboard 1 (Platform Overview) → identify failing component
2. Golden Signals (Section 3.1) → confirm
3. Runbook (Section 3.4) → fix

## Platform Engineer Quick Reference

### Day 1
- [ ] Read Section 1 (Design Decisions) and Section 2 (Architecture)
- [ ] Review ID management (Appendix Section 4) and thresholds (Section 3.5)

### Tuning & Capacity
| Component     | Tune                                         | Capacity Threshold          |
|---------------|----------------------------------------------|-----------------------------|
| Redis         | `maxmemory`, defrag, slot balance            | Scale at 70% memory         |
| Kafka         | Partitions, consumer groups, replication     | Rebalance when queue grows  |
| Envoy         | Circuit breakers, retry budgets, pool sizes  | —                           |
| Elasticsearch | Heap, shard count, ILM                       | Add nodes at 30% disk free  |
| MinIO         | —                                            | Expand at 20% free          |
| ID Generators | Range sizes (SEID/TEID/Port)                 | —                           |

---

The PCG/UPF platform is a cloud-native 5G user plane deployed on Kubernetes. It processes mobile data traffic at scale using Envoy, gRPC, Redis, Kafka, etcd, MinIO, Elasticsearch, Prometheus/Grafana, and TLS/mTLS/DTLS. See Section 1 (Design Decisions) for why each technology was chosen.

---

## Glossary

| Acronym  | Meaning                                          |
|----------|--------------------------------------------------|
| PCG      | Packet Core Gateway                              |
| UPF      | User Plane Function (3GPP 5G)                    |
| CRE      | Core Runtime Environment                         |
| TMRA     | Timer & Route Management Application             |
| TMPD     | Timer Processing Daemon                          |
| PCUPDP   | PCG User Plane Data Path                         |
| PCUPPEP  | PCG User Plane Policy Enforcement Point          |
| PCUPUPCP | PCG User Plane Control Path                      |
| PCUPLEP  | PCG L2TP Endpoint Processing                     |
| DATATW   | Data Transport Worker                            |
| KVDB     | Key-Value Database                               |
| SEID     | Session Endpoint Identifier (PFCP)               |
| TEID     | Tunnel Endpoint Identifier (GTP)                 |
| DTLS     | Datagram Transport Layer Security                |
| mTLS     | Mutual TLS                                       |
| SIP TLS  | Service Identity Provider TLS (cert management)  |
| DDC      | Data Distribution & Collection                   |
| RBAC     | Role-Based Access Control                        |
| ILM      | Index Lifecycle Management (Elasticsearch)       |

---

## 1. Design Decisions
### Why These Technologies?
- **Envoy over NGINX/HAProxy** — Native gRPC support, built-in mTLS, dynamic config via xDS API, fine-grained observability out of the box. Essential for a service mesh architecture.
- **gRPC over REST** — Binary protocol (protobuf) is faster and smaller than JSON. HTTP/2 multiplexing avoids head-of-line blocking. Strongly typed contracts prevent integration bugs.
- **Redis over Memcached/PostgreSQL** — Sub-millisecond latency for session lookups. Cluster mode gives horizontal scaling. Built-in replication for HA. Supports atomic ID generation (INCR/ranges).
- **Kafka over RabbitMQ** — Higher throughput for event streaming. Durable message log allows replay. Partition-based parallelism scales with consumers. Better fit for event-driven architectures.
- **etcd over Consul/ZooKeeper** — Native Kubernetes integration (K8s itself uses etcd). Strong consistency via Raft. Watch API for real-time config changes.
- **MinIO over cloud-native S3** — On-prem S3-compatible storage. Erasure coding provides data durability without cloud dependency. Runs inside the same Kubernetes cluster.
- **Elasticsearch over Loki/Splunk** — Full-text search across logs. Scales horizontally with sharding. Rich query language (KQL/Lucene). Proven at telecom log volumes.
- **Prometheus + Grafana over Datadog/commercial APM** — Open-source, no licensing cost. Pull-based model works well in Kubernetes. PromQL is the industry standard for metric queries. Grafana provides flexible dashboarding.
- **mTLS everywhere** — Zero-trust networking. Every service authenticates both sides of every connection. SIP TLS automates cert rotation so it's operationally invisible.
- **DTLS on user plane** — User-plane traffic is UDP (GTP-U). TLS doesn't work over UDP, so DTLS provides equivalent encryption without switching to TCP.

### Networking
- **mTLS enforced** between all services — monitored via `non_mtls_*_conn_count` (must be 0)
- **DTLS** on user plane (`pcupdp_dtls_*`) for encrypted data transport
- **Envoy** as L7 proxy with circuit breakers, retries, RBAC

### Data Layer
- **Redis Cluster** for session state, ID generation, caching
- **KVDB Replicator** for geo-redundant DB sync (`pckvdbrdr_bulk_sync_*`)
- **DB Proxy** pattern — all DB access goes through proxy with latency tracking

### Messaging
- **Kafka** for async event processing between components
- Per-component producers/consumers with independent error tracking

### Storage
- **MinIO** with erasure coding for object storage
- **Elasticsearch** for centralized logging via DDC (`ddc_log_elasticsearch_*`)

### ID Management
- Distributed ID generators for SEID, TEID, Port IDs, Call IDs
- Range-based allocation with pre-allocation to avoid hot paths

---

## 2. Architecture Overview

```text
┌─────────────────────────────────────────────────────────┐
│                    Kubernetes Cluster                     │
│                                                          │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌─────────┐ │
│  │  Envoy   │→ │  gRPC    │→ │  App     │→ │  Redis  │ │
│  │  Proxy   │  │  Server  │  │  Logic   │  │  Cluster│ │
│  │ (ingress)│  │ (CRE)    │  │ (tmra,   │  │         │ │
│  └──────────┘  └──────────┘  │  pcupdp, │  └─────────┘ │
│                              │  pcuppep) │              │
│  ┌──────────┐               └────┬──────┘              │
│  │  Kafka   │←───────────────────┘                     │
│  │  Cluster │                                           │
│  └──────────┘  ┌──────────┐  ┌──────────┐              │
│                │  MinIO   │  │  Elastic  │              │
│  ┌──────────┐  │ (storage)│  │  search   │              │
│  │  etcd    │  └──────────┘  │  (logs)   │              │
│  │  Cluster │                └──────────┘              │
│  └──────────┘                                           │
│                                                          │
│  ┌─────────────────────────────────────────────────────┐ │
│  │  TLS/mTLS everywhere — SIP TLS, DTLS (user plane)  │ │
│  └─────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────┘
```

### Components
| Component       | Role                                           | Key Prefixes      |
|-----------------|------------------------------------------------|-------------------|
| Envoy           | Ingress proxy, TLS termination, routing, RBAC  | `envoy_*`         |
| CRE gRPC        | Core runtime gRPC services                     | `cre_grpc_*`      |
| TMRA            | Timer & route management                       | `tmra_*`          |
| TMPD            | Timer processing daemon                        | `tmpd_*`          |
| PCUPDP          | User plane data path (UPF)                     | `pcupdp_*`        |
| PCUPPEP         | User plane PEP (policy enforcement)            | `pcuppep_*`       |
| PCUPUPCP        | User plane control path                        | `pcupupcp_*`      |
| PCUPLEP         | L2TP endpoint processing                       | `pcuplep_*`       |
| DATATW          | Data transport worker                          | `datatw_*`        |
| KVDB Replicator | Cross-site DB replication                      | `pckvdbrdr_*`     |
| Redis           | Session/state store, caching                   | `redis_*`         |
| Kafka           | Event streaming, notifications                 | `kafka_*`         |
| etcd            | Config store, service discovery                | `etcd_*`          |
| MinIO           | Object storage                                 | `minio_*`         |
| Elasticsearch   | Log aggregation, search                        | `elasticsearch_*` |
| SIP TLS         | Certificate management                         | `sip_tls_*`       |

---

## 3. Observability

### 3.1 Golden Signals per Component

| Component | Latency | Traffic | Errors | Saturation |
|-----------|---------|---------|--------|------------|
| gRPC | `grpc_server_handled_total` by duration | `grpc_server_started_total` | `grpc_server_handled_total{code!="OK"}` | `cre_grpc_load_protection_thresholds_total` |
| Envoy | `envoy_cluster_upstream_rq_time_bucket` | `envoy_http_downstream_rq_total` | `envoy_http_downstream_rq_xx{code="5"}` | `envoy_cluster_upstream_cx_active` |
| Redis | `redis_commands_duration_seconds_total` | `redis_commands_processed_total` | `redis_rejected_connections_total` | `redis_memory_used_bytes / redis_memory_max_bytes` |
| Kafka | `kafka_network_RequestMetrics_*Percentile` | `tmra_kafka_published_messages_total` | `tmra_kafka_errors_total` | `kafka_server_Request_queue_size` |
| DB Proxy | `db_proxy_latency_seconds` | `db_proxy_command_requests_total` | `db_tracker_errors_total` | `db_proxy_hiredis_write_buffer_bytes` |
| Elasticsearch | `elasticsearch_indices_search_query_time_seconds` | `elasticsearch_indices_search_query_total` | `elasticsearch_cluster_health_status` | `elasticsearch_jvm_memory_used_bytes` |
| MinIO | `minio_s3_requests_ttfb_seconds_distribution` | `minio_s3_requests_total` | `minio_s3_requests_5xx_errors_total` | `minio_capacity_usable_free_total` |

### 3.2 Critical Alerts

#### TLS / Security
```yaml
- alert: CertExpiringIn7Days
  expr: TLS_Certificate_Expiration_Time_in_Seconds < 604800
  severity: critical

- alert: NonMtlsConnections
  expr: non_mtls_remote_conn_count + non_mtls_local_conn_count > 0
  severity: critical

- alert: TlsHandshakeFailures
  expr: rate(tls_handshake_exceeded[5m]) > 0
  severity: warning

- alert: DtlsConnectionFailures
  expr: rate(pcupdp_dtls_connection_failures_total[5m]) > 10
  severity: warning
```

#### gRPC
```yaml
- alert: GrpcHighErrorRate
  expr: rate(grpc_server_handled_total{grpc_code!="OK"}[5m]) / rate(grpc_server_handled_total[5m]) > 0.05
  severity: critical

- alert: GrpcLoadProtectionTriggered
  expr: rate(cre_grpc_load_protection_thresholds_total[5m]) > 0
  severity: warning
```

#### Redis
```yaml
- alert: RedisMemoryHigh
  expr: redis_memory_used_bytes / redis_memory_max_bytes > 0.9
  severity: critical

- alert: RedisFragmentation
  expr: redis_mem_fragmentation_ratio > 2.0
  severity: warning

- alert: RedisClusterFail
  expr: redis_cluster_state != 1
  severity: critical

- alert: RedisReplicationDown
  expr: redis_master_link_up == 0
  severity: critical
```

#### Kafka
```yaml
- alert: KafkaErrors
  expr: rate(tmra_kafka_errors_total[5m]) > 0
  severity: warning

- alert: KafkaUnderReplicated
  expr: kafka_server_ReplicaManager_Value{name="UnderReplicatedPartitions"} > 0
  for: 5m
  severity: critical
```

#### Database
```yaml
- alert: DbProxyLatencyHigh
  expr: histogram_quantile(0.99, rate(tmra_db_proxy_latency_seconds_bucket[5m])) > 0.05
  severity: warning

- alert: IdAllocationExhausted
  expr: rate(pcupdp_db_id_gen_port_allocator_port_id_range_allocate_max_attempts_exceeded_total[5m]) > 0
  severity: critical

- alert: KvdbReplicationStalled
  expr: pckvdbrdr_bulk_sync_percent < 100 and rate(pckvdbrdr_bulk_sync_percent[10m]) == 0
  severity: critical
```

#### Envoy
```yaml
- alert: Envoy5xxHigh
  expr: rate(envoy_http_downstream_rq_xx{envoy_response_code_class="5"}[5m]) / rate(envoy_http_downstream_rq_total[5m]) > 0.05
  severity: critical

- alert: EnvoyUpstreamUnhealthy
  expr: envoy_cluster_membership_total - envoy_cluster_membership_healthy > 0
  severity: warning

- alert: EnvoyConnectionOverflow
  expr: rate(envoy_cluster_upstream_cx_overflow[5m]) > 0
  severity: warning
```

#### MinIO
```yaml
- alert: MinioDriveOffline
  expr: minio_cluster_drive_offline_total > 0
  severity: critical

- alert: MinioCapacityLow
  expr: minio_capacity_usable_free_total / minio_capacity_usable_total < 0.1
  severity: critical
```

#### Elasticsearch
```yaml
- alert: ElasticsearchRed
  expr: elasticsearch_cluster_health_status{color="red"} == 1
  severity: critical

- alert: ElasticsearchHeapHigh
  expr: elasticsearch_jvm_memory_used_bytes / elasticsearch_jvm_memory_max_bytes > 0.85
  severity: warning

- alert: ElasticsearchUnassignedShards
  expr: elasticsearch_cluster_health_unassigned_shards > 0
  for: 5m
  severity: warning
```

#### Linux / Network
```yaml
- alert: TcpRetransmissionsHigh
  expr: rate(tmra_linux_proc_net_snmp_tcp_retranssegs_total[5m]) / rate(tmra_linux_proc_net_snmp_tcp_outsegs_total[5m]) > 0.05
  severity: warning

- alert: UdpBufferErrors
  expr: rate(tmra_linux_proc_net_snmp_udp_rcvbuferrors_total[5m]) > 100
  severity: warning
```

### 3.3 Key Dashboards

#### Dashboard 1: Platform Overview
| Panel              | Query                                                                                                        |
|--------------------|--------------------------------------------------------------------------------------------------------------|
| gRPC RPS           | `sum(rate(grpc_server_started_total[5m]))`                                                                   |
| gRPC Error %       | `sum(rate(grpc_server_handled_total{grpc_code!="OK"}[5m])) / sum(rate(grpc_server_handled_total[5m])) * 100` |
| Envoy 5xx          | `sum(rate(envoy_http_downstream_rq_xx{envoy_response_code_class="5"}[5m]))`                                  |
| Active Connections | `envoy_http_downstream_cx_active`                                                                            |
| Cert Expiry Days   | `min(TLS_Certificate_Expiration_Time_in_Seconds) / 86400`                                                    |
| Non-mTLS Count     | `non_mtls_remote_conn_count + non_mtls_local_conn_count`                                                     |

#### Dashboard 2: Data Layer
| Panel            | Query                                                                                                          |
|------------------|----------------------------------------------------------------------------------------------------------------|
| Redis Memory %   | `redis_memory_used_bytes / redis_memory_max_bytes * 100`                                                       |
| Redis Ops/s      | `rate(redis_commands_processed_total[5m])`                                                                     |
| Redis Hit Ratio  | `rate(redis_keyspace_hits_total[5m]) / (rate(redis_keyspace_hits_total[5m]) + rate(redis_keyspace_misses_total[5m]))` |
| DB Proxy p99     | `histogram_quantile(0.99, rate(tmra_db_proxy_latency_seconds_bucket[5m]))`                                     |
| KVDB Sync %      | `pckvdbrdr_bulk_sync_percent`                                                                                  |
| Kafka Produce Rate | `sum(rate(tmra_kafka_published_messages_total[5m]))`                                                         |

#### Dashboard 3: User Plane (UPF)
| Panel               | Query                                                                                              |
|---------------------|----------------------------------------------------------------------------------------------------|
| DTLS Active         | `pcupdp_dptls_active_connections`                                                                  |
| DTLS Failures       | `rate(pcupdp_dtls_connection_failures_total[5m])`                                                  |
| Packet TX/RX        | `rate(pcupdp_pktio_linux_transmitted_packets_total[5m])` / `rate(pcupdp_pktio_linux_received_packets_total[5m])` |
| Port ID Alloc Fails | `rate(pcupdp_db_id_gen_port_allocator_port_id_range_allocate_max_attempts_exceeded_total[5m])`     |
| TCP Retransmissions  | `rate(pcupdp_linux_proc_net_snmp_tcp_retranssegs_total[5m])`                                      |
| UDP Buffer Errors   | `rate(pcupdp_linux_proc_net_snmp_udp_rcvbuferrors_total[5m])`                                      |

#### Dashboard 4: Storage & Logs
| Panel                 | Query                                                                        |
|-----------------------|------------------------------------------------------------------------------|
| MinIO Health          | `minio_cluster_health_status`                                                |
| MinIO Capacity % Used | `(1 - minio_capacity_usable_free_total / minio_capacity_usable_total) * 100` |
| MinIO S3 Errors       | `rate(minio_s3_requests_5xx_errors_total[5m])`                               |
| ES Cluster Status     | `elasticsearch_cluster_health_status`                                        |
| ES Indexing Rate      | `rate(elasticsearch_indices_indexing_index_total[5m])`                        |
| ES JVM Heap %         | `elasticsearch_jvm_memory_used_bytes / elasticsearch_jvm_memory_max_bytes * 100` |

### 3.4 Troubleshooting Runbook

For detailed per-component troubleshooting steps with metrics and actions, see the Appendix:

| Issue                  | Go To                                                       |
|------------------------|-------------------------------------------------------------|
| gRPC failures          | Appendix Section 1 — gRPC Troubleshooting                   |
| TLS/cert issues        | Appendix Section 2 — TLS Troubleshooting                    |
| Redis problems         | Appendix Section 8 — Redis Troubleshooting                  |
| Kafka lag              | Appendix Section 5 — Kafka Troubleshooting                  |
| DB/Replication issues  | Appendix Section 4 — Database Troubleshooting               |
| Envoy routing failures | Appendix Section 6 — Envoy Troubleshooting                  |
| MinIO storage          | Appendix Section 7 — MinIO Troubleshooting                  |
| Elasticsearch          | Appendix Section 9 — Elasticsearch Troubleshooting          |
| Linux/Network          | Appendix Section 3 — Linux Process Metrics Troubleshooting  |

### 3.5 Alert Thresholds Summary

| Component     | Metric                         | Warning   | Critical     |
|---------------|--------------------------------|-----------|--------------|
| TLS           | Cert expiry                    | < 30 days | < 7 days     |
| TLS           | `non_mtls_*_conn_count`        | > 0       | > 0          |
| gRPC          | Error rate                     | > 1%      | > 5%         |
| Redis         | Memory usage                   | > 80%     | > 90%        |
| Redis         | Fragmentation ratio            | > 1.5     | > 2.0        |
| Elasticsearch | Cluster status                 | Yellow    | Red          |
| Elasticsearch | JVM heap                       | > 75%     | > 85%        |
| MinIO         | Drive offline                  | > 0       | > quorum     |
| Kafka         | Under-replicated partitions    | > 0       | > 0 for 5min |
| Envoy         | 5xx rate                       | > 1%      | > 5%         |
| Linux         | TCP retransmissions            | > 1%      | > 5%         |
| DB            | ID alloc max attempts exceeded | > 0       | > 0          |
| DB            | Proxy latency p99              | > 10ms    | > 50ms       |

---

# Appendix: Per-Metric Reference

## Appendix Summary

### Table A — Key Metrics by Component

| #   | Component     | Metric                                              | Purpose                    |
|-----|---------------|-----------------------------------------------------|----------------------------|
| 1   | gRPC          | `grpc_server_started_total`                         | Traffic volume baseline    |
| 1   | gRPC          | `grpc_server_handled_total`                         | Completed RPCs             |
| 1   | gRPC          | `cre_grpc_load_protection_thresholds_total`         | Overload detection         |
| 1   | gRPC          | `tmra_rs_client_mgr_grpc_client_ongoing_requests`   | In-flight requests         |
| 1   | gRPC          | `etcd_grpc_proxy_cache_hits_total`                  | Cache efficiency           |
| 2   | TLS/mTLS      | `TLS_Certificate_Expiration_Time_in_Seconds`        | Cert expiry countdown      |
| 2   | TLS/mTLS      | `tls_handshake_exceeded`                            | Failed handshakes          |
| 2   | TLS/mTLS      | `non_mtls_remote_conn_count`                        | Unencrypted remote conns   |
| 2   | TLS/mTLS      | `non_mtls_local_conn_count`                         | Unencrypted local conns    |
| 2   | TLS/mTLS      | `pcupdp_dtls_connection_failures_total`             | DTLS failures              |
| 3   | Linux         | `tcp_retranssegs_total`                             | TCP retransmissions        |
| 3   | Linux         | `tcp_currestab`                                     | Active TCP connections     |
| 3   | Linux         | `udp_rcvbuferrors_total`                            | UDP buffer overflows       |
| 3   | Linux         | `ip_indiscards_total`                               | IP layer drops             |
| 3   | Linux         | `icmp_indestunreachs_total`                         | Reachability failures      |
| 4   | Database      | `db_proxy_latency_seconds`                          | Proxy request latency      |
| 4   | Database      | `db_proxy_hiredis_write_buffer_bytes`               | Write backpressure         |
| 4   | Database      | `pckvdbrdr_bulk_sync_percent`                       | Replication progress       |
| 4   | Database      | `*_max_attempts_exceeded_total`                     | ID allocation failures     |
| 5   | Kafka         | `tmra_kafka_published_messages_total`               | Produce rate               |
| 5   | Kafka         | `tmra_kafka_consumed_messages_total`                | Consume rate               |
| 5   | Kafka         | `tmra_kafka_errors_total`                           | Kafka errors               |
| 5   | Kafka         | `kafka_server_ReplicaManager_Value`                 | Under-replicated parts     |
| 6   | Envoy         | `envoy_http_downstream_rq_xx`                       | Response code classes      |
| 6   | Envoy         | `envoy_cluster_upstream_rq_timeout`                 | Backend timeouts           |
| 6   | Envoy         | `envoy_cluster_membership_healthy`                  | Healthy endpoints          |
| 6   | Envoy         | `envoy_cluster_upstream_cx_overflow`                | Connection pool overflow   |
| 6   | Envoy         | `envoy_http_rbac_denied`                            | RBAC denials               |
| 7   | MinIO         | `minio_cluster_health_status`                       | Cluster health             |
| 7   | MinIO         | `minio_cluster_drive_offline_total`                 | Offline drives             |
| 7   | MinIO         | `minio_capacity_usable_free_total`                  | Free capacity              |
| 7   | MinIO         | `minio_s3_requests_5xx_errors_total`                | S3 server errors           |
| 7   | MinIO         | `minio_node_drive_latency_us`                       | Drive latency              |
| 8   | Redis         | `redis_memory_used_bytes`                           | Memory consumption         |
| 8   | Redis         | `redis_mem_fragmentation_ratio`                     | RAM fragmentation          |
| 8   | Redis         | `redis_cluster_state`                               | Cluster OK/FAIL            |
| 8   | Redis         | `redis_master_link_up`                              | Replication link           |
| 8   | Redis         | `redis_keyspace_hits_total`                         | Cache hits                 |
| 8   | Redis         | `redis_evicted_keys_total`                          | Eviction pressure          |
| 9   | Elasticsearch | `elasticsearch_cluster_health_status`               | Green/Yellow/Red           |
| 9   | Elasticsearch | `elasticsearch_jvm_memory_used_bytes`               | JVM heap usage             |
| 9   | Elasticsearch | `elasticsearch_cluster_health_unassigned_shards`    | Unassigned shards          |
| 9   | Elasticsearch | `elasticsearch_jvm_gc_collection_seconds_sum`       | GC pressure                |
| 9   | Elasticsearch | `elasticsearch_filesystem_data_free_bytes`          | Disk free space            |

### Table B — Critical Alerts by Component

| #   | Component     | Alert Condition                | Severity |
|-----|---------------|--------------------------------|----------|
| 1   | gRPC          | Error rate > 5%                | Critical |
| 1   | gRPC          | Load protection triggered      | Warning  |
| 2   | TLS/mTLS      | Cert expiry < 7 days           | Critical |
| 2   | TLS/mTLS      | Non-mTLS connections > 0       | Critical |
| 3   | Linux         | TCP retransmissions > 5%       | Critical |
| 3   | Linux         | UDP buffer errors > 100/s      | Warning  |
| 4   | Database      | DB proxy p99 > 50ms            | Critical |
| 4   | Database      | ID allocation exhausted        | Critical |
| 4   | Database      | KVDB replication stalled       | Critical |
| 5   | Kafka         | Under-replicated > 0 for 5min  | Critical |
| 5   | Kafka         | Kafka errors > 0               | Warning  |
| 6   | Envoy         | 5xx rate > 5%                  | Critical |
| 6   | Envoy         | Upstream unhealthy             | Warning  |
| 6   | Envoy         | Connection overflow            | Warning  |
| 7   | MinIO         | Drive offline > 0              | Critical |
| 7   | MinIO         | Usable capacity < 10% free     | Critical |
| 8   | Redis         | Memory > 90%                   | Critical |
| 8   | Redis         | Cluster FAIL                   | Critical |
| 8   | Redis         | Replication down               | Critical |
| 8   | Redis         | Fragmentation > 2.0            | Warning  |
| 9   | Elasticsearch | Cluster RED                    | Critical |
| 9   | Elasticsearch | JVM heap > 85%                 | Warning  |
| 9   | Elasticsearch | Unassigned shards > 0          | Warning  |

### Table C — Key PromQL Queries

| #   | Component     | Description            | Query                                                                        |
|-----|---------------|------------------------|------------------------------------------------------------------------------|
| 1   | gRPC          | Error rate             | `rate(grpc_server_handled_total{grpc_code!="OK"}[5m])`                       |
|     |               |                        | `/ rate(grpc_server_handled_total[5m])`                                      |
| 2   | TLS/mTLS      | Days to cert expiry    | `TLS_Certificate_Expiration_Time_in_Seconds / 86400`                         |
| 2   | TLS/mTLS      | Non-mTLS count         | `non_mtls_remote_conn_count + non_mtls_local_conn_count`                     |
| 3   | Linux         | TCP retransmit rate    | `rate(tmra_linux_proc_net_snmp_tcp_retranssegs_total[5m])`                   |
| 4   | Database      | Proxy latency p99      | `histogram_quantile(0.99,`                                                   |
|     |               |                        | `rate(tmra_db_proxy_latency_seconds_bucket[5m]))`                            |
| 5   | Kafka         | Consumer lag           | `sum(tmra_kafka_published_messages_total)`                                   |
|     |               |                        | `- sum(tmra_kafka_consumed_messages_total)`                                  |
| 6   | Envoy         | 5xx error rate         | `rate(envoy_http_downstream_rq_xx{code_class="5"}[5m])`                     |
|     |               |                        | `/ rate(envoy_http_downstream_rq_total[5m])`                                 |
| 7   | MinIO         | Remaining capacity %   | `minio_capacity_usable_free_total`                                           |
|     |               |                        | `/ minio_capacity_usable_total * 100`                                        |
| 8   | Redis         | Memory usage %         | `redis_memory_used_bytes / redis_memory_max_bytes * 100`                     |
| 8   | Redis         | Cache hit ratio        | `redis_keyspace_hits_total`                                                  |
|     |               |                        | `/ (redis_keyspace_hits_total + redis_keyspace_misses_total)`                |
| 9   | Elasticsearch | JVM heap %             | `elasticsearch_jvm_memory_used_bytes`                                        |
|     |               |                        | `/ elasticsearch_jvm_memory_max_bytes * 100`                                 |
| 9   | Elasticsearch | Indexing rate          | `rate(elasticsearch_indices_indexing_index_total[5m])`                        |

### Table D — Troubleshooting Quick Reference

| #   | Component     | Symptom              | Action                                             |
|-----|---------------|----------------------|----------------------------------------------------|
| 1   | gRPC          | High error rate      | Check logs for UNAVAILABLE/DEADLINE_EXCEEDED       |
| 1   | gRPC          | Connection churn     | Check DNS, load balancer, or cert issues           |
| 1   | gRPC          | Cache misses         | Increase etcd proxy cache size                     |
| 2   | TLS/mTLS      | Cert expiring        | Rotate cert, check auto-renewal                    |
| 2   | TLS/mTLS      | Handshake failures   | Check cert chain, cipher mismatch, clock skew      |
| 2   | TLS/mTLS      | Non-mTLS traffic     | Enforce mTLS policy, check sidecar injection       |
| 3   | Linux         | High retransmissions | Check network congestion, MTU, packet loss         |
| 3   | Linux         | UDP buffer errors    | Increase `net.core.rmem_max`, scale receivers      |
| 3   | Linux         | IP discards          | Check memory pressure, socket buffer sizes         |
| 4   | Database      | High latency         | Check Redis cluster health, network                |
| 4   | Database      | Write buffer growing | Redis is slow — check memory, eviction             |
| 4   | Database      | Replication stuck    | Check source DB, cross-site network                |
| 4   | Database      | ID exhaustion        | Expand ID range, check for leaks                   |
| 5   | Kafka         | Consumer lag         | Scale consumers, check processing time             |
| 5   | Kafka         | Produce errors       | Check broker health, disk space, replication       |
| 5   | Kafka         | Under-replicated     | Check broker connectivity, disk health             |
| 6   | Envoy         | 5xx spike            | Check backend health, circuit breakers             |
| 6   | Envoy         | Connection overflow  | Increase circuit breaker limits                    |
| 6   | Envoy         | TLS errors           | Check cert chain, expiry, SAN mismatch             |
| 6   | Envoy         | RBAC denials         | Review RBAC policy, check source identity          |
| 7   | MinIO         | 5xx errors           | Check drive health, quorum, memory                 |
| 7   | MinIO         | Slow TTFB            | Check drive latency, network, CPU                  |
| 7   | MinIO         | Drive offline        | Replace drive, check erasure set health            |
| 7   | MinIO         | Capacity low         | Add drives, enable ILM policies                    |
| 8   | Redis         | High memory          | Increase maxmemory, check key TTLs                 |
| 8   | Redis         | Fragmentation > 1.5  | Restart Redis or enable active defrag              |
| 8   | Redis         | Replication lag      | Check network, slave disk I/O                      |
| 8   | Redis         | Cluster FAIL         | Check failed nodes, rebalance slots                |
| 9   | Elasticsearch | Red cluster          | Check failed nodes, disk space, shard allocation   |
| 9   | Elasticsearch | Slow search          | Optimize queries, increase heap, add nodes         |
| 9   | Elasticsearch | High GC              | Reduce heap pressure, check bulk indexing          |
| 9   | Elasticsearch | Disk full            | Add nodes, delete old indices, enable ILM          |

---

## 1. gRPC Metrics

> **gRPC** is the primary RPC framework used for inter-service communication on the platform. All core services (CRE, TMRA, etcd proxy) expose gRPC endpoints. Metrics here tell you whether requests are flowing, how fast they complete, and when the system is under load protection.

### Key Counters
| Metric                          | What It Tells You                                              |
|---------------------------------|----------------------------------------------------------------|
| `grpc_server_started_total`     | Total RPCs started — baseline for traffic volume               |
| `grpc_server_handled_total`     | Total RPCs completed — compare with started to find in-flight  |
| `grpc_server_msg_received_total`| Messages received by server                                    |
| `grpc_server_msg_sent_total`    | Messages sent by server                                        |

### CRE gRPC (Custom Runtime Environment)
| Metric                                            | Purpose                                        |
|---------------------------------------------------|------------------------------------------------|
| `cre_grpc_client_total`                           | Active gRPC clients                            |
| `cre_grpc_consumer_total` / `cre_grpc_producer_total` | Consumer/producer counts                  |
| `cre_grpc_db_conn_pl_total`                       | DB connection pool via gRPC                    |
| `cre_grpc_timer_pool_timer_operations_total`      | Timer pool operations                          |
| `cre_grpc_load_protection_thresholds_total`       | Load protection triggers — spikes mean overload|

### TMRA gRPC Client
| Metric                                              | Purpose                            |
|-----------------------------------------------------|------------------------------------|
| `tmra_ra_grpc_client_sm_recreate_seconds`           | Client state machine recreate latency |
| `tmra_rs_client_mgr_grpc_client_requests_total`     | Outgoing requests                  |
| `tmra_rs_client_mgr_grpc_client_responses_total`    | Incoming responses                 |
| `tmra_rs_client_mgr_grpc_client_ongoing_requests`   | In-flight — high = bottleneck      |

### etcd gRPC Proxy
| Metric                                                    | Purpose              |
|-----------------------------------------------------------|----------------------|
| `etcd_grpc_proxy_cache_hits_total` / `cache_misses_total` | Cache hit ratio      |
| `etcd_grpc_proxy_watchers_coalescing_total`               | Watcher efficiency   |
| `etcd_network_client_grpc_sent_bytes_total`               | Network throughput   |

### Monitoring
```promql
# gRPC error rate
rate(grpc_server_handled_total{grpc_code!="OK"}[5m]) / rate(grpc_server_handled_total[5m])

# In-flight RPCs
grpc_server_started_total - grpc_server_handled_total

# CRE load protection triggers
rate(cre_grpc_load_protection_thresholds_total[5m])
```

### Troubleshooting
| Symptom          | Check                                    | Action                                          |
|------------------|------------------------------------------|--------------------------------------------------|
| High error rate  | `grpc_server_handled_total` by code      | Check logs for UNAVAILABLE/DEADLINE_EXCEEDED     |
| Slow responses   | `tmra_ra_grpc_client_sm_recreate_seconds`| Check downstream service health                  |
| Connection churn | `cre_grpc_client_total` trend            | Check DNS, load balancer, or cert issues         |
| Cache misses     | `etcd_grpc_proxy_cache_misses_total`     | Increase cache size or check key patterns        |

---

## 2. TLS / mTLS Metrics

> **TLS/mTLS** secures all service-to-service traffic. SIP TLS manages certificate lifecycle (issuance, rotation, expiry). DTLS secures user-plane UDP traffic. Any non-mTLS connection in production is a security incident.

### Certificate Health
| Metric                                                  | What It Tells You                          |
|---------------------------------------------------------|--------------------------------------------|
| `TLS_Certificate_Expiration_Time_in_Seconds`            | Time until cert expires — alert at <30 days|
| `sip_tls_internal_certificate_total`                    | Total internal certs managed               |
| `sip_tls_internal_user_ca_total`                        | User CA certificates                       |
| `ddc_tls_certificate_read_total` / `ddc_tls_ca_read_total` | Cert/CA read operations                |

### Handshake & Connection
| Metric                                                    | What It Tells You                                    |
|-----------------------------------------------------------|------------------------------------------------------|
| `tls_handshake_exceeded`                                  | Failed handshakes — spikes = cert or config issue    |
| `outstanding_tls_handshake`                               | In-progress handshakes — high = slow negotiation     |
| `non_mtls_remote_conn_count` / `non_mtls_local_conn_count`| Non-mTLS connections — should be 0 in prod           |
| `perb_mtls_encryption_status`                             | mTLS encryption status                               |

### Envoy TLS Inspector
| Metric                                        | Purpose                       |
|-----------------------------------------------|-------------------------------|
| `envoy_tls_inspector_tls_found` / `tls_not_found`   | TLS vs non-TLS traffic ratio |
| `envoy_tls_inspector_sni_found` / `sni_not_found`   | SNI detection success        |
| `envoy_tls_inspector_client_hello_too_large`  | Oversized client hellos       |

### DTLS (User Plane)
| Metric                                                    | Purpose                        |
|-----------------------------------------------------------|--------------------------------|
| `pcupdp_dtls_connection_events_total_total`               | DTLS connection lifecycle      |
| `pcupdp_dtls_connection_failures_total`                   | Failed DTLS connections        |
| `pcupdp_dtls_transmitted_bytes_total` / `received_bytes_total` | DTLS throughput           |
| `pcupdp_dptls_ssl_handshakes_total`                       | SSL handshake count            |
| `pcupdp_dptls_active_connections`                         | Current active DTLS sessions   |

### SIP TLS Resources
| Metric                            | Purpose                       |
|-----------------------------------|-------------------------------|
| `sip_tls_cpu_usage_ratio_bucket`  | TLS CPU usage distribution    |
| `sip_tls_memory_usage_ratio_bucket`| TLS memory usage distribution|
| `sip_tls_kms_http_requests_total` | KMS requests for key material |

### Monitoring
```promql
# Days until cert expiry
TLS_Certificate_Expiration_Time_in_Seconds / 86400

# Handshake failure rate
rate(tls_handshake_exceeded[5m])

# Non-mTLS connections (should be 0)
non_mtls_remote_conn_count + non_mtls_local_conn_count

# DTLS failure rate
rate(pcupdp_dtls_connection_failures_total[5m]) / rate(pcupdp_dtls_connection_events_total_total[5m])
```

### Troubleshooting
| Symptom             | Check                                                    | Action                                              |
|---------------------|----------------------------------------------------------|-----------------------------------------------------|
| Cert expiring soon  | `TLS_Certificate_Expiration_Time_in_Seconds`             | Rotate cert, check auto-renewal                     |
| Handshake failures  | `tls_handshake_exceeded`, `outstanding_tls_handshake`    | Check cert chain, cipher mismatch, clock skew       |
| Non-mTLS traffic    | `non_mtls_*_conn_count`                                  | Enforce mTLS policy, check sidecar injection        |
| DTLS drops          | `pcupdp_dtls_dropped_packets_total`                      | Check network path, MTU, firewall rules             |
| High TLS CPU        | `sip_tls_cpu_usage_ratio`                                | Scale TLS termination, offload to hardware          |

---

## 3. Linux Process Metrics

> Each application pod exposes Linux kernel network counters (from `/proc/net/snmp`) prefixed by component name. These are your low-level indicators for TCP retransmissions, UDP buffer overflows, and IP-layer drops — essential for diagnosing network-level issues that don't show up at the application layer.

### Network Stack (SNMP)
| Category   | Key Metrics                                                          | Purpose                  |
|------------|----------------------------------------------------------------------|--------------------------|
| TCP        | `tcp_currestab`, `tcp_activeopens_total`, `tcp_passiveopens_total`   | Connection state         |
| TCP Errors | `tcp_retranssegs_total`, `tcp_inerrs_total`, `tcp_attemptfails_total`| Retransmissions & failures|
| UDP        | `udp_indatagrams_total`, `udp_outdatagrams_total`                    | UDP throughput           |
| UDP Errors | `udp_inerrors_total`, `udp_rcvbuferrors_total`, `udp_sndbuferrors_total` | Buffer overflows     |
| IP         | `ip_inreceives_total`, `ip_outrequests_total`, `ip_indiscards_total` | IP layer health          |
| ICMP       | `icmp_indestunreachs_total`, `icmp_outtimeexcds_total`               | Reachability issues      |

### Prefixes: `tmra_`, `tmpd_`, `datatw_`, `pcupdp_`, `pcuppep_`, `pcupupcp_`
These are per-component views of the same Linux SNMP counters. Compare across components to isolate which pod/process has network issues.

### Monitoring
```promql
# TCP retransmission rate (any component)
rate(tmra_linux_proc_net_snmp_tcp_retranssegs_total[5m])

# UDP buffer errors
rate(tmra_linux_proc_net_snmp_udp_rcvbuferrors_total[5m])

# Active TCP connections
tmra_linux_proc_net_snmp_tcp_currestab
```

### Troubleshooting
| Symptom              | Check                    | Action                                              |
|----------------------|--------------------------|-----------------------------------------------------|
| High retransmissions | `tcp_retranssegs_total`  | Check network congestion, MTU, packet loss          |
| UDP buffer errors    | `udp_rcvbuferrors_total` | Increase `net.core.rmem_max`, scale receivers       |
| Connection failures  | `tcp_attemptfails_total` | Check target health, firewall, connection limits    |
| IP discards          | `ip_indiscards_total`    | Check memory pressure, socket buffer sizes          |

---

## 4. Database Metrics

> All database access goes through a **DB Proxy** layer backed by **Redis**. The proxy tracks latency, buffer pressure, and command throughput. The **KVDB Replicator** handles cross-site geo-redundant replication. **ID Generators** allocate unique identifiers (SEID, TEID, Port IDs) using range-based allocation — exhaustion here causes session creation failures.

### DB Proxy (Redis-backed)
| Metric                                              | Purpose                                  |
|-----------------------------------------------------|------------------------------------------|
| `db_proxy_sync_total` / `db_proxy_async_total`      | Sync vs async operations                 |
| `db_proxy_latency_seconds`                          | Request latency — histogram              |
| `db_proxy_live_contexts`                            | Active DB connections                    |
| `db_proxy_hiredis_write_buffer_bytes`               | Write buffer size — high = backpressure  |
| `db_proxy_command_requests_total` / `replies_total` | Command throughput                       |

### DB Tracker
| Metric                                    | Purpose                    |
|-------------------------------------------|----------------------------|
| `db_tracker_tracking_requests_total`      | Key tracking requests      |
| `db_tracker_invalidation_statistics_total`| Cache invalidation events  |
| `db_tracker_tracked_entries`              | Currently tracked keys     |
| `db_tracker_connections_total`            | Tracker connections        |

### DB Timers
| Metric                     | Purpose              |
|----------------------------|----------------------|
| `db_processed_timers_total`| Timers processed     |
| `db_timer_operations_total`| Timer CRUD operations|

### KVDB Replicator (pckvdbrdr)
| Metric                                       | Purpose                    |
|----------------------------------------------|----------------------------|
| `pckvdbrdr_bulk_sync_percent`                | Bulk sync progress (0-100) |
| `pckvdbrdr_active_bulk_syncs`                | Ongoing bulk syncs         |
| `pckvdbrdr_receiver_message_latency_seconds` | Replication lag            |
| `pckvdbrdr_connection_pool_connections`       | Pool size                  |
| `pckvdbrdr_memory_usage_bytes`               | Replicator memory          |
| `pckvdbrdr_concurrency_control_*_flows`      | Flow control state         |

### ID Generators
| Metric                                              | Purpose                          |
|-----------------------------------------------------|----------------------------------|
| `pcupdp_db_id_gen_port_allocator_*`                 | Port ID allocation               |
| `pcuppep_db_id_gen_seid_*` / `teid_*`              | SEID/TEID allocation             |
| `*_range_allocate_max_attempts_exceeded_total`      | Allocation failures — critical   |

### Monitoring
```promql
# DB proxy latency p99
histogram_quantile(0.99, rate(tmra_db_proxy_latency_seconds_bucket[5m]))

# Write buffer pressure
tmra_db_proxy_hiredis_write_buffer_bytes_sum / tmra_db_proxy_hiredis_write_buffer_bytes_count

# Replication lag
pckvdbrdr_receiver_message_latency_seconds

# ID allocation failures
rate(pcupdp_db_id_gen_port_allocator_port_id_range_allocate_max_attempts_exceeded_total[5m])
```

### Troubleshooting
| Symptom              | Check                            | Action                                        |
|----------------------|----------------------------------|-----------------------------------------------|
| High latency         | `db_proxy_latency_seconds`       | Check Redis cluster health, network           |
| Write buffer growing | `hiredis_write_buffer_bytes`     | Redis is slow — check memory, eviction        |
| Replication stuck    | `pckvdbrdr_bulk_sync_percent`    | Check source DB, network between sites        |
| ID exhaustion        | `*_max_attempts_exceeded_total`  | Expand ID range, check for leaks              |
| Tracker errors       | `db_tracker_errors_total`        | Check Redis connectivity, key patterns        |

---

## 5. Kafka Metrics

> **Apache Kafka** handles asynchronous event streaming between platform components. TMRA and other services produce/consume messages for notifications, timer events, and state changes. Broker-level metrics track cluster health; application-level metrics (prefixed `tmra_kafka_*`) track produce/consume rates and errors.

### Broker Health
| Metric                                  | Purpose                    |
|-----------------------------------------|----------------------------|
| `kafka_server_KafkaServer_Value`        | Broker state               |
| `kafka_cluster_Partition_Value`         | Partition count            |
| `kafka_server_ReplicaManager_Value`     | Under-replicated partitions|
| `kafka_server_BrokerTopicMetrics_*`     | Per-topic throughput       |

### Request Performance
| Metric                                                    | Purpose                    |
|-----------------------------------------------------------|----------------------------|
| `kafka_network_RequestMetrics_*`                          | Request latency percentiles|
| `kafka_server_Request_queue_size`                         | Request queue depth        |
| `kafka_server_Fetch_queue_size` / `Produce_queue_size`   | Fetch/produce queue depth  |

### Application Producers/Consumers
| Metric                                       | Purpose              |
|----------------------------------------------|----------------------|
| `tmra_kafka_published_messages_total`        | Messages published   |
| `tmra_kafka_consumed_messages_total`         | Messages consumed    |
| `tmra_kafka_published_bytes_total`           | Publish throughput   |
| `tmra_kafka_errors_total`                    | Kafka errors         |
| `tmra_kafka_consumers` / `tmra_kafka_producers` | Active client count|

### Monitoring
```promql
# Message produce rate
rate(tmra_kafka_published_messages_total[5m])

# Consumer lag (compare published vs consumed)
sum(tmra_kafka_published_messages_total) - sum(tmra_kafka_consumed_messages_total)

# Kafka error rate
rate(tmra_kafka_errors_total[5m])

# Under-replicated partitions
kafka_server_ReplicaManager_Value{name="UnderReplicatedPartitions"}
```

### Troubleshooting
| Symptom          | Check                          | Action                                            |
|------------------|--------------------------------|---------------------------------------------------|
| Consumer lag     | Published vs consumed totals   | Scale consumers, check processing time            |
| Produce errors   | `kafka_errors_total`           | Check broker health, disk space, replication      |
| High latency     | `RequestMetrics` percentiles   | Check disk I/O, network, partition balance        |
| Under-replicated | `ReplicaManager`               | Check broker connectivity, disk health            |

---

## 6. Envoy Proxy Metrics

> **Envoy** is the L7 ingress proxy sitting in front of all services. It handles TLS termination, routing, RBAC enforcement, retries, and circuit breaking. Downstream metrics = client-facing traffic; upstream metrics = traffic to backend pods. When something breaks, Envoy metrics are usually the first place you'll see it.

### Downstream (Client-facing)
| Metric                                  | Purpose            |
|-----------------------------------------|--------------------|
| `envoy_http_downstream_rq_total`        | Total requests     |
| `envoy_http_downstream_rq_xx`           | Response codes     |
| `envoy_http_downstream_cx_active`       | Active connections |
| `envoy_http_downstream_rq_time_bucket`  | Request latency    |

### Upstream (Backend)
| Metric                                          | Purpose                |
|-------------------------------------------------|------------------------|
| `envoy_cluster_upstream_rq_total`               | Requests to backends   |
| `envoy_cluster_upstream_rq_xx`                  | Backend response codes |
| `envoy_cluster_upstream_cx_connect_ms_bucket`   | Connection setup time  |
| `envoy_cluster_upstream_rq_timeout`             | Request timeouts       |
| `envoy_cluster_upstream_rq_retry`               | Retries                |

### Cluster Health
| Metric                                    | Purpose                      |
|-------------------------------------------|------------------------------|
| `envoy_cluster_membership_healthy`        | Healthy endpoints            |
| `envoy_cluster_membership_total`          | Total endpoints              |
| `envoy_cluster_lb_healthy_panic`          | Panic mode (all unhealthy)   |
| `envoy_cluster_upstream_cx_connect_fail`  | Connection failures          |

### TLS & Security
| Metric                                    | Purpose                      |
|-------------------------------------------|------------------------------|
| `envoy_cluster_ssl_handshake`             | Successful TLS handshakes    |
| `envoy_cluster_ssl_connection_error`      | TLS errors                   |
| `envoy_http_rbac_denied`                  | RBAC policy denials          |
| `envoy_listener_ssl_fail_verify_error`    | Cert verification failures   |

### Monitoring
```promql
# 5xx error rate
rate(envoy_http_downstream_rq_xx{envoy_response_code_class="5"}[5m]) / rate(envoy_http_downstream_rq_total[5m])

# Upstream latency p99
histogram_quantile(0.99, rate(envoy_cluster_upstream_rq_time_bucket[5m]))

# Unhealthy endpoints
envoy_cluster_membership_total - envoy_cluster_membership_healthy

# Connection pool exhaustion
envoy_cluster_upstream_cx_overflow
```

### Troubleshooting
| Symptom             | Check                                                  | Action                                        |
|---------------------|--------------------------------------------------------|-----------------------------------------------|
| 5xx spike           | `upstream_rq_xx`, `upstream_cx_connect_fail`           | Check backend health, circuit breakers        |
| High latency        | `upstream_rq_time`, `downstream_rq_time`               | Compare to isolate envoy vs backend           |
| Connection overflow | `upstream_cx_overflow`, `upstream_cx_pool_overflow`    | Increase circuit breaker limits               |
| TLS errors          | `ssl_connection_error`, `ssl_fail_verify_*`            | Check cert chain, expiry, SAN mismatch        |
| RBAC denials        | `http_rbac_denied`                                     | Review RBAC policy, check source identity     |

---

## 7. MinIO (Object Storage)

> **MinIO** provides S3-compatible object storage with erasure coding for data durability. It stores artifacts, backups, and bulk data. Key concerns: drive health (offline drives reduce fault tolerance), capacity planning, and S3 request error rates.

### Cluster Health
| Metric                                                 | Purpose              |
|--------------------------------------------------------|----------------------|
| `minio_cluster_health_status`                          | Overall cluster health|
| `minio_cluster_nodes_online_total` / `offline_total`   | Node availability    |
| `minio_cluster_drive_online_total` / `offline_total`   | Drive availability   |
| `minio_cluster_write_quorum`                           | Write quorum status  |

### Storage Capacity
| Metric                                                 | Purpose                      |
|--------------------------------------------------------|------------------------------|
| `minio_capacity_raw_total` / `raw_free_total`          | Raw disk capacity            |
| `minio_capacity_usable_total` / `usable_free_total`   | Usable after erasure coding  |
| `minio_cluster_usage_total_bytes`                      | Total data stored            |
| `minio_cluster_usage_object_total`                     | Total objects                |

### Request Performance
| Metric                                                 | Purpose              |
|--------------------------------------------------------|----------------------|
| `minio_s3_requests_total`                              | Total S3 requests    |
| `minio_s3_requests_4xx_errors_total` / `5xx_errors_total` | Error rates       |
| `minio_s3_requests_ttfb_seconds_distribution`          | Time to first byte   |
| `minio_s3_traffic_received_bytes` / `sent_bytes`       | Throughput           |

### Node Resources
| Metric                          | Purpose              |
|---------------------------------|----------------------|
| `minio_node_cpu_avg_load1`      | CPU load             |
| `minio_node_mem_used_perc`      | Memory usage %       |
| `minio_node_drive_free_bytes`   | Per-drive free space |
| `minio_node_drive_latency_us`   | Drive latency        |

### Monitoring
```promql
# S3 error rate
rate(minio_s3_requests_5xx_errors_total[5m]) / rate(minio_s3_requests_total[5m])

# Usable capacity remaining %
minio_capacity_usable_free_total / minio_capacity_usable_total * 100

# Drive latency
minio_node_drive_latency_us

# Offline drives (should be 0)
minio_cluster_drive_offline_total
```

### Troubleshooting
| Symptom        | Check                          | Action                                        |
|----------------|--------------------------------|-----------------------------------------------|
| 5xx errors     | `s3_requests_5xx_errors_total` | Check drive health, quorum, memory            |
| Slow TTFB      | `s3_requests_ttfb_seconds`     | Check drive latency, network, CPU             |
| Drive offline  | `cluster_drive_offline_total`  | Replace drive, check erasure set health       |
| Capacity low   | `capacity_usable_free_total`   | Add drives, enable lifecycle/ILM policies     |

---

## 8. Redis

> **Redis** is the primary in-memory data store used for session state, caching, and ID generation. It runs in cluster mode for horizontal scaling. Key concerns: memory usage (eviction starts when full), fragmentation (wastes RAM), replication lag (data loss risk), and cluster state (slot failures break reads/writes).

### Connection & Clients
| Metric                              | Purpose                        |
|-------------------------------------|--------------------------------|
| `redis_connected_clients`           | Current client connections     |
| `redis_blocked_clients`             | Clients blocked on BLPOP etc.  |
| `redis_rejected_connections_total`  | Rejected (maxclients hit)      |
| `redis_connected_slaves`            | Replica count                  |

### Memory
| Metric                            | Purpose                                  |
|-----------------------------------|------------------------------------------|
| `redis_memory_used_bytes`         | Total memory used                        |
| `redis_memory_max_bytes`          | Max memory configured                    |
| `redis_mem_fragmentation_ratio`   | Fragmentation — >1.5 is bad              |
| `redis_evicted_keys_total`        | Keys evicted due to memory pressure      |

### Performance
| Metric                                    | Purpose            |
|-------------------------------------------|--------------------|  
| `redis_commands_processed_total`          | Command throughput |
| `redis_commands_duration_seconds_total`   | Command latency    |
| `redis_keyspace_hits_total` / `misses_total` | Cache hit ratio |
| `redis_slowlog_length`                    | Slow queries count |

### Replication
| Metric                                | Purpose                              |
|---------------------------------------|--------------------------------------|
| `redis_master_repl_offset`            | Master replication offset            |
| `redis_connected_slave_offset_bytes`  | Slave offset — lag = master - slave  |
| `redis_master_link_up`                | Replication link status              |

### Cluster
| Metric                                    | Purpose            |
|-------------------------------------------|--------------------|  
| `redis_cluster_state`                     | Cluster OK or FAIL |
| `redis_cluster_slots_ok` / `slots_fail`  | Slot health        |
| `redis_cluster_known_nodes`               | Nodes in cluster   |

### Monitoring
```promql
# Memory usage %
redis_memory_used_bytes / redis_memory_max_bytes * 100

# Cache hit ratio
redis_keyspace_hits_total / (redis_keyspace_hits_total + redis_keyspace_misses_total)

# Command latency
rate(redis_commands_duration_seconds_total[5m]) / rate(redis_commands_processed_total[5m])

# Replication lag bytes
redis_master_repl_offset - on(instance) redis_connected_slave_offset_bytes
```

### Troubleshooting
| Symptom         | Check                                        | Action                                    |
|-----------------|----------------------------------------------|-------------------------------------------|
| High memory     | `memory_used_bytes`, `evicted_keys_total`    | Increase maxmemory, check key TTLs        |
| Fragmentation   | `mem_fragmentation_ratio` > 1.5              | Restart Redis or enable active defrag     |
| Slow commands   | `slowlog_length`, `commands_duration`        | Optimize queries, avoid O(N) commands     |
| Replication lag | Slave offset vs master offset                | Check network, slave disk I/O             |
| Cluster FAIL    | `cluster_state`, `slots_fail`                | Check failed nodes, rebalance slots       |

---

## 9. Elasticsearch

> **Elasticsearch** is the centralized log aggregation and search engine, fed by DDC collectors. It stores application logs, audit trails, and diagnostic data. Key concerns: cluster health (red = data loss), JVM heap pressure (causes GC pauses), unassigned shards (blocks indexing), and disk capacity.

### Cluster Health
| Metric                                            | Purpose              |
|---------------------------------------------------|----------------------|
| `elasticsearch_cluster_health_status`             | Green/Yellow/Red     |
| `elasticsearch_cluster_health_active_shards`      | Active shards        |
| `elasticsearch_cluster_health_unassigned_shards`  | Unassigned — should be 0 |
| `elasticsearch_cluster_health_number_of_nodes`    | Node count           |

### Indexing & Search
| Metric                                                        | Purpose          |
|---------------------------------------------------------------|------------------|
| `elasticsearch_indices_indexing_index_total`                   | Documents indexed|
| `elasticsearch_indices_indexing_index_time_seconds_total`      | Indexing time    |
| `elasticsearch_indices_search_query_total`                     | Search queries   |
| `elasticsearch_indices_search_query_time_seconds`              | Search latency   |

### Resources
| Metric                                        | Purpose        |
|-----------------------------------------------|----------------|
| `elasticsearch_jvm_memory_used_bytes`         | JVM heap usage |
| `elasticsearch_jvm_memory_max_bytes`          | JVM heap max   |
| `elasticsearch_jvm_gc_collection_seconds_sum` | GC time        |
| `elasticsearch_os_cpu_percent`                | CPU usage      |
| `elasticsearch_filesystem_data_free_bytes`    | Disk free      |

### Monitoring
```promql
# Cluster status (0=green, 1=yellow, 2=red)
elasticsearch_cluster_health_status

# JVM heap usage %
elasticsearch_jvm_memory_used_bytes / elasticsearch_jvm_memory_max_bytes * 100

# Indexing rate
rate(elasticsearch_indices_indexing_index_total[5m])

# Search latency
rate(elasticsearch_indices_search_query_time_seconds[5m]) / rate(elasticsearch_indices_search_query_total[5m])
```

### Troubleshooting
| Symptom      | Check                                | Action                                              |
|--------------|--------------------------------------|-----------------------------------------------------|
| Red cluster  | `unassigned_shards`, node count      | Check failed nodes, disk space, shard allocation    |
| Slow search  | `search_query_time`, `jvm_gc`        | Optimize queries, increase heap, add nodes          |
| High GC      | `jvm_gc_collection_seconds`          | Reduce heap pressure, check bulk indexing           |
| Disk full    | `filesystem_data_free_bytes`         | Add nodes, delete old indices, enable ILM           |


