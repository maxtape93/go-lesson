package hotel_provider

import (
	"encoding/json"
	"fmt"
	"go-lesson/domain"
	"net/http"
)

type Client struct {
	http *http.Client
}

func (c Client) GetHotel(hotelID int64) (domain.Hotel, error) {
	url := fmt.Sprintf(
		"http://hotel-provider/check?hotel_id=%d",
		hotelID,
	)

	resp, err := http.Get(url)
	if err != nil {
		return domain.Hotel{}, fmt.Errorf("provider request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Hotel{}, fmt.Errorf("provider status: %d", resp.StatusCode)
	}

	var dto hotelResponse
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return domain.Hotel{}, fmt.Errorf("decode provider response: %w", err)
	}

	return domain.Hotel{
		ID:     dto.ID,
		Name:   dto.Name,
		Active: dto.Active,
	}, nil
}
