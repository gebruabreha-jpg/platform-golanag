# 1Booking TypeScript Platform

A full-featured booking platform backend built with **Node.js**, **TypeScript**, **Express**, and **PostgreSQL**. Supports flight bookings, peer-to-peer shipping, housing marketplace, and services marketplace — modeled on the [1booking-golang](..) project.

## 📁 Stack

| Layer     | Technology                 |
|-----------|----------------------------|
| Runtime   | Node.js ≥ 18 / TypeScript  |
| Framework | Express                    |
| ORM       | Drizzle ORM / Prisma       |
| Database  | PostgreSQL 16              |
| Cache     | Redis 7                    |
| Auth      | JWT (Access + Refresh)     |
| Password  | bcrypt (12 rounds)         |
| Validation| Zod + express-validator    |
| Logging   | Pino (structured JSON)     |
| Transport | REST (Express + TypeScript)|
| Testing   | Jest + Supertest           |
| Docker    | Docker Compose             |

## 📚 Features

### Services
- **Flight Booking** – Search, book, and manage flight reservations
- **Peer-to-Peer Shipping** – Send/receive items via verified travelers
- **Housing Marketplace** – Room/accommodation listings & bookings
- **Services Marketplace** – Find lawyers, doctors, and skilled professionals
- **Job Board** – Posting and discovering opportunities
- **Scholarships** – Listing and tracking academic opportunities
- **Community Forums** – Category-based communities & posts
- **Trust & Verification** – Review and rating system with trust scores
- **Payment & Escrow** – Secure escrow-backed transactions via Stripe

### Cross-Cutting Concerns
- **JWT Auth** – Access token + refresh token rotation with Redis revocation
- **RBAC** – DIASPORA / LOCAL / MERCHANT / ADMIN roles with per-route guards
- **Rate Limiting** – Redis-backed sliding window (100 req/min per IP+route)
- **Input Validation** – Zod schemas with expressive error messages
- **Structured Logging** – Pino JSON logs with request correlation IDs
- **Health Checks** – `/health` endpoint, ready/live probes for Docker/K8s
- **Open API** – Shared types, interfaces, and response helpers

## 🏗️ Project Structure

```
1booking-typescript/
├── src/
│   ├── index.ts            # ─── App entry, DI wiring, server startup
│   ├── config/             # ─── Environment config validation (Zod)
│   ├── types/              # ─── Domain & API type definitions
│   ├── constants/roles.ts and index.ts  # ─── RBAC enums
│   ├── errors/             # ─── AppError class, error codes, handler
│   ├── utils/              # ─── Logger, crypto, response, pagination, date
│   ├── middleware/         # ─── Auth, rate-limit, CORS, notFound, errorHandler
│   ├── schemas/            # ─── Zod validation schemas
│   ├── repositories/       # ─── Repository interfaces
│   ├── services/           # ─── Service interfaces
│   ├── controllers/        # ─── HTTP route handlers
│   ├── routes/             # ─── Express routers
│   ├── clients/            # ─── PostgreSQL, Redis, Stripe clients
│   ├── database/migrations # ─── SQL migration files
│   └── metrics/            # ─── OpenTelemetry / Prometheus metrics
├── tests/                  # ─── Jest unit, integration, e2e tests
├── Makefile                # ─── Build, test, lint, docker targets
├── Dockerfile              # ─── Multi-stage production build
├── docker-compose.yml      # ─── PostgreSQL, Redis, backend, pgadmin
├── tsconfig.json           # ─── Compiler options
└── .env.example            # ─── Environment variable template
```

## 🚀 Quick Start

```bash
# 1 – Clone & install dependencies
cd onepagecommerce-typescript
npm install

# 2 – Start infra services
docker compose up -d postgres redis pgadmin

# 3 – Configure environment variables
cp .env.example .env
# edit .env – set DATABASE_URL, REDIS_URL, JWT_SECRET …

# 4 – Run migrations
npm run db:migrate

# 5 – Start the dev server
npm run dev
```

