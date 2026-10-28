package repository

import (
	"context"
	"fmt"
	"go-lesson/domain"
)

func (r *Repository) CreateBooking(ctx context.Context, booking *domain.Booking) (int64, error) {
	bookingDB := toDBBooking(booking)

	var id int64

	err := r.DB.QueryRowContext(
		ctx,
		`INSERT INTO "bookings"
         ("hotel_id", "user_id", "nights", "price", "status", "email")
         VALUES ($1, $2, $3, $4, $5, $6)
         RETURNING id`,
		bookingDB.HotelID,
		bookingDB.UserID,
		bookingDB.Nights,
		bookingDB.Price,
		bookingDB.Status,
		bookingDB.Email,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("insert booking: %w", err)
	}

	return id, nil
}
