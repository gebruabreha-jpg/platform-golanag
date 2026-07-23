# URL Shortener:-
CLI Handler → Service → Repository → Model
Layer	Package	Responsibility
CLI Entry	cmd/cli	Parse CLI args, call service, print output
Business Logic	internal/service	Validate URLs, generate short codes
Data Access	internal/repository	File-based storage (JSON)
Domain Model	internal/model	URL struct with code, long URL, timestamp
Config	internal/config	Load settings from environment variables

Project Structure
3URL-Shortener/
├── cmd/cli/
│   └── main.go              # CLI entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Environment-based configuration
│   ├── model/
│   │   └── url.go           # URL domain model
│   ├── repository/
│   │   └── repository.go    # File-based URL storage
│   └── service/
│       └── task_service.go  # URL shortening logic
├── go.mod
└── README.md


Linting
A .golangci.yml configuration is recommended for this project. Example:

run:
  timeout: 120s

linters:
  enable:
    - govet
    - staticcheck
    - gosimple
    - unused
    - gofmt
    - goimports
    - gci