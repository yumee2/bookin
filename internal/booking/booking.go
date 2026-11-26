package booking

import (
	"context"
	"fmt"
)

type HouseRepository interface {
	GetHouses(ctx context.Context) ([]House, error)
	AddHouse(ctx context.Context, houseType HouseType, capacity int) (*House, error)
	RemoveHouse(ctx context.Context, houseID int) error
}

type BookingService struct {
	houseRepo HouseRepository
}

func NewBookingService(repo HouseRepository) *BookingService {
	return &BookingService{houseRepo: repo}
}

func (b *BookingService) GetHouses(ctx context.Context) ([]House, error) {
	houses, err := b.houseRepo.GetHouses(ctx)
	if err != nil {
		return nil, err
	}

	return houses, nil
}

func (b *BookingService) AddHouse(ctx context.Context, houseTypeRaw string, capacity int) (*House, error) {
	houseType, err := ParseHouseType(houseTypeRaw)
	if err != nil {
		return nil, fmt.Errorf("Invalid House Type")
	}

	house, err := b.houseRepo.AddHouse(ctx, houseType, capacity)
	if err != nil {
		return nil, fmt.Errorf("Error trying to add a house: ", err)
	}

	return house, nil
}

func (b *BookingService) RemoveHouse(ctx context.Context, houseID int) error {
	if houseID < 0 {
		return fmt.Errorf("House ID is invalid")
	}

	return b.houseRepo.RemoveHouse(ctx, houseID)
}
