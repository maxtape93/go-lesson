package api

import (
	"context"
	"encoding/json"
	"errors"
	"go-lesson/domain"
	"go-lesson/usecase"
	"net/http"
)

type ServerAPI struct {
	service usecase.Service
}

func (s *ServerAPI) CreateBooking(w http.ResponseWriter, r *http.Request) {

	var b Booking

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return

	}

	request, err := toCreateBookingRequest(b)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	booking, err := s.service.CreateBooking(context.Background(), request)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrHotelUnavailable):
			http.Error(w, "hotel unavailable", http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(toBookingResponse(booking))
}
