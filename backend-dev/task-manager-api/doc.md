I'll create a Go-based task manager REST API with an industry-grade project structure that can be easily extended. This will give you a solid foundation to learn from.

```go
// Project Structure
// task-manager/
// ├── cmd/
// │   └── api/
// │       └── main.go
// ├── internal/
// │   ├── config/
// │   │   └── config.go
// │   ├── domain/
// │   │   └── task.go
// │   ├── repository/
// │   │   ├── interface.go
// │   │   └── postgres/
// │   │       └── task_repository.go
// │   ├── service/
// │   │   └── task_service.go
// │   ├── handler/
// │   │   └── task_handler.go
// │   ├── middleware/
// │   │   └── logging.go
// │   └── errs/
// │       └── errors.go
// ├── pkg/
// │   └── logger/
// │       └── logger.go
// ├── go.mod
// ├── go.sum
// ├── Makefile
// ├── .env.example
// ├── docker-compose.yml
// └── README.md

// Let's implement the core files:

// 1. First, let's define the module and dependencies in go.mod
// go.mod
module github.com/yourusername/task-manager

go 1.22

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.4.0
	github.com/jackc/pgx/v5 v5.5.0
	github.com/joho/godotenv v1.5.1
	github.com/rs/zerolog v1.31.0
	github.com/stretchr/testify v1.8.4
)

// 2. Main entry point
// cmd/api/main.go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/task-manager/internal/config"
	"github.com/yourusername/task-manager/internal/handler"
	"github.com/yourusername/task-manager/internal/middleware"
	"github.com/yourusername/task-manager/internal/repository/postgres"
	"github.com/yourusername/task-manager/internal/service"
	"github.com/yourusername/task-manager/pkg/logger"
)

func main() {
	// Initialize logger
	log := logger.New()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to database
	db, err := postgres.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	// Initialize repository
	taskRepo := postgres.NewTaskRepository(db)

	// Initialize service
	taskService := service.NewTaskService(taskRepo)

	// Initialize handlers
	taskHandler := handler.NewTaskHandler(taskService)

	// Initialize router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logger(log))

	// Register routes
	v1 := router.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", taskHandler.ListTasks)
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("/:id", taskHandler.GetTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// Create HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Info().Msgf("Starting server on port %d", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting")
}

// 3. Configuration
// internal/config/config.go
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        int
	Environment string
	DatabaseURL string
	LogLevel    string
}

func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid port: %w", err)
	}

	return &Config{
		Port:        port,
		Environment: getEnv("ENVIRONMENT", "development"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// 4. Domain model
// internal/domain/task.go
package domain

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone       TaskStatus = "DONE"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateTaskInput struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type UpdateTaskInput struct {
	Title       *string     `json:"title,omitempty"`
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty"`
	DueDate     *time.Time  `json:"due_date,omitempty"`
}

// 5. Repository interface
// internal/repository/interface.go
package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/yourusername/task-manager/internal/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	List(ctx context.Context) ([]*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// 6. Postgres implementation
// internal/repository/postgres/task_repository.go
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/yourusername/task-manager/internal/domain"
	"github.com/yourusername/task-manager/internal/errs"
)

type TaskRepository struct {
	db *sql.DB
}

func NewConnection(connectionString string) (*sql.DB, error) {
	config, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDB(*config)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	query := `
		INSERT INTO tasks (id, title, description, status, due_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.DueDate,
		task.CreatedAt,
		task.UpdatedAt,
	)

	return err
}

func (r *TaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	query := `
		SELECT id, title, description, status, due_date, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	var task domain.Task
	var dueDate sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&dueDate,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}

	if dueDate.Valid {
		task.DueDate = &dueDate.Time
	}

	return &task, nil
}

func (r *TaskRepository) List(ctx context.Context) ([]*domain.Task, error) {
	query := `
		SELECT id, title, description, status, due_date, created_at, updated_at
		FROM tasks
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task

	for rows.Next() {
		var task domain.Task
		var dueDate sql.NullTime

		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&dueDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}

		if dueDate.Valid {
			task.DueDate = &dueDate.Time
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *domain.Task) error {
	query := `
		UPDATE tasks
		SET title = $1, description = $2, status = $3, due_date = $4, updated_at = $5
		WHERE id = $6
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Status,
		task.DueDate,
		task.UpdatedAt,
		task.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *TaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

// 7. Service layer
// internal/service/task_service.go
package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yourusername/task-manager/internal/domain"
	"github.com/yourusername/task-manager/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, input domain.CreateTaskInput) (*domain.Task, error) {
	now := time.Now()

	// Set default status if not provided
	if input.Status == "" {
		input.Status = domain.TaskStatusTodo
	}

	task := &domain.Task{
		ID:          uuid.New(),
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		DueDate:     input.DueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskService) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	return s.repo.List(ctx)
}

