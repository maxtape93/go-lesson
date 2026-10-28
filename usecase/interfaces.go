package usecase

import (
	"context"
	"go-lesson/domain"
)

type HotelProvider interface {
	GetHotel(hotelID int64) (domain.Hotel, error)
}

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *domain.Booking) (int64, error)
}

type Sender interface {
	SendEmail(bookingID int64, email string) error
}

type Service interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest) (*domain.Booking, error)
}
