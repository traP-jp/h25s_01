package handler

import (
	"backend/internal/domain/repository"
	"time"
)

type StationHandler struct {
	stationRepo repository.StationRepository
}

func NewStationHandler(stationRepo repository.StationRepository) *StationHandler {
	return &StationHandler{
		stationRepo: stationRepo,
	}
}

type Station struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type APIV1StationsPostRequest struct {
	Name string `json:"name"`
}
