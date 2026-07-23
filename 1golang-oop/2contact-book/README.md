# Contact Book API
old_code:-
the old Go code reads OS environment variables, not .env files, so even if you create .env it won't be loaded automatically.for old with os.Setenv(...) inside main.go. it worked  and we use config.go  to read from OS.


new_code:-
Use a real .env file (standard for Go projects) using godotenv
Update cmd/server/main.go — add the autoload import at the top, remove the hardcoded os.Setenv(...) and Create .env with:-
PORT=8080
DATABASE_URL=postgresql://neondb_owner:npg_SFsIvO9H5YZn@ep-frosty-voice-alup0zcy-pooler.c-3.eu-central-1.aws.neon.tech/book?sslmode=require&channel_binding=require
APP_ENV=prod


For the CURRENT state:-
cmd/server/main.go - entry point, import  config(PORT and DATABASE_URL), loads .env, connects to DB, then internal/database/db.go - opens PostgreSQL pool and .env - holds secrets.

internal/model/contact.go — struct + JSON tags
internal/repository/contact.go — interface + Postgres SQL
internal/service/contact.go — business logic
internal/handler/contact.go — HTTP layer
cmd/server/main.go — register routes + wire everything(entry point → dependency injection → start server.)

What main.go always does:-
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
main.go creates ALL objects
Each package only creates its OWN objects
Packages never create objects from other packages
This is the Composition Root pattern. It means:


test:-
If you want to test service, you can pass a mock repository — because service doesn't create its own repo
If you want to test handler, you can pass a mock service — because handler doesn't create its own service
main.go is the only place that knows about all packages


The Pattern: 
1,Composition Root:-The Composition Root is a single place (in this case, main.go) where all objects are created and wired together. Every other package only creates its own objects.This is described by Mark Seemann in Dependency Injection in .NET and is a standard pattern in Go, Rust, Java, and C#.

2,Dependency Inversion Principle (SOLID):-High-level modules depend on abstractions, not concretions.
Applied in this project:
  handler depends on ContactServiceInterface (abstraction),
  not *service.ContactService (concrete type).
This means:-
  - The handler doesn't know which service implementation is used
  - You can swap the service with a mock for testing
  - You can create a different service implementation without changing the handler
Before (bad):-
    type ContactHandler struct {
        contacts *service.ContactService  // concrete type
    }
After (good):-
    type ContactHandler struct {
        contacts service.ContactServiceInterface  // interface
    }

3,Inversion of Control (IoC):-
Object creation is controlled externally, not internally.

Applied in this project:
  Packages don't create each other's objects.
  main.go (the composition root) creates ALL objects and wires them together.

Each package only provides constructors for its own type:
    config.Load()          → creates Config
    database.Connect()     → creates *sql.DB
    logger.New()           → creates *slog.Logger
    repository.NewPostgresContactRepository() → creates repo
    service.NewContactService()                → creates service
    handler.NewContactHandler()                → creates handler

No package imports another package to create its objects.
This means:
  - No hidden dependencies
  - No tight coupling between packages
  - Easy to test each package in isolation
