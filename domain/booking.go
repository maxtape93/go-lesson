package domain

const StatusNew = "New"

const (
	pricePerNight  = 5000
	discountNights = 7
	discountFactor = 0.9
)

func (b *Booking) Calculate() {
	b.Price = float64(b.Nights) * pricePerNight
	if b.Nights >= discountNights {
		b.Price *= discountFactor
	}
	b.Status = StatusNew
}
