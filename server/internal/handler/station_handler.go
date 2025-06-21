package handler

import (
	"backend/internal/domain/repository"
	"backend/internal/domain/model"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StationHandler struct {
	stationRepo repository.StationRepository
}

type StationDto struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewStationHandler(stationRepo repository.StationRepository) *StationHandler {
	return &StationHandler{
		stationRepo: stationRepo,
	}
}

func FromModelStation(s *model.Station) *StationDto {
	return &StationDto{
		ID: s.ID.String(),
		Name: s.Name,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
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

func (h *StationHandler) GetStations(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid Station ID")
	}

	Station, err := h.stationRepo.FindByID(c.Request().Context(), id)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to get user: %v", err))
	}

	return c.JSON(http.StatusOK, FromModelStation(Station))
}
