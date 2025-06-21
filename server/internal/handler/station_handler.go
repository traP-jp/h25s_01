package handler

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"fmt"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StationHandler struct {
	stationRepo repository.StationRepository
	shopRepo repository.ShopRepository
}

type StationDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewStationHandler(stationRepo repository.StationRepository, shopRepo repository.ShopRepository) *StationHandler {
	return &StationHandler{
		stationRepo: stationRepo,
		shopRepo: shopRepo,
	}
}

func FromModelStation(s *model.Station) *StationDto {
	return &StationDto{
		ID:        s.ID.String(),
		Name:      s.Name,
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

func FromModel(m *model.Station) *Station {
	return &Station{
		ID:        m.ID.String(),
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

type APIV1StationsPostRequest struct {
	Name string `json:"name"`
}

func errorResponse(c echo.Context, status int, msg string) error {
	return c.JSON(status, map[string]string{"error": msg})
}

func (h *StationHandler) CreateStation(c echo.Context) error {
	var req APIV1StationsPostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if err := validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required),
	); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	station, err := model.NewStation(req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	if err := h.stationRepo.Save(c.Request().Context(), station); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, FromModel(station))

}

func (h *StationHandler) GetStations(c echo.Context) error {
	stations, err := h.stationRepo.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	responses := make([]*Station, len(stations))
	for i, v := range stations {
		responses[i] = FromModel(v)
	}

	return c.JSON(http.StatusCreated, responses)
}

func (h *StationHandler) UpdateStation(c echo.Context) error {
	stationid := c.Param("id")
	id, err := uuid.Parse(stationid)
	var req APIV1StationsPostRequest
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid station ID")
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if err := validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required),
	); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	station, err := h.stationRepo.FindByID(c.Request().Context(), id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Invalid name",
		})
	}
	station.Name = req.Name
	station.UpdatedAt = time.Now()
	if err := h.stationRepo.Save(c.Request().Context(), station); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, FromModel(station))
}

func (h *StationHandler) DeleteStation(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid station ID")
	}

	if err := h.stationRepo.Delete(c.Request().Context(), id); err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to delete station: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "station deleted successfully",
	})
}

func (h *StationHandler) GetStationDetail(c echo.Context) error {
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
func (h *StationHandler) GetShopAroundStation(c echo.Context) error {
	stationid := c.Param("id")
	id, err := uuid.Parse(stationid)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid station ID")
	}
	shops, err := h.shopRepo.FindByStation(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	responses := shops

	return c.JSON(http.StatusCreated, responses)
}
