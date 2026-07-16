# 2cli-task-manager

A small command-line task manager written in Go.

## Features
- Add, list, complete, and delete tasks from the terminal
- Tasks persisted to a local JSON file (`~/.taskmanager/tasks.json`)
- Simple, dependency-free CLI built on the standard library

## Build
```
go build -o taskmgr ./cmd/cli
```

## Usage
```
taskmgr add "Write the report"          # create a task
taskmgr list                            # list all tasks
taskmgr done <id>                      # mark a task complete
taskmgr delete <id>                    # delete a task
```

## Layout
```
cmd/cli        entrypoint + flag parsing
internal/task  model / repository / service (same clean layering as the API)
internal/config environment + file paths
```
