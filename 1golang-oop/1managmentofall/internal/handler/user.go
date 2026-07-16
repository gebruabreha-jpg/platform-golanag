package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"task-manager-api/internal/model"
	"task-manager-api/internal/service"
	"task-manager-api/pkg/response"
)

type UserHandler struct {
	users *service.UserService
}

func NewUserHandler(users *service.UserService) *UserHandler {
	return &UserHandler{users: users}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.users.Register(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			response.Fail(c, http.StatusConflict, "email already registered")
			return
		}
		response.Fail(c, http.StatusInternalServerError, "internal error")
		return
	}
	response.OK(c, http.StatusCreated, toUserResponse(user))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.users.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.Fail(c, http.StatusUnauthorized, "invalid credentials")
		return
	}
	response.OK(c, http.StatusOK, toUserResponse(*user))
}

func toUserResponse(u model.User) UserResponse {
	return UserResponse{ID: u.ID, Email: u.Email, CreatedAt: u.CreatedAt}
}
