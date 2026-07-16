package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"cli-task-manager/internal/config"
	"cli-task-manager/internal/task/repository"
	"cli-task-manager/internal/task/service"
)

func main() {
	cfg := config.Load()
	repo := repository.NewFileTaskRepository(cfg.StorePath)
	tm := service.NewTaskManager(repo)

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		fs := flag.NewFlagSet("add", flag.ExitOnError)
		fs.Parse(os.Args[2:])
		if fs.NArg() == 0 {
			fail("add requires a title, e.g. `taskmgr add \"buy milk\"`")
		}
		title := joinArgs(fs.Args())
		task, err := tm.Add(title)
		if err != nil {
			fail(err.Error())
		}
		fmt.Printf("added task #%d: %s\n", task.ID, task.Title)

	case "list":
		tasks, err := tm.List()
		if err != nil {
			fail(err.Error())
		}
		if len(tasks) == 0 {
			fmt.Println("no tasks yet")
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tDONE\tTITLE")
		for _, t := range tasks {
			done := " "
			if t.Done {
				done = "x"
			}
			fmt.Fprintf(w, "%d\t[%s]\t%s\n", t.ID, done, t.Title)
		}
		w.Flush()

	case "done":
		id := firstID(os.Args[2:], "done")
		task, err := tm.Complete(id)
		if err != nil {
			fail(err.Error())
		}
		fmt.Printf("completed task #%d: %s\n", task.ID, task.Title)

	case "delete":
		id := firstID(os.Args[2:], "delete")
		if err := tm.Delete(id); err != nil {
			fail(err.Error())
		}
		fmt.Printf("deleted task #%d\n", id)

	case "help", "-h", "--help":
		usage()

	default:
		fail(fmt.Sprintf("unknown command %q", os.Args[1]))
	}
}

func firstID(args []string, cmd string) int {
	if len(args) == 0 {
		fail(fmt.Sprintf("%s requires a task id, e.g. `taskmgr %s 1`", cmd, cmd))
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fail(fmt.Sprintf("invalid id %q", args[0]))
	}
	return id
}

func joinArgs(args []string) string {
	out := ""
	for i, a := range args {
		if i > 0 {
			out += " "
		}
		out += a
	}
	return out
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, "error:", msg)
	os.Exit(1)
}

func usage() {
	fmt.Print(`taskmgr - a tiny CLI task manager

Usage:
  taskmgr add "task title"     add a new task
  taskmgr list                 list all tasks
  taskmgr done <id>           mark a task complete
  taskmgr delete <id>         delete a task
  taskmgr help                show this help

Tasks are stored in TASK_STORE (default ~/.taskmanager/tasks.json).
`)
}
