package notification

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	http *http.Client
}

func (c *Client) SendEmail(bookingID int64, email string) error {
	msg := emailMessage{
		Email:     email,
		BookingID: bookingID,
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal email: %w", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://notification-service/send",
		bytes.NewReader(body),
	)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("send email: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("notification status: %d", resp.StatusCode)
	}
	return nil
}
