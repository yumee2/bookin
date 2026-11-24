package booking

import (
	"context"
)

type HouseRepository interface {
	GetHouses(ctx context.Context) ([]House, error)
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
