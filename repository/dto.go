package repository

type Booking struct {
	ID      int64   `db:"id"`
	HotelID int64   `db:"hotel_id"`
	UserID  int64   `db:"user_id"`
	Nights  int     `db:"nights"`
	Price   float64 `db:"price"`
	Status  string  `db:"status"`
	Email   string  `db:"email"`
}
