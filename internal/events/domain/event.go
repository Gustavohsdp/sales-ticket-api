package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired          = errors.New("Event name is required")
	ErrEventDateMustBeInTheFuture = errors.New("Event date must be in the future")
	ErrEventCapactyZerp           = errors.New("Capacity must be greater than 0")
	ErrEventPriceZero             = errors.New("Price must be greater than 0")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

func (e Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrEventDateMustBeInTheFuture
	}

	if e.Capacity <= 0 {
		return ErrEventCapactyZerp
	}

	if e.Price <= 0 {
		return ErrEventPriceZero
	}

	return nil
}

func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)

	if err != nil {
		return nil, err
	}

	e.Spots = append(e.Spots, *spot)

	return spot, nil
}
