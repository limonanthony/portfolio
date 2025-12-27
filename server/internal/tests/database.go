package tests

import (
	"fmt"
	"testing"

	"github.com/limonanthony/portfolio/internal/database"
	"github.com/limonanthony/portfolio/internal/database/migrations"
	"gorm.io/driver/sqlite"
)

func NewDatabase(t *testing.T) (*database.Database, error) {
	dbName := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
	db, err := database.NewDatabase(sqlite.Open(dbName))
	if err != nil {
		return nil, err
	}

	if err := migrations.RunMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}
