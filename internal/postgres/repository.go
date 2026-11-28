// This package provides functions to work with postgres repository
package postgres

import (
	"context"
	"fmt"
	"yumee/booking/internal/booking"

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

	err = db.AutoMigrate(&HouseModel{}, &ReservationModel{})
	if err != nil {
		return nil, fmt.Errorf("migrate database: %v", err)
	}

	return &Repository{db: db}, nil
}

// Return all houses with ID and its type
func (r *Repository) GetHouses(ctx context.Context) ([]booking.House, error) {
	var houseModels []HouseModel

	if err := r.db.WithContext(ctx).Find(&houseModels).Error; err != nil {
		return nil, fmt.Errorf("get houses: %w", err)
	}

	houses := make([]booking.House, len(houseModels))

	for i, houseModel := range houseModels {
		houses[i] = toHouse(houseModel)
	}

	return houses, nil
}

// optional
func (r *Repository) AddHouse(ctx context.Context, houseType booking.HouseType, capacity int) (*booking.House, error) {
	houseModel := HouseModel{
		Type:     houseType,
		Capacity: capacity,
	}

	if err := r.db.WithContext(ctx).Create(&houseModel).Error; err != nil {
		return nil, fmt.Errorf("create houses: %w", err)
	}

	house := toHouse(houseModel)

	return &house, nil
}

// optional
func (r *Repository) RemoveHouse(ctx context.Context, houseID int) error {
	result := r.db.WithContext(ctx).Where("id = ?", houseID).Delete(&HouseModel{})
	if result.Error != nil {
		return fmt.Errorf("delete houses: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return booking.HouseDoesNotExist
	}

	return nil
}
