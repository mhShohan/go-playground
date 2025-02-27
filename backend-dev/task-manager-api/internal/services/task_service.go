package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mhShohan/go-playground/task-manager-api/task-manager/internal/domain"
	"github.com/mhShohan/go-playground/task-manager-api/task-manager/internal/repository"
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
