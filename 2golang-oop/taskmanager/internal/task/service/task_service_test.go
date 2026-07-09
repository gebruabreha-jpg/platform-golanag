package service

import (
	"testing"
)

func TestTaskManager_AddTask(t *testing.T) {
	tm := NewTaskManager()
	task := tm.AddTask("Test task")
	if task.Title != "Test task" {
		t.Errorf("expected title 'Test task', got %s", task.Title)
	}
	if task.Done != false {
		t.Errorf("expected done=false, got %v", task.Done)
	}
	if task.ID != 1 {
		t.Errorf("expected id=1, got %d", task.ID)
	}
}

func TestTaskManager_GetTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task 1")
	task := tm.GetTask(1)
	if task == nil {
		t.Error("expected task, got nil")
	}
	if task.ID != 1 {
		t.Errorf("expected id=1, got %d", task.ID)
	}
	task = tm.GetTask(999)
	if task != nil {
		t.Error("expected nil for nonexistent task")
	}
}

func TestTaskManager_DeleteTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task 1")
	if !tm.DeleteTask(1) {
		t.Error("expected delete to succeed")
	}
	if tm.DeleteTask(999) {
		t.Error("expected delete to fail for nonexistent task")
	}
}

func TestTaskManager_UpdateTask(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task 1")
	newTitle := "Updated"
	newDone := true
	task := tm.UpdateTask(1, &newTitle, &newDone)
	if task == nil {
		t.Error("expected task, got nil")
	}
	if task.Title != "Updated" {
		t.Errorf("expected title 'Updated', got %s", task.Title)
	}
	if task.Done != true {
		t.Errorf("expected done=true, got %v", task.Done)
	}
}

func TestTaskManager_ListTasks(t *testing.T) {
	tm := NewTaskManager()
	tm.AddTask("Task 1")
	tm.AddTask("Task 2")
	tasks := tm.ListTasks()
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}