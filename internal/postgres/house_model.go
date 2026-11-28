package postgres

import "yumee/booking/internal/booking"

type HouseModel struct {
	ID       int               `gorm:"primaryKey"`
	Type     booking.HouseType `gorm:"type:text;not null"`
	Capacity int               `gorm:"not null"`
}

// To implement Tabler interaface and rename the table in postgres
func (HouseModel) TableName() string {
	return "houses"
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
