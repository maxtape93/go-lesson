package domain

import "errors"

var (
	ErrHotelUnavailable = errors.New("hotel unavailable")
)

type CreateBookingResponse struct {
}

type Hotel struct {
	ID     int64
	Name   string
	Active bool
}

type Booking struct {
	ID      *int64
	HotelID int64
	UserID  int64
	Nights  int
	Price   float64
	Status  string
	Email   string
}
