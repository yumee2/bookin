package postgres

import (
	"context"
	"fmt"

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

// Return all houses with ID and its type
func (r *Repository) getHouses(ctx context.Context) ([]House, error) {
	var houses []House

	if err := r.db.WithContext(ctx).Find(&houses).Error; err != nil {
		return nil, fmt.Errorf("get houses: %w", err)
	}

	return houses, nil
}

// optional
// func (r *Repository) addHouseType() House {
//
// }
