# Reliability Patterns for Microservices (3R: Reliable, Robust, Resilient)
| Term                  | Meaning                                   | Think |
| Reliable              | Works correctly and consistently          | "Does it work as expected?" |
| Robust                | Handles errors/bad input without crashing | "Does it handle problems well?" |
| Resilient             | Recovers quickly from failures            | "Does it bounce back after failure?" |

## Scenario: E-Commerce Online Shop
Three services: **Order Service**, **Payment Service**, **Stock Service**
---
# Phase 1 — Build It (Get 3 services running)
## 1. API Gateway (Single Entry Point)
You need one front door. Without this, clients call services directly and you have chaos.
```
/api/orders   → Gateway → Order Service
/api/payments → Gateway → Payment Service
/api/stock    → Gateway → Stock Service
```

```java
@Bean
public RouteLocator routes(RouteLocatorBuilder builder) {
    return builder.routes()
        .route("order-service", r -> r.path("/orders/**")
            .uri("lb://ORDER-SERVICE"))
        .route("payment-service", r -> r.path("/payments/**")
            .uri("lb://PAYMENT-SERVICE"))
        .build();
}
```

## 2. Database per Service

Each service owns its data. No shared DB — otherwise you have a distributed monolith.

```
Order Service   → order_db   (orders table)
Payment Service → payment_db (payments table)
Stock Service   → stock_db   (inventory table)
```

```java
// Order Service entity
@Entity
public class Order {
    @Id @GeneratedValue
    private Long id;
    private String status; // CREATED, COMPLETED, CANCELLED
}
```

Each service has its own `application.yml` pointing to its own DB.

## 3. Saga Pattern (Simple Choreography)

You need this because: Order → Payment → Stock spans 3 DBs. No single transaction can cover all 3.

```
1. Create Order  ──→  2. Take Payment  ──→  3. Reserve Stock
        ↑ cancel            ↑ refund              ✗ fails
```

**Order Service** — starts the saga:

```java
@RestController
@RequestMapping("/orders")
public class OrderController {

    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;

    @PostMapping
    public String createOrder() {
        // save order to DB with status=CREATED
        kafkaTemplate.send("order-events", "ORDER_CREATED");
        return "Order Created";
    }
}
```

**Payment Service** — listens, processes, continues or rolls back:

```java
@KafkaListener(topics = "order-events")
public void handleOrder(String event) {
    if (event.equals("ORDER_CREATED")) {
        String result = paymentService.processPayment();
        if (result.equals("Payment Success")) {
            kafkaTemplate.send("payment-events", "PAYMENT_SUCCESS");
        } else {
            kafkaTemplate.send("order-events", "ORDER_CANCELLED");
        }
    }
}
```

**Stock Service** — listens, reserves or triggers refund:

```java
@KafkaListener(topics = "payment-events")
public void handlePayment(String event) {
    if (event.equals("PAYMENT_SUCCESS")) {
        String result = stockService.reserveStock();
        if (result.equals("Stock Reserved")) {
            kafkaTemplate.send("order-events", "ORDER_COMPLETED");
        } else {
            kafkaTemplate.send("payment-events", "PAYMENT_REFUND");
        }
    }
}
```

**Compensation handlers** (undo on failure):

```java
@KafkaListener(topics = "payment-events")
public void refund(String event) {
    if (event.equals("PAYMENT_REFUND")) {
        // reverse the payment
    }
}

@KafkaListener(topics = "order-events")
public void cancel(String event) {
    if (event.equals("ORDER_CANCELLED")) {
        // mark order as cancelled
    }
}
```

**Phase 1 architecture:**
```
Client → API Gateway → Order Service → Kafka → Payment Service → Kafka → Stock Service
                                                      ↕
                                              Compensation (Saga)
```
At this point you have 3 working services with event-driven communication and rollback. **Ship it.**
---










# Phase 2 — Harden It (When things start failing)
Add these when you see real failures in production logs.
## 4. Circuit Breaker
**When to add:** Payment service starts timing out and Order service hangs waiting for it.

