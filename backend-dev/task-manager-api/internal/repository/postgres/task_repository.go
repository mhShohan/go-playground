package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/mhShohan/go-playground/task-manager-api/task-manager/internal/domain"
	"github.com/mhShohan/go-playground/task-manager-api/task-manager/internal/errors"
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
