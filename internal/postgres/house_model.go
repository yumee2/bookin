package postgres

import "yumee/booking/internal/booking"

type HouseType string

type HouseModel struct {
	ID       uint              `gorm:"primaryKey"`
	Type     booking.HouseType `gorm:"type:text;not null"`
	Capacity int               `gorm:"not null"`
}

func toHouseModel(h booking.House) HouseModel {
	return HouseModel{
		ID:       h.ID,
		Type:     h.Type,
		Capacity: h.Capacity,
	}
}

func toHouse(m HouseModel) booking.House {
	return booking.House{
		ID:       m.ID,
		Type:     booking.HouseType(m.Type),
		Capacity: m.Capacity,
	}
}
