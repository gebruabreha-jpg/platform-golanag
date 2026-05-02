# System Design Heuristics & Patterns
## 1. High Latency → CDN
**Problem:** Users far from origin server experience slow load times.  
**Example:** Expedia serves hotel images via CloudFront CDN. A user in Tokyo gets images from a nearby edge node (~30ms) instead of US-East origin (~200ms).

## 2. Read Bottleneck → Cache
**Problem:** Database overwhelmed by repeated read queries.  
**Example:** Expedia caches popular hotel search results in Redis (TTL=5min). 80% of searches hit cache instead of querying the DB every time.
```
User -> API -> Redis (HIT?) -> return cached result
                  (MISS?) -> DB -> store in Redis -> return
```

## 3. Write Spike → Queue
**Problem:** Sudden burst of writes (e.g., flash sale) overwhelms the DB.  
**Example:** During a Black Friday sale, booking requests are pushed to an SQS/Kafka queue. Workers consume at a steady rate, preventing DB overload.
```
User -> API -> Kafka Topic -> Consumer Workers -> DB
         (buffered, async)     (controlled rate)
```

## 4. Large Files → Object Storage (S3/Blob)
**Problem:** Storing images/videos in a relational DB is expensive and slow.  
**Example:** Hotel photos are stored in S3. The DB only stores the S3 URL reference.

## 5. Complex Pre-computation → Precompute + Store
**Problem:** Expensive queries computed on every request.  
**Example:** Expedia precomputes "top 10 hotels per city" nightly and stores results in a read-optimized table, avoiding expensive aggregation at query time.

## 6. High Request Rate → Rate Limiter
**Problem:** A single client or bot floods the API.  
**Example:** API Gateway limits each user to 100 search requests/minute. Excess requests get HTTP 429.

## 7. Too Much Data on One DB → Sharding
**Problem:** Single DB can't handle the data volume.  
**Example:** Shard bookings DB by `user_id % N`. User 12345 always goes to shard 5.

**Caveat:** Sharding breaks cross-shard joins. Often better to start with read replicas first (can cut latency ~40%) before sharding.

## 8. Single Point of Failure → Replication
**Problem:** One server goes down, entire system is unavailable.  
**Example:** Primary DB in us-east-1 replicates to a standby in us-west-2. On failure, DNS failover switches traffic in <30s.

## 9. Consistency + Contention → Partition by Access Pattern
**Problem:** Data grouped the wrong way causes lock contention.  
**Example:** Instead of one `bookings` table for all hotels, partition by `hotel_id`. Concurrent bookings for different hotels never contend for the same lock.

## Quick Decision Table
| Symptom                       | Pattern         | Tool     |
| High latency (static)         | CDN             | CloudFront, Akamai |
| Read bottleneck               | Cache           | Redis, Memcached |
| Write spike                   | Queue           | Kafka, SQS, RabbitMQ |
| Large files                   | Object Storage  | S3, GCS |
| Expensive queries             | Precompute      | Materialized Views, Cron |
| Request flood                 | Rate Limiter    | API Gateway, Nginx |
| Data too large                | Sharding        | Consistent Hashing |
| Single point of failure       | Replication/HA  | Primary-Replica, Multi-AZ |
| Lock contention               | Partition       | Partition by key |

![alt text](image-1.png)
![alt text](image-2.png)
![alt text](image-3.png)
![alt text](image-4.png)
![alt text](image-5.png)
![alt text](image-6.png)
