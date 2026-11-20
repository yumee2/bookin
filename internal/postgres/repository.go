package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// maybe change to private later
type Repository struct {
	db *gorm.DB
}

// Init a repository and do the basic migrations
func NewRepository(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&House{})

	return &Repository{db: db}, nil
}
