# Contact Book API

## Run
```bash
go mod tidy
go run ./cmd/server
```

## Test
```bash
go test ./...
```

## Endpoints
- `GET /healthz` - health check
- `POST /contacts` - create contact
- `GET /contacts/:id` - get contact by ID

option1:-
the current Go code reads OS environment variables, not .env files, so even if you create .env it won't be loaded automatically.
for now with os.Setenv(...) inside main.go. it will work  so Just run it:
go mod tidy
go run ./cmd/server

option2:-
 Use a real .env file (standard for Go projects)
 install godotenv:-
go get github.com/joho/godotenv
go mod tidy

Update cmd/server/main.go — add the autoload import at the top, remove the hardcoded os.Setenv(...)

Create .env in 2contact-book/:

PORT=8080
DATABASE_URL=postgresql://neondb_owner:npg_SFsIvO9H5YZn@ep-frosty-voice-alup0zcy-pooler.c-3.eu-central-1.aws.neon.tech/book?sslmode=require&channel_binding=require
APP_ENV=prod

 Run:
go run ./cmd/server


For the CURRENT state (DB connectivity only):-
cmd/server/main.go - entry point, loads env, connects to DB
internal/config/config.go - reads PORT and DATABASE_URL
internal/database/db.go - opens PostgreSQL pool
.env - holds secrets
go.mod - dependencies


You don't need internal/config/config.go for a tiny connectivity check. It's a common Go pattern, but optional.Simpler version of main.go — no config package needed.

db.go only makes sense later when multiple files need to reuse the connection logic. For now, delete it and keep main.go self-contained.db.go only makes sense later when multiple files need to reuse the connection logic. For now, delete it and keep main.go self-contained.


internal/model/contact.go — struct + JSON tags
internal/repository/contact.go — interface + Postgres SQL
internal/service/contact.go — business logic
internal/handler/contact.go — HTTP layer
cmd/server/main.go — register routes + wire everything(entry point → dependency injection → start server.)
What main.go always does:
Load config (.env → config)
Initialize dependencies (DB pool, repository, service, handler)
Register routes
Start HTTP server
Graceful shutdown
What changes as features grow:
More imports — new services, handlers, repos
More initialization — NewXService(), NewXHandler()
More routes — r.POST(...), r.GET(...)
More middleware — auth, logging, CORS