.PHONY: help dev build test clean docker-up docker-down logs migrate

help:
	@echo "ConnectMe Development Commands"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  dev            Start all services in development mode"
	@echo "  build          Build all Docker images"
	@echo "  test           Run tests for all services"
	@echo "  lint           Run linters for all services"
	@echo "  migrate        Run database migrations"
	@echo "  docker-up      Start services with Docker Compose"
	@echo "  docker-down    Stop all services"
	@echo "  logs           Show logs from all services"
	@echo "  clean          Remove build artifacts and containers"
	@echo "  frontend-dev   Start frontend dev server only"
	@echo "  backend-dev    Start backend dev server only"
	@echo "  ai-dev         Start AI service dev server only"

dev:
	docker-compose up -d

build:
	docker-compose build

test:
	@echo "Running backend tests..."
	cd backend && go test ./...
	@echo ""
	@echo "Running AI service tests..."
	cd ai-service && pytest -v
	@echo ""
	@echo "Running frontend tests..."
	cd frontend && npm test

lint:
	@echo "Linting backend..."
	cd backend && golangci-lint run
	@echo ""
	@echo "Linting AI service..."
	cd ai-service && black --check app/ && flake8 app/
	@echo ""
	@echo "Linting frontend..."
	cd frontend && npm run lint

migrate:
	cd backend && goose -dir migrations postgres $(DATABASE_URL) up

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

logs:
	docker-compose logs -f

clean:
	docker-compose down -v
	rm -rf backend/dist
	rm -rf frontend/.next
	rm -rf ai-service/__pycache__
	rm -rf ai-service/*.pyc

frontend-dev:
	cd frontend && npm run dev

backend-dev:
	cd backend && go run cmd/server/main.go

ai-dev:
	cd ai-service && uvicorn app.main:app --reload

# Database utilities
db-shell:
	docker-compose exec postgres psql -U connectme_user -d connectme

db-backup:
	docker-compose exec postgres pg_dump -U connectme_user connectme > backup_$(shell date +%Y%m%d_%H%M%S).sql

# Production deploy (example)
deploy-prod:
	@echo "Deploying to production..."
	# Add your deployment commands here
