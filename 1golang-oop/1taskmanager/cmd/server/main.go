package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"task-manager-api/internal/config"
	"task-manager-api/internal/task/handler"
	"task-manager-api/internal/task/repository"
	"task-manager-api/internal/task/service"
	"task-manager-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	log := logger.New()

	var repo repository.TaskRepository
	if cfg.UsePostgres() {
		pg, err := repository.NewPostgresTaskRepository(cfg.DatabaseURL)
		if err != nil {
			log.Error("failed to connect to database", "error", err)
			os.Exit(1)
		}
		defer pg.Close()
		repo = pg
		log.Info("using Postgres repository")
	} else {
		repo = repository.NewInMemoryTaskRepository()
		log.Info("using in-memory repository (set DATABASE_URL for Postgres)")
	}

	tm := service.NewTaskManager(repo)
	h := handler.NewTaskHandler(tm)

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/tasks", h.GetTasks)
	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks/:id", h.GetTask)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)

	srv := &http.Server{
		Addr:    cfg.Addr(),
		Handler: r,
	}

	go func() {
		log.Info("server listening", "addr", cfg.Addr())
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("forced shutdown", "error", err)
		os.Exit(1)
	}
	log.Info("server stopped")
}
