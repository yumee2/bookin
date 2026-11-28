package booking

import "time"

type ReservationStatus string

const (
	Confirmed ReservationStatus = "confirmed"
	Cancelled ReservationStatus = "cancelled"
	Pending   ReservationStatus = "pending"
)

type Reservation struct {
	ID               int               `json:"id"`
	HouseID          int               `json:"house_id"`
	GuestPhoneNumber string            `json:"guest_phone_number"`
	CheckIn          time.Time         `json:"check_in"`
	CheckOut         time.Time         `json:"check_out"`
	GuestCount       int               `json:"guest_count"`
	Status           ReservationStatus `json:"status"`
	CreatedAt        time.Time         `json:"created_at"`
	Updatedt         time.Time         `json:"updated_at"`
}
