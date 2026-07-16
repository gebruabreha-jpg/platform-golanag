package main

import (
	"task-manager-api/internal/task/handler"
	"task-manager-api/internal/task/service"

	"github.com/gin-gonic/gin"
)

func main() {
	tm := service.NewTaskManager()
	h := handler.NewTaskHandler(tm)

	r := gin.Default()

	r.GET("/tasks", h.GetTasks)
	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks/:id", h.GetTask)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)

	r.Run(":8080")
}
