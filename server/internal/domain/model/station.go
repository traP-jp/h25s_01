package model

import (
	"github.com/google/uuid"
	"time"
)

type Station struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStation(name string) (*Station, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, ErrInvalidStationName
	}

	return &Station{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
