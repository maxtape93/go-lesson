package usecase

import (
	"context"
	"fmt"
	"go-lesson/domain"
)

type CreateBookingUseCase struct {
	hotel  HotelProvider
	repo   BookingRepository
	sender Sender
}

func (uc CreateBookingUseCase) CreateBooking(ctx context.Context, in *CreateBookingRequest) (*domain.Booking, error) {
	hotel, err := uc.hotel.GetHotel(in.HotelID)

	if err != nil {
		return nil, fmt.Errorf("get hotel error: %w", err)
	}

	if !hotel.Active {
		return nil, domain.ErrHotelUnavailable
	}

	domainBooking := toDomainBooking(in)

	domainBooking.Calculate()

	bookingID, err := uc.repo.CreateBooking(ctx, domainBooking)

	if err != nil {
		return nil, fmt.Errorf("create booking: %w", err)
	}

	domainBooking.ID = &bookingID

	err = uc.sender.SendEmail(bookingID, domainBooking.Email)

	if err != nil {
		return nil, fmt.Errorf("send email: %w", err)
	}

	return domainBooking, nil
}
