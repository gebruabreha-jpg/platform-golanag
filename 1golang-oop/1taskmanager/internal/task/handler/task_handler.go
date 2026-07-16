package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"task-manager-api/internal/task/service"
	"task-manager-api/pkg/response"
)

type TaskHandler struct {
	tm *service.TaskManager
}

func NewTaskHandler(tm *service.TaskManager) *TaskHandler {
	return &TaskHandler{tm: tm}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	response.OK(c, http.StatusOK, h.tm.ListTasks())
}

type TaskCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input TaskCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	task := h.tm.AddTask(input.Title, input.Description)
	response.OK(c, http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	idStr := c.Param("id")
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	task := h.tm.GetTask(taskID)
	if task != nil {
		response.OK(c, http.StatusOK, task)
		return
	}
	response.Fail(c, http.StatusNotFound, "Task not found")
}

type TaskUpdate struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var input TaskUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	task := h.tm.UpdateTask(taskID, input.Title, input.Done)
	if task != nil {
		response.OK(c, http.StatusOK, task)
		return
	}
	response.Fail(c, http.StatusNotFound, "Task not found")
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if h.tm.DeleteTask(taskID) {
		c.Status(http.StatusNoContent)
		return
	}
	response.Fail(c, http.StatusNotFound, "Task not found")
}
