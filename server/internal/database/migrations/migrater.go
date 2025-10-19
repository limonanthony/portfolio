package migrations

import (
	"github.com/limonanthony/portfolio/internal/database"
	"github.com/limonanthony/portfolio/internal/reviews"
)

func RunMigrations(db *database.Database) error {
	return db.Db().AutoMigrate(&reviews.Review{})
}
