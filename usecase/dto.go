package usecase

type CreateBookingRequest struct {
	ID      int64
	HotelID int64
	UserID  int64
	Nights  int
	Price   float64
	Status  string
	Email   string
}
