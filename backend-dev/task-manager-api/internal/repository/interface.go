package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhShohan/go-playground/task-manager-api/task-manager/internal/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	List(ctx context.Context) ([]*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}
