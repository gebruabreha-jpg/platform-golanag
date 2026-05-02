1,CRE (Cloud Routing Engine) — eric-pc-routing-engine:-
The brain of routing in PCG
Runs BGP sessions with DCGW (one per network instance: sgi1, sgi2, media, ran, signaling)
Advertises NAT prefixes to DCGW (e.g., 16.0.0.0/19 for CGNAT pool)
Withdraws prefixes when NAT IPs are released
Handles BFD for fast failure detection
In GeoRed: manages GSA state (ACTIVE/STANDBY), controls which cluster advertises routes
From your TCs:
Then check correct information on cluster 0 on cre:
  | ActionCLI                                                                       | Entry           | ExpectedResult |
  | network-instances network-instance sgi1 routing bgp 10852 rib neighbors-summary | estab-peers-num | 2              |
CRE — Does Routing, NOT Forwarding
CRE never touches user packets. It only:
Publishes routes — tells DCGW "send traffic for 16.0.0.0/19 to me"
Withdraws routes — tells DCGW "stop sending traffic for this prefix"
Learns routes — receives routes from DCGW (e.g., default route to internet)
Installs routes — pushes forwarding rules into DP's forwarding table
Think of CRE as a map maker. It draws the map (routing table) but doesn't drive the car (forward packets).
CRE says to DCGW via BGP:
  "I have 16.0.0.0/19 → send it to me"
  "I have 19.20.0.64/27 → send it to me"
DCGW updates its routing table:
  16.0.0.0/19 → next-hop PCG Cluster A
  19.20.0.64/27 → next-hop PCG Cluster A

CRE — Routes & Prefixes (no sessions, no packets)
CRE is the routing control plane. It publishes/withdraws BGP routes to DCGW so traffic reaches the right cluster. It never touches user packets.
Key metrics: BGP peers up? Prefixes correct? Active or standby?



2,RA (Routing Aggregator) — eric-pc-routing-aggregator:-
Aggregates routing information from multiple DPs
Tells CRE which prefixes to advertise based on active NAT allocations across all DP pods
Bridges between DP (many pods) and CRE (few pods)


3,DP (Data Plane) — eric-pc-up-data-plane:-
The packet processing engine — forwards all user traffic
Handles: GTP encap/decap, CGNAT translation, DPI, QoS, IPFIX reporting
CPU roles:
ingress_egress — packet I/O (NIC to memory)
input_output — packet processing (NAT, DPI, forwarding)
control — session management
spare — overflow/background tasks
Multiple pods (2-30 depending on environment)
DP — Does BOTH Forwarding AND Packet Processing
DP does everything with the actual packets:-
Forwarding (moving packets):
Packet arrives on ran_link (from gNodeB)
  → DP looks up forwarding table: "destination 8.8.8.8 → send out sgi1_link"
  → Packet sent out sgi1_link toward DCGW
Packet Processing (transforming packets):-
Same packet, but before forwarding:
1. GTP decap:     Remove GTP tunnel header (ran side)
2. DPI:           Inspect → "this is HTTP traffic"
3. CGNAT:         Change source IP 10.0.0.5 → 16.0.0.100:12345
4. QoS:           Mark packet priority
5. IPFIX:         Generate NAT translation event record
6. Forwarding:    Send out sgi1_link
DP — Forwarding + Packet Processing
DP does everything with actual packets: forward them (routing table lookup) and process them (GTP, NAT, DPI, QoS, IPFIX).
Session vs Flow:
1 session = 1 subscriber's NAT mapping (e.g., UE gets public IP 16.0.0.100)
1 flow = 1 TCP/UDP connection (e.g., YouTube stream on port 12345)
1 session → many flows. 140K sessions → 33K+ active flows
Key metrics: throughput (Gbps), packet drop (0 PPM), NAT sessions/flows, CPU per role, memory



4,PEP (PFCP Endpoint) — eric-pc-up-pfcp-endpoint:-
Receives PFCP messages from SMF/PGW-C
Session establishment, modification, deletion
Translates PFCP rules into DP forwarding rules
Distributes sessions across DP pods


5,KVDB (Key-Value Database) — eric-pc-kvdb-rd-server:-
Stores session state (NAT mappings, PFCP sessions)
Replicates data between Active/Standby clusters (GeoRed bulk sync)
pckvdbrdr_bulk_sync_percent = 100 means all data synced
KVDB — Session Replication (GeoRed)
Stores session state, replicates between Active/Standby clusters.
Key metrics: bulk sync = 100%, session parity < 0.1% between clusters






Metrics by Component
CRE:-
estab-peers-num — BGP peers per NI
gsa prefix, prefix-count — advertised prefixes
pcupdp_geored_client_state — 2=ACTIVE, 3=STANDBY
gsa-state — ACTIVE (WITH PEER) / STANDBY
ready-to-switchover — Yes/No

