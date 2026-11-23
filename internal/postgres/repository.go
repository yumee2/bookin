// This package provides functions to work with postgres repository
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
func (r *Repository) addHouse(ctx context.Context, houseType HouseType, capacity int) (*House, error) {
	house := House{
		Type:     houseType,
		Capacity: capacity,
	}

	if err := r.db.WithContext(ctx).Create(&house).Error; err != nil {
		return nil, fmt.Errorf("create houses: %w", err)
	}

	return &house, nil
}

// optional
func (r *Repository) removeHouse(ctx context.Context, houseID int) error {
	result := r.db.WithContext(ctx).Where("id = ?", houseID).Delete(&House{})
	if result.Error != nil {
		return fmt.Errorf("delete houses: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("house with this ID was not found")
	}

	return nil
}
