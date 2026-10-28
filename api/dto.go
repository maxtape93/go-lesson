package api

type Booking struct {
	ID      int64   `json:"id"`
	HotelID int64   `json:"hotel_id"`
	UserID  int64   `json:"user_id"`
	Nights  int     `json:"nights"`
	Price   float64 `json:"price"`
	Status  string  `json:"status"`
	Email   string  `json:"email"`
}

type BookingResponse struct {
	ID      int64   `json:"id"`
	HotelID int64   `json:"hotel_id"`
	UserID  int64   `json:"user_id"`
	Nights  int     `json:"nights"`
	Price   float64 `json:"price"`
	Status  string  `json:"status"`
	Email   string  `json:"email"`
}
