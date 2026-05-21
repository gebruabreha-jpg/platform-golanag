# 1booking-zspringjava

Spring Boot backend for the 1booking platform.

## Features

- JWT-based authentication with access/refresh tokens
- PostgreSQL database with JPA/Hibernate
- Redis caching
- Flyway database migrations
- OpenAPI/Swagger documentation
- Input validation

## Development

### Prerequisites

- Java 17+
- Maven 3.9+
- PostgreSQL
- Redis

### Run

```bash
./mvnw spring-boot:run
```

### Test

```bash
./mvnw test
```

### API Docs

http://localhost:8080/swagger-ui.html

## Environment Variables

```env
DATABASE_URL=jdbc:postgresql://localhost:5432/onepagecommerce_db
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=postgres
REDIS_HOST=localhost
REDIS_PORT=6379
JWT_SECRET=your-secret-key-min-32-chars
JWT_REFRESH_SECRET=your-refresh-secret-key
```