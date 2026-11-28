package postgres

import "time"

type ReservationModel struct {
	ID               int        `gorm:"primaryKey;autoIncrement"`
	HouseID          int        `gorm:"not null;index"`
	House            HouseModel `gorm:"foreignKey:HouseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	GuestPhoneNumber string     `gorm:"type:varchar(32);not null"`
	CheckIn          time.Time  `gorm:"type:date;not null"`
	CheckOut         time.Time  `gorm:"type:date;not null;check:reservation_dates,check_out > check_in"`
	GuestCount       int        `gorm:"not null;check:reservation_guest_count,guest_count > 0"`
	Status           string     `gorm:"type:varchar(20);not null;check:reservation_status,status IN ('confirmed','cancelled', 'pending')"`
	CreatedAt        time.Time  `gorm:"not null"`
	UpdatedAt        time.Time  `gorm:"not null"`
}

// To implement Tabler interaface and rename the table in postgres
func (ReservationModel) TableName() string {
	return "reservations"
}