func (s *TaskService) UpdateTask(ctx context.Context, id uuid.UUID, input domain.UpdateTaskInput) (*domain.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if input.Title != nil {
		task.Title = *input.Title
	}

	if input.Description != nil {
		task.Description = *input.Description
	}

	if input.Status != nil {
		task.Status = *input.Status
	}

	if input.DueDate != nil {
		task.DueDate = input.DueDate
	}

	task.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// 8. HTTP handlers
// internal/handler/task_handler.go
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourusername/task-manager/internal/domain"
	"github.com/yourusername/task-manager/internal/errs"
	"github.com/yourusername/task-manager/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input domain.CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.CreateTask(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.service.GetTask(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks, err := h.service.ListTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var input domain.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.UpdateTask(c.Request.Context(), id, input)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = h.service.DeleteTask(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// 9. Custom errors
// internal/errs/errors.go
package errs

import "errors"

var (
	ErrNotFound = errors.New("resource not found")
)

// 10. Middleware
// internal/middleware/logging.go
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Logger(log zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		if raw != "" {
			path = path + "?" + raw
		}

		// Process request
		c.Next()

		// Log after request
		end := time.Now()
		latency := end.Sub(start)

		log.Info().
			Str("method", c.Request.Method).
			Str("path", path).
			Int("status", c.Writer.Status()).
			Dur("latency", latency).
			Str("client_ip", c.ClientIP()).
			Msg("API request")
	}
}

// 11. Logger package
// pkg/logger/logger.go
package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	return zerolog.New(output).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(zerolog.InfoLevel)
}

// 12. Docker Compose for local development
// docker-compose.yml
version: '3.8'

services:
  postgres:
    image: postgres:16
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: taskdb
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./migrations/init.sql:/docker-entrypoint-initdb.d/init.sql

  api:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - ENVIRONMENT=development
      - DATABASE_URL=postgres://postgres:postgres@postgres:5432/taskdb?sslmode=disable
      - LOG_LEVEL=debug
    depends_on:
      - postgres
    volumes:
      - .:/app

volumes:
  postgres_data:

// 13. SQL initialization script
// migrations/init.sql
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL,
    due_date TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
CREATE INDEX IF NOT EXISTS idx_tasks_created_at ON tasks(created_at);

// 14. Makefile
# Makefile
.PHONY: build run test lint clean docker-up docker-down

build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

clean:
	rm -rf bin/

```

This Go task manager REST API follows industry best practices with a clean architecture that separates concerns and allows for easy extension. Here's a breakdown of the key components:

### Project Structure

- **Clean Architecture**: Separated into domain, repository, service, and handler layers
- **Dependency Injection**: Components are wired together in the main function
- **Configuration Management**: Environment-based configuration with sensible defaults
- **Error Handling**: Custom error types with appropriate HTTP responses
- **Middleware**: Logging middleware for request tracking

### Technical Features

- **RESTful API**: Full CRUD operations for tasks
- **PostgreSQL Database**: Structured data storage with appropriate indexes
- **Graceful Shutdown**: Handles OS signals to shut down cleanly
- **Input Validation**: Request payload validation with appropriate error messages
- **Containerization**: Docker setup for local development
- **Modularity**: Well-defined interfaces allowing for different implementations

### Getting Started

1. Clone the repository
2. Run `make docker-up` to start PostgreSQL
3. Run `make run` to start the API server
4. Access the API at http://localhost:8080

### API Endpoints

- `GET /api/v1/tasks` - List all tasks
- `POST /api/v1/tasks` - Create a new task
- `GET /api/v1/tasks/:id` - Get a specific task
- `PUT /api/v1/tasks/:id` - Update a task
- `DELETE /api/v1/tasks/:id` - Delete a task
- `GET /health` - Health check endpoint

This structure is designed to scale well as requirements grow. You can easily add new features by creating new domain models, repositories, services, and handlers without modifying existing code.

Would you like me to explain any specific part of the implementation in more detail?
