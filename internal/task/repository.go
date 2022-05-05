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
	Create(ctx context.Context, task Task) error
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
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM tasks").Scan(&count)

	return count, err
}

func (r repository) GetAll(ctx context.Context) ([]Task, error) {
	var tasks []Task
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	err = rows.Err()

	return tasks, err
}

func (r repository) Create(ctx context.Context, task Task) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tasks VALUES ($1, $2, $3, $4, $5)", task.ID, task.Name, task.Description, task.CreatedAt, task.UpdatedAt)

	return err
}

func (r repository) Update(ctx context.Context, task Task) error {
	_, err := r.db.ExecContext(ctx, "UPDATE tasks SET name = $1, description = $2, updated_at = $3 WHERE id = $4", task.Name, task.Description, task.UpdatedAt, task.ID)
	return err
}

func (r repository) Delete(ctx context.Context, id string) error {
	task, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, "DELETE FROM tasks WHERE id = $1", task.ID)

	return err
}
