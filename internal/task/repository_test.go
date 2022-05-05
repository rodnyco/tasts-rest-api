package task

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rodnyco/tasks-rest-api/internal/config"
	"log"
	"testing"
)


func TestRepository_Get(t *testing.T) {
	cfg := config.DefaultConfig()

	driverName := "postgres"
	dataSource := cfg.DSN

	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatalln(err)
	}

	repository := NewRepository(db)
	res, err := repository.Get(context.Background(), "1")
	if err != nil {
		log.Fatalln(err)
	}

	if res.ID != "1" {
		t.Error("Expected 1 got ", res.ID)
	}
}
