package main

import (
	"crypto/rand"
	"encoding/hex"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"contact-book-api/internal/config"
	"contact-book-api/internal/database"
	"contact-book-api/internal/handler"
	"contact-book-api/internal/repository"
	"contact-book-api/internal/service"
	"contact-book-api/pkg/logger"
	"contact-book-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		logger.New().Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	log := logger.New()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	contactRepo := repository.NewPostgresContactRepository(db)
	contactSvc := service.NewContactService(contactRepo)
	contactHandler := handler.NewContactHandler(contactSvc)

	r := gin.Default()

	r.Use(traceIDMiddleware())

	r.GET("/healthz", func(c *gin.Context) {
		response.OK(c.Writer, http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/readyz", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			response.Fail(c.Writer, http.StatusServiceUnavailable, "database not ready")
			return
		}
		response.OK(c.Writer, http.StatusOK, gin.H{"status": "ready"})
	})

	r.GET("/metrics", func(c *gin.Context) {
		h := promhttp.Handler()
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/v1/contacts", contactHandler.Create)
	r.GET("/v1/contacts", contactHandler.GetAll)
	r.GET("/v1/contacts/:id", contactHandler.GetByID)
	r.PUT("/v1/contacts/:id", contactHandler.Update)
	r.DELETE("/v1/contacts/:id", contactHandler.Delete)

	srv := &http.Server{
		Addr:    cfg.Addr(),
		Handler: r,
	}

	go func() {
		log.Info("server listening", "addr", cfg.Addr())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutting down...")

	timeout := time.Duration(cfg.ShutdownTimeoutSeconds) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("forced shutdown", "error", err)
		os.Exit(1)
	}
	log.Info("server stopped")
}

func traceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Request-ID")
		if traceID == "" {
			bytes := make([]byte, 8)
			if _, err := rand.Read(bytes); err == nil {
				traceID = hex.EncodeToString(bytes)
			}
		}
		c.Set("trace_id", traceID)
		c.Writer.Header().Set("X-Request-ID", traceID)
		c.Next()
	}
}