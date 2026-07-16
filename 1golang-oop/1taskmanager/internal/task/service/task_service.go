package service


import(
	"sync"
	"task-manager-api/internal/task/model"
)

type TaskManager struct {
	mu   sync.Mutex
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
	tm.mu.Lock()
	defer tm.mu.Unlock()
	task := model.Task{ID: tm.counter, Title: title, Done: false}
	tm.counter++
	tm.tasks = append(tm.tasks, task)
	return task
}

func (tm *TaskManager) GetTask(taskID int) *model.Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	for i := range tm.tasks {
		if tm.tasks[i].ID == taskID {
			return &tm.tasks[i]
		}
	}
	return nil
}

func (tm *TaskManager) DeleteTask(taskID int) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	for i, task := range tm.tasks {
		if task.ID == taskID {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (tm *TaskManager) UpdateTask(taskID int, title *string, done *bool) *model.Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	for i := range tm.tasks {
		if tm.tasks[i].ID == taskID {
			if title != nil {
				tm.tasks[i].Title = *title
			}
			if done != nil {
				tm.tasks[i].Done = *done
			}
			return &tm.tasks[i]
		}
	}
	return nil 
}

func (tm *TaskManager) ListTasks() []model.Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	return tm.tasks
}