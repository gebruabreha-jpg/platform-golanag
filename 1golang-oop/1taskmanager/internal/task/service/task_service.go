package service

import (
	"task-manager-api/internal/task/model"
	"task-manager-api/internal/task/repository"
)

type TaskManager struct {
	repo repository.TaskRepository
}

func NewTaskManager(repo repository.TaskRepository) *TaskManager {
	return &TaskManager{repo: repo}
}

func (tm *TaskManager) AddTask(title string) model.Task {
	return tm.repo.Add(title)
}

func (tm *TaskManager) GetTask(taskID int) *model.Task {
	return tm.repo.Get(taskID)
}

func (tm *TaskManager) DeleteTask(taskID int) bool {
	return tm.repo.Delete(taskID)
}

func (tm *TaskManager) UpdateTask(taskID int, title *string, done *bool) *model.Task {
	return tm.repo.Update(taskID, title, done)
}

func (tm *TaskManager) ListTasks() []model.Task {
	return tm.repo.List()
}