DP — Traffic:-
pcupdp_gtp_sent_bytes_total / pcupdp_gtp_received_bytes_total — GTP throughput
pcupdp_payload_received_bytes_total / pcupdp_payload_sent_bytes_total — payload throughput
pcupdp_ip_interface_received_bytes_total / pcupdp_ip_interface_transmitted_bytes_total — per-interface throughput
pcupdp_core_interface_packet_drop_ppm / pcupdp_access_interface_packet_drop_ppm — packet loss
pcupdp_upf_engine_dropped_packets_total — dropped packets
pcupdp_upf_sent_error_messages_total — error messages (target: 0)

DP — NAT:-
pcupdp_nat_active_sessions — subscribers with NAT mapping
pcupdp_nat_active_flows — active TCP/UDP connections
pcupdp_nat_ip_addresses{policy_id} — public IPs in use
pcupdp_nat_received_packets_total / pcupdp_nat_sent_packets_total — NAT PPS
pcupdp_nat_dropped_packets_total — NAT drops
pcupdp_nat_received_bytes_total / pcupdp_nat_sent_bytes_total — NAT bytes
pcupdp_nat_ueipv4_passthrough_checks_total{result} — NPT checks
pcupdp_nat_flow_passthroughs_total{reason} — passthrough flows
pcupdp_nat_garbage_collector_ips — IPs being cleaned up
pcupdp_nat_pool_active_flows / pcupdp_nat_pool_sent_packets_total / pcupdp_nat_pool_received_packets_total — per-policy counters
pcupdp_nat_sessions{policy_id} — sessions per policy

DP — CPU & Memory:-
pcupdp_cpu_load_15s_percent{role=ingress_egress|input_output|control|spare} — CPU per role
pcupdp_cpu_load_1h_peak_percent{role} — peak CPU
pcupdp_cpu_load_15s_peak_percent{role} — 15s peak
pcupdp_memory_used_bytes — total DP memory
pcupdp_packet_detection_max_session_flow_limit_hits_total — flow table overflow (target: 0)
pcupdp_packet_detection_max_memory_flow_hits_total — memory overflow (target: 0)

DP — DPI:-
pcupdp_application_packets_total{direction, application} — packets per app
pcupdp_dpi_active_sessions{type} — DPI sessions

DP — IPFIX Reporting:
pcupdp_event_reporting_succeeded_events_total{type=session_establishment|session_deletion|translation_44_create|translation_44_delete|translation_64_create|translation_64_delete|port_set_allocation|port_set_deallocation} — event rates
pcupdp_reporting_overload_total{threshold="high"} — overload events
pcupdp_reporting_overload_restore_total — overload recovery
pcupdp_event_reporting_transaction_round_trip_time_high_prio_max_seconds — reporting latency
pcupdp_ipfix_session_events_sent_to_erc_total / pcupdp_ipfix_port_set_events_sent_to_erc_total — IPFIX event counts

DP — Firewall:
pcupdp_fw_session_creations_total / pcupdp_fw_session_deletions_total — FW session events
pcupdp_fw_unlicensed_skipped_sessions — skipped sessions

DP — License:
pcupdp_license_feature_state{feature_id, license_key} — license status (1=active, 2=enabled)

PEP:
pcuppep_pfcp_total_sessions — total active sessions
pcuppep_pfcp_received_messages_total{status, type=session_establishment_request|session_modification_request|session_deletion_request|session_report_response} — PFCP message rates
pcuppep_pfcp_sent_messages_total{status="no_response", type} — failed messages
pcuppep_pfcp_sent_messages_total{retransmission="true", type} — retransmissions

KVDB:
pckvdbrdr_bulk_sync_percent — sync progress (target: 100)
pckvdbrdr_bulk_sync_duration_seconds — sync time
Session Parity (cross-cluster):
pcuppep_pfcp_total_sessions differ < 0.1%
pcupdp_upf_pfcp_cached_sessions differ < 0.1%

Container-Level (PMRM — all components):
pmrm_container_cpu_usage_percent{workload_name} — CPU per component
pmrm_container_mem_usage_bytes{workload_name} — memory per component
pmrm_container_mem_usage_percent{workload_name} — memory % (delta for leak detection)

DP Session Counters:-
pcupdp_upf_ipv4_sessions / pcupdp_upf_ipv6_sessions / pcupdp_upf_dual_stack_sessions — sessions by IP type
pcupdp_upf_pfcp_cached_sessions — cached sessions in DP

Dallas/TRex:
total signaling success rate — > 99.99%
total signaling resending rate — < 0.05%
traffic UL/DL loss ppm — < 10

System Crashes Attempted — 0
Error Indication sent — < 200

