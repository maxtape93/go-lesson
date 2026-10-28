package api

import (
	"go-lesson/domain"
	"go-lesson/usecase"
)

func toCreateBookingRequest(b Booking) (request *usecase.CreateBookingRequest, error) {
	err := validateBooking(b)
	if err != nil {
		return nil, err
	}

	return &usecase.CreateBookingRequest{
		ID:      b.ID,
		HotelID: b.HotelID,
		UserID:  b.UserID,
		Nights:  b.Nights,
		Price:   b.Price,
		Status:  b.Status,
		Email:   b.Email,
	}, nil
}

func toBookingResponse(b *domain.Booking) BookingResponse {
	return BookingResponse{
		ID:      *b.ID,
		HotelID: b.HotelID,
		UserID:  b.UserID,
		Nights:  b.Nights,
		Price:   b.Price,
		Status:  b.Status,
		Email:   b.Email,
	}
}
