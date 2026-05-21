.PHONY: help dev start start:prod build test lint typecheck docker-up docker-down logs clean db:migrate ci

NODE_DIR := node_modules/.bin
DC := docker compose

help:
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"
	@echo "   Booking TypeScript – Development Commands        ┃"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫"
	@echo "" && \
	echo "  make dev          Start all services (Docker Compose)" && \
	echo "  make docker-up    Start infra only (postgres + redis)" && \
	echo "  make docker-down  Stop all services" && \
	echo "  make logs         Tail Docker Compose logs" && \
	echo "" && \
	echo "  make build        Build TypeScript → dist/" && \
	echo "  make start        Run dist/index.js (prod mode)" && \
	echo "  make dev:nodemon  Dev server with hot reload (nodemon)" && \
	echo "" && \
	echo "  make test         Run Jest test suite (CI)" && \
	echo "  make test:watch   Run Jest in watch mode" && \
	echo "  make lint         Run ESLint" && \
	echo "  make lint:fix     Run ESLint with --fix" && \
	echo "  make typecheck    Run tsc --noEmit" && \
	echo "" && \
	echo "  make db:migrate   Apply SQL migrations to Postgres" && \
	echo "  make clean        Remove dist/ + Docker volumes" && \
	echo "" && \
	echo "  make ci           Full CI: install → lint → typecheck → test"

# ── Docker ──────────────────────────────────────────────────────────────────────
dev: docker-up

docker-up:
	$(DC) up -d postgres redis pgadmin

docker-down:
	$(DC) down

logs:
	$(DC) logs -f --tail=100

clean:
	rm -rf dist coverage .nyc_output
	$(DC) down -v --remove-orphans

# ── Build & start ───────────────────────────────────────────────────────────────
build:
	@npm run build

start: build
	@npm run start:prod

# ── Testing ─────────────────────────────────────────────────────────────────────
test:
	@jest --coverage

test:watch:
	@jest --watch

# ── Lint & type check ──────────────────────────────────────────────────────────
lint:
	@echo "→ Running ESLint…" && $(NODE_DIR)/eslint src --ext .ts --cache

lint:fix:
	@echo "→ Running ESLint with --fix…" && $(NODE_DIR)/eslint src --ext .ts --cache --fix

typecheck:
	@echo "→ Running TypeScript compiler…" && tsc --noEmit -p tsconfig.json

# ── Database ───────────────────────────────────────────────────────────────────
db:migrate:
	@echo "→ Running migrations…" && \
	psql $$DATABASE_URL -f src/database/migrations/001_create_users_table.sql && \
	psql $$DATABASE_URL -f src/database/migrations/002_create_refresh_tokens_table.sql && \
	psql $$DATABASE_URL -f src/database/migrations/003_create_communities_table.sql && \
	psql $$DATABASE_URL -f src/database/migrations/004_create_housing_listings_table.sql && \
	psql $$DATABASE_URL -f src/database/migrations/005_create_marketplace_items_table.sql && \
	psql $$DATABASE_URL -f src/database/migrations/006_create_services_table.sql && \
	psql $$DATABASE_URL -f src/database/migrations/007_create_transactions_table.sql && \
	psql $$DATABASE_URL -f src/database/migrations/008_create_posts_table.sql && \
	psql $$DATABASE_URL -f src/database/migrations/009_create_reviews_table.sql

# ── Full CI pipeline ────────────────────────────────────────────────────────────
ci: install lint typecheck test
install:
	npm ci
