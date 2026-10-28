package usecase

import "go-lesson/domain"

func toDomainBooking(in *CreateBookingRequest) *domain.Booking {
	return &domain.Booking{
		HotelID: in.HotelID,
		UserID:  in.UserID,
		Nights:  in.Nights,
		Price:   in.Price,
		Status:  in.Status,
		Email:   in.Email,
	}
}
