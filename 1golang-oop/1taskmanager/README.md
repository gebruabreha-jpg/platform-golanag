run:
go mod tidy
go build ./...
go vet ./...

go clean
go mod tidy
gofmt -w .
golangci-lint run
go build ./cmd/myapp
go run main.go
go test ./...


 Go best practices:-
Structure:
- cmd/server/main.go - entry point
- internal/task/model/task.go - Task struct
- internal/task/service/task_service.go - TaskManager logic
- internal/task/handler/task_handler.go - HTTP handlers

Run: go run ./cmd/server
Test: go test ./... -v

✅ Short answer (very important)
1. Why struct fields are capital?
👉 Because Go only exports (makes accessible) fields that start with a capital letter.

2. Why JSON tags are used?
👉 To control how the field name looks in JSON output (usually lowercase like id, title).

🧠 Conclusion (super short)
Capital fields → required so Go can access/serialize them (public)
JSON tags → control API output format (clean JSON for frontend)




model.py (Task struct) → Go struct
logic.py (TaskManager) → Go with slice of Task structs and methods
api.py (FastAPI) → Go with net/http or a framework like Gin/Echo/Fiber
Tests → Go testing package
Key differences Go vs Python:
No classes - use structs and methods
No optional parameters - use pointers or separate structs
Static typing with explicit error handling



Go (Golang) does not have classes in the same way as Python, Java, or C++.
Instead, Go uses a simpler model based on:
1. Structs (data)
type Task struct {
    ID    int
    Title string
    Done  bool
}
2. Methods (behavior attached to structs)

You can attach functions to a struct using a receiver:

func (t Task) IsDone() bool {
    return t.Done
}

func (t *Task) MarkDone() {
    t.Done = true
}
Key idea
A struct ≈ class fields
methods on structs ≈ class methods
but there is no inheritance, no class keyword
Instead of inheritance, Go uses:
composition (embedding structs)
type Base struct {
    ID int
}

type Task struct {
    Base
    Title string
}











Even with Gin, you still use net/http because Gin is built on top of Go’s standard HTTP system.
⚙️ What Gin does
Gin is just a wrapper/framework around Go’s built-in HTTP server.
So under the hood:
net/http (core engine)
        ↓
Gin (helper layer)
📦 Why we still need net/http
Because it gives you standard HTTP constants and types, like:

1. HTTP status codes
Without net/http, you’d write numbers:
c.JSON(200, data)
With net/http:
c.JSON(http.StatusOK, data)
👉 clearer + readable