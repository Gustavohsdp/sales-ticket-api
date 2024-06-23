package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner1 struct {
	BaseURL string
}

type Partner1ReservationRequest struct {
	Spots      []string `json:"spots"`
	TicketType string   `json:"ticket_kind"`
	Email      string   `json:"email"`
}

type Partner1ReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketType string `json:"ticket_kind"`
	Status     string `json:"status"`
	EventID    string `json:"event_id"`
}

func (p *Partner1) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	partnerRequest := Partner1ReservationRequest{
		Spots:      req.Spots,
		TicketType: req.TicketType,
		Email:      req.Email,
	}

	body, err := json.Marshal(partnerRequest)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/events/%s/reserve", p.BaseURL, req.EventID)

	httpRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return nil, err
	}

	// Close the response body

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Unexpected status code: %d", httpResponse.StatusCode)
	}

	var partnerResponse []Partner1ReservationResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&partnerResponse); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResponse))

	for i, r := range partnerResponse {
		responses[i] = ReservationResponse{
			ID:     r.ID,
			Spot:   r.Spot,
			Status: r.Status,
		}
	}

	return responses, nil
}
