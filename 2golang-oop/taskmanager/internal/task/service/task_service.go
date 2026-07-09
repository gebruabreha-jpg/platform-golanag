package service

import "task-manager-api/internal/task/model"

type TaskManager struct {
	tasks   []model.Task
	counter int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:   []model.Task{},
		counter: 1,
	}
}

func (tm *TaskManager) AddTask(title string) model.Task {
	task := model.Task{ID: tm.counter, Title: title, Done: false}
	tm.counter++
	tm.tasks = append(tm.tasks, task)
	return task
}

func (tm *TaskManager) GetTask(taskID int) *model.Task {
	for i := range tm.tasks {
		if tm.tasks[i].ID == taskID {
			return &tm.tasks[i]
		}
	}
	return nil
}

func (tm *TaskManager) DeleteTask(taskID int) bool {
	for i, task := range tm.tasks {
		if task.ID == taskID {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (tm *TaskManager) UpdateTask(taskID int, title *string, done *bool) *model.Task {
	task := tm.GetTask(taskID)
	if task != nil {
		if title != nil {
			task.Title = *title
		}
		if done != nil {
			task.Done = *done
		}
	}
	return task
}

func (tm *TaskManager) ListTasks() []model.Task {
	return tm.tasks
}