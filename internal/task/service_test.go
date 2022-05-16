package task

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_service_CRUD(t *testing.T) {
	s := NewService(&mockRepository{})
	ctx := context.Background()

	// first count
	count, _ := s.Count(ctx)
	assert.Equal(t, 0, count)

	// successful creating
	task, err := s.Create(ctx, CreateTaskRequest{
		Name:        "Test",
		Description: "description",
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, task.ID)
	id := task.ID
	assert.Equal(t, "test", task.Name)
	assert.Equal(t, "description", task.Description)
	assert.NotEmpty(t, task.UpdatedAt)
	assert.NotEmpty(t, task.CreatedAt)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)
	
	// validation error
	_, err = s.Create(ctx, CreateTaskRequest{
		Name:        "",
		Description: "",
	})
	assert.NotNil(t, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)


	_, _ = s.Create(ctx, CreateTaskRequest{
		Name:        "test 2",
		Description: "description 2",
	})

	// update
	task, err = s.Update(ctx, id, UpdateTaskRequest{Name: "test updated"})
	assert.Nil(t, err)
	assert.Equal(t, "test updated", task.Name)
	_, err = s.Update(ctx, "none", UpdateTaskRequest{Name: "test updated"})
	assert.NotNil(t, err)

	// validation error in updated
	_, err = s.Update(ctx, id, UpdateTaskRequest{
		Name:        "",
		Description: "",
	})
	assert.NotNil(t, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 2, count)

	// get
	_, err = s.Get(ctx, "none")
	assert.NotNil(t, err)
	task, err = s.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, "test updated", task.Name)
	assert.Equal(t, id, task.ID)

	// get all
	tasks, _ := s.GetAll(ctx)
	assert.Equal(t, 2, len(tasks))

	// delete
	_, err = s.Delete(ctx, "none")
	assert.NotNil(t, err)
	task, err = s.Delete(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, id, task.ID)
	count, err = s.Count(ctx)
	assert.Equal(t, 1, count)
}

type mockRepository struct {
	items []Task
}

func (m mockRepository) Get(ctx context.Context, id string) (Task, error) {
	for _, item := range m.items {
		if item.ID == id {
			return item, nil
		}
	}

	return Task{}, sql.ErrNoRows
}

func (m mockRepository) Count(ctx context.Context) (int, error) {
	return len(m.items), nil
}

func (m mockRepository) GetAll(ctx context.Context) ([]Task, error) {
	return m.items, nil
}

func (m *mockRepository) Create(ctx context.Context, task Task) error {
	m.items = append(m.items, task)

	return nil
}

func (m *mockRepository) Update(ctx context.Context, task Task) error {
	for i, item := range m.items {
		if item.ID == task.ID {
			m.items[i] = task
			break
		}
	}

	return nil
}

func (m *mockRepository) Delete(ctx context.Context, id string) error {
	for i, item := range m.items {
		if item.ID == id {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			break
		}
	}

	return nil
}

