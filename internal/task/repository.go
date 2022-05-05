package task

import (
	"context"
	"database/sql"
)

type Repository interface {
	Get(ctx context.Context, id string) (Task, error)
	Count(ctx context.Context) (int, error)
	// GetAll TODO: add pagination
	GetAll(ctx context.Context) ([]Task, error)
	Update(ctx context.Context, task Task) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) Get(ctx context.Context, id string) (Task, error) {
	var task Task
	err := r.db.QueryRowContext(ctx, "SELECT * FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.UpdatedAt)

	return task, err
}

func (r repository) Count(ctx context.Context) (int, error) {
	panic("implement me")
}

func (r repository) GetAll(ctx context.Context) ([]Task, error) {
	panic("implement me")
}

func (r repository) Update(ctx context.Context, task Task) error {
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