`curl http://localhost:3100/health` should return:

```json
{"status":"ok","timestamp":"2025-01-01T00:00:00.000Z"}
```

## 🛠️ Makefile Targets

| Target        | Description                                 |
|---------------|---------------------------------------------|
| `make dev`          | Start all services (Docker Compose) |
| `make build`        | Build all Docker images             |
| `make test`         | Run all test suites                 |
| `make lint`         | Run ESLint + Prettier checks        |
| `make typecheck`    | Run TypeScript compiler             |
| `make docker-up`    | Start infra services                |
| `make docker-down`  | Stop all services                   |
| `make logs`         | Tail all service logs               |
| `make clean`        | Remove build artifacts & containers |
| `make db:migrate`   | Apply SQL migrations                |
| `make ci`           | Full CI: install → lint → typecheck → test |

## 🔐 Authentication Flow

```
POST /auth/register  →  Creates user record (bcrypt password)
POST /auth/login     →  Verifies password → issues access + refresh tokens
GET  /auth/refresh   →  Verifies refresh token hash → issues new access token
POST /auth/logout    →  Revokes refresh token in Redis
```

- **Access token** – short-lived (15 min, configurable), stored in memory / httpOnly cookie
- **Refresh token** – long-lived (7 days), stored hashed in `refresh_tokens` table + Redis TTL
- **Logout** – refresh token revoked immediately; access token blacklisted via TTL in Redis

## 🗄️ Database Schema

| Table              | Key Columns                                                   |
|--------------------|---------------------------------------------------------------|
| `users`            | id, email, password_hash, role, is_verified, trust_score      |
| `refresh_tokens`   | id, user_id, token_hash, expires_at, revoked                  |
| `communities`      | id, name, category, is_private, member_count, moderator_id    |
| `posts`            | id, community_id, user_id, type, title, content               |
| `housing_listings` | id, landlord_id, title, bedrooms, rent, latitude, longitude   |
| `marketplace_items`| id, seller_id, title, category, price, condition              |
| `transactions`     | id, buyer_id, seller_id, type, amount, status, escrow         |
| `trust_reviews`    | id, reviewer_id, subject_id, subject_type, rating             |
| `lawyers`          | id, user_id, specialization, consultation_fee                 |
| `lawyer_bookings`  | id, lawyer_id, user_id, scheduled_at, type, status            |

## ⚡ API Response Shape

All responses follow a consistent envelope:

### Success
```json
{
  "success": true,
  "data": { ... },
  "meta": { "timestamp": "2025-01-01T00:00:00.000Z", "path": "/users/profile" }
}
```

### Paginated List
```json
{
  "success": true,
  "data": [ ... ],
  "meta": {
    "page": 1,
    "limit": 20,
    "total": 100,
    "totalPages": 5
  }
}
```

### Error
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Request validation failed",
    "details": [ ... ]
  }
}
```

## 🧑‍💻 Development

```bash
# Run type checker
npm run typecheck

# Lint + fix
npm run lint && npm run lint:fix

# Run tests (watch mode)
npm run test:watch

# Build for production
npm run build
npm run start:prod
```

## 🧪 Testing

```
tests/
├── setup.ts            # Jest setup + global mocks
├── unit/
│   ├── utils/          # Unit tests for utility functions
│   └── services/       # Unit tests for service logic
├── integration/
│   └── controllers/    # Controller tests with mocked repositories
└── e2e/
    └── auth.e2e-spec.ts# End-to-end authentication flow
```

Run once before every PR:

```bash
npm ci
npm run lint && npm run typecheck && npm test
```

## 🔍 Code Quality

- **ESLint** – TypeScript config + `security` plugin
- **Prettier** – Enforced formatting on save
- **TypeScript** – Strict mode, no `any` without explicit opt-in
- **Husky** – Pre-commit hook runs `lint-staged` for `src/**/*.ts`
