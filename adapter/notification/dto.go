package notification

type emailMessage struct {
	Email     string `json:"email"`
	BookingID int64  `json:"booking_id"`
}