```java
@CircuitBreaker(name = "paymentCB", fallbackMethod = "fallback")
public String processPayment() {
    return paymentClient.call(); // external call
}

public String fallback(Exception e) {
    return "Payment service is temporarily unavailable";
}
```

States: `CLOSED` (normal) → `OPEN` (failing, skip calls) → `HALF_OPEN` (test if recovered)

## 5. Retry with Exponential Backoff

**When to add:** You see transient network errors (connection reset, timeout) in logs.

```
1st fail → wait 1s → retry
2nd fail → wait 2s → retry
3rd fail → wait 4s → give up
```

```java
@Retry(name = "paymentRetry")
@CircuitBreaker(name = "paymentCB", fallbackMethod = "fallback")
public String processPayment() {
    return paymentClient.call();
}
```

`application.yml`:

```yaml
resilience4j:
  retry:
    instances:
      paymentRetry:
        maxAttempts: 3
        waitDuration: 1s
        exponentialBackoffMultiplier: 2
```

## 6. Bulkhead

**When to add:** Stock service is slow and it's eating all threads, making Payment service unresponsive too.

```java
@Bulkhead(name = "stockBulkhead", type = Bulkhead.Type.THREADPOOL)
public String reserveStock() {
    return stockClient.call();
}
```

```yaml
resilience4j:
  bulkhead:
    instances:
      stockBulkhead:
        maxConcurrentCalls: 10
```

Each service gets its own thread pool — one slow service can't starve the others.

## 7. Idempotence

**When to add:** Users double-click "Pay" or Kafka retries deliver duplicate messages.

```java
public void processPayment(String requestId) {
    if (redisTemplate.hasKey(requestId)) return; // already processed

    // process payment
    redisTemplate.opsForValue().set(requestId, "done", 24, TimeUnit.HOURS);
}
```

**Phase 2 architecture:**

```
Client → API Gateway → Order Service → Kafka → Payment Service → Kafka → Stock Service
                                                  ↕ Retry/CB          ↕ Bulkhead
                                              Compensation (Saga)
                                              Idempotence (Redis)
```
---














# Phase 3 — Scale It (When traffic grows)

Add these only when monitoring shows bottlenecks.

## 8. Cache

**When to add:** DB queries for product/stock data are slow and repetitive.

```java
@Cacheable("products")
public Product getProduct(Long id) {
    return repository.findById(id).orElse(null);
}
```

Use Redis. Set TTL so stale data expires.

## 9. CQRS + Event Sourcing

**When to add:** Read traffic is 100x write traffic, and your single DB can't handle both.

**CQRS** — separate write model from read model:

```java
// WRITE → primary DB
@PostMapping("/orders")
public void create(@RequestBody Order order) {
    repository.save(order);
    kafkaTemplate.send("order-events", "ORDER_CREATED:" + order.getId());
}

// READ → denormalized read store (separate DB or Elasticsearch)
@GetMapping("/orders")
public List<OrderView> getAll() {
    return readRepository.findAll();
}
```

**Event Sourcing** — store events, rebuild state by replaying:

```
OrderCreated → OrderPaid → OrderShipped → OrderDelivered
```

Only add this if you need full audit trail or temporal queries. Otherwise it's overkill.

---

## Quick Reference

| Phase | Pattern | Purpose | Tool |
|---|---|---|---|
| 1 - Build | API Gateway | Routing | Spring Cloud Gateway |
| 1 - Build | DB per Service | Loose coupling | Separate datasources |
| 1 - Build | Saga | Distributed txn | Kafka events |
| 2 - Harden | Circuit Breaker | Stop cascading failure | Resilience4j |
| 2 - Harden | Retry + Backoff | Handle transient errors | Resilience4j |
| 2 - Harden | Bulkhead | Isolate resources | Resilience4j |
| 2 - Harden | Idempotence | Prevent duplicates | Redis |
| 3 - Scale | Cache | Reduce DB load | Redis, Spring Cache |
| 3 - Scale | CQRS | Optimize read/write | Separate models |

![alt text](image.png)
