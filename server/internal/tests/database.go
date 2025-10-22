package tests

import (
	"fmt"
	"testing"

	"github.com/limonanthony/portfolio/internal/database"
	"github.com/limonanthony/portfolio/internal/database/migrations"
	"gorm.io/driver/sqlite"
)

// NewDatabase creates a new test database with migrations applied
func NewDatabase(t *testing.T) (*database.Database, error) {
	// Use unique database name for each test to avoid shared state
	dbName := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
	db, err := database.NewDatabase(sqlite.Open(dbName))
	if err != nil {
		return nil, err
	}

	// Run migrations using the migrations package
	if err := migrations.RunMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}
