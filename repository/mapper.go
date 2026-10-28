package repository

import "go-lesson/domain"

func toDBBooking(booking *domain.Booking) Booking {
	return Booking{
		HotelID: booking.HotelID,
		Nights:  booking.Nights,
		UserID:  booking.UserID,
		Price:   booking.Price,
		Status:  booking.Status,
		Email:   booking.Email,
	}
}
