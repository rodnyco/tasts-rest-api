package task

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rodnyco/tasks-rest-api/internal/config"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)


func TestRepository(t *testing.T) {
	cfg := config.TestConfig()

	driverName := "postgres"
	dataSource := cfg.DSN

	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatalln(err)
	}
	repo := NewRepository(db)

	ctx := context.Background()

	// count 0
	count, err := repo.Count(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, count)

	// create
	err = repo.Create(ctx, Task{
		ID:          "testID",
		Name:        "Test Task",
		Description: "Description of the task",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	assert.Nil(t, err)
	count1, _ := repo.Count(ctx)
	assert.Equal(t, 1, count1)

	// get
	task, err := repo.Get(ctx, "testID")
	assert.Nil(t, err)
	assert.Equal(t, "Test Task", task.Name)
	_, err = repo.Get(ctx, "undefined id")
	assert.Equal(t, sql.ErrNoRows, err)

	// update
	err = repo.Update(ctx, Task{
		ID:          "testID",
		Name:        "Test Task updated",
		Description: "",
		UpdatedAt:   time.Time{},
	})
	assert.Nil(t, err)
	task, _ = repo.Get(ctx, "testID")
	assert.Equal(t, "Test Task updated", task.Name)


	// getAll
	tasks, err := repo.GetAll(ctx)
	assert.Nil(t, err)
	assert.Equal(t, count1, len(tasks))

	// delete
	err = repo.Delete(ctx, "testID")
	assert.Nil(t, err)
	_, err = repo.Get(ctx, "testID")
	assert.Equal(t, sql.ErrNoRows, err)
	err = repo.Delete(ctx, "testID")
	assert.Equal(t, sql.ErrNoRows, err)
}
