package database

import (
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

type Dialect = gorm.Dialector

func NewDatabase(dialect Dialect) (*Database, error) {
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Db() *gorm.DB {
	return d.db
}
