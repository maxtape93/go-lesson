package api

import (
	"fmt"
)

const minNights = 0
const maxNights = 30

func validateBooking(b Booking) error {
	if b.HotelID == 0 || b.UserID == 0 || b.Nights <= minNights {
		return fmt.Errorf("invalid booking")
	}

	if b.Nights > maxNights {
		return fmt.Errorf("too many night")
	}

	return nil
}
