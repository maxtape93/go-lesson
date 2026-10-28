package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Booking struct {
	ID      int64   `json:"id"`
	HotelID int64   `json:"hotel_id"`
	UserID  int64   `json:"user_id"`
	Nights  int     `json:"nights"`
	Price   float64 `json:"price"`
	Status  string  `json:"status"`
	Email   string  `json:"email"`
}

type App struct {
	DB *sql.DB
}

func (a *App) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var b Booking

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	if b.HotelID == 0 || b.UserID == 0 || b.Nights <= 0 {
		http.Error(w, "invalid booking", 400)
		return
	}

	if b.Nights > 30 {
		http.Error(w, "too many nights", 400)
		return
	}

	// Проверяем доступность отеля
	url := fmt.Sprintf(
		"http://hotel-provider/check?hotel_id=%d",
		b.HotelID,
	)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "provider error", 500)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		http.Error(w, "hotel unavailable", 400)
		return
	}

	// Рассчитываем стоимость
	b.Price = float64(b.Nights) * 5000

	if b.Nights >= 7 {
		b.Price = b.Price * 0.9
	}

	b.Status = "new"

	// Сохраняем бронирование
	err = a.DB.QueryRowContext(
		r.Context(),
		`INSERT INTO "bookings"
         ("hotel_id", "user_id", "nights", "price", "status", "email")
         VALUES ($1, $2, $3, $4, $5, $6)
         RETURNING id`,
		b.HotelID,
		b.UserID,
		b.Nights,
		b.Price,
		b.Status,
		b.Email,
	).Scan(&b.ID)

	if err != nil {
		http.Error(w, "db error", 500)
		return
	}

	// Отправляем email
	body, _ := json.Marshal(map[string]any{
		"email":      b.Email,
		"booking_id": b.ID,
	})

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://notification-service/send",
		bytes.NewReader(body),
	)
	if err != nil {
		http.Error(w, "notification error", 500)
		return
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "notification error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func main() {
	// Инициализация DB опущена
	var db *sql.DB

	app := &App{DB: db}

	http.HandleFunc("/bookings", app.CreateBooking)
	http.ListenAndServe(":8080", nil)
}
