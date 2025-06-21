package handler

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ShopHandler struct {
	shopRepo repository.ShopRepository
}

type ShopDto struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Stations       []string  `json:"stations,omitempty"`
	PostCode       string    `json:"post_code,omitempty"`
	Address        string    `json:"address,omitempty"`
	Latitude       float64   `json:"latitude,omitempty"`
	Longitude      float64   `json:"longitude,omitempty"`
	Images         []string  `json:"images,omitempty"`
	PaymentMethods []string  `json:"payment_methods,omitempty"`
	Registerer     string    `json:"registerer,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

func NewShopHandler(shopRepo repository.ShopRepository) *ShopHandler {
	return &ShopHandler{
		shopRepo: shopRepo,
	}
}

func FromModelShop(s *model.Shop) *ShopDto {
	stations := make([]string, len(s.Stations))
	for i, station := range s.Stations {
		stations[i] = station.String()
	}

	images := make([]string, len(s.Images))
	for i, img := range s.Images {
		images[i] = img.Path
	}

	return &ShopDto{
		ID:             s.ID.String(),
		Name:           string(s.Name),
		Stations:       stations,
		Address:        s.Address,
		PostCode:       string(s.PostCode),
		Latitude:       s.Latitude,
		Longitude:      s.Longitude,
		Images:         images,
		PaymentMethods: s.PaymentMethods,
		Registerer:     string(s.Registerer),
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}

type Shop struct {
	ID string `json:"id"`

	Name string `json:"name"`

	PostCode string `json:"post_code,omitempty"`

	Address string `json:"address,omitempty"`

	Latitude float64 `json:"latitude,omitempty"`

	Longitude float64 `json:"longitude,omitempty"`

	Images []string `json:"images,omitempty"`

	PaymentMethods []string `json:"payment_methods,omitempty"`

	// 関連する駅のID配列
	Stations []string `json:"stations,omitempty"`

	Registerer string `json:"registerer,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func FromModelToShop(m *model.Shop) *Shop {
	stations := make([]string, len(m.Stations))
	for i, station := range m.Stations {
		stations[i] = station.String()
	}

	images := make([]string, len(m.Images))
	for i, img := range m.Images {
		images[i] = img.Path
	}

	return &Shop{
		ID:             m.ID.String(),
		Name:           string(m.Name),
		Address:        m.Address,
		PostCode:       string(m.PostCode),
		Latitude:       m.Latitude,
		Longitude:      m.Longitude,
		Images:         images,
		PaymentMethods: m.PaymentMethods,
		Stations:       stations,
		Registerer:     string(m.Registerer),
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

type APIV1ShopsPostRequest struct {
	Name string `json:"name"`

	PostCode string `json:"post_code,omitempty"`

	Address string `json:"address"`

	Latitude float64 `json:"latitude,omitempty"`

	Longitude float64 `json:"longitude,omitempty"`

	Images []string `json:"images,omitempty"`

	PaymentMethods []string `json:"payment_methods,omitempty"`

	// 関連する駅のID配列
	Stations []string `json:"stations,omitempty"`

	Registerer string `json:"registerer,omitempty"`
}

type APIV1ShopsIDImagesPost200Response struct {
	ImageURL string `json:"image_url,omitempty"`
}

type APIV1ShopsIDImagesDeleteRequest struct {
	ImageURL string `json:"image_url"`
}

func (h *ShopHandler) GetShopDetail(c echo.Context) error {
	shopID := c.Param("id")
	uuidShopID, err := uuid.Parse(shopID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid shop ID format",
		})
	}
	shop, err := h.shopRepo.FindByID(c.Request().Context(), uuidShopID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	if shop == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Shop not found",
		})
	}

	return c.JSON(http.StatusOK, FromModelToShop(shop))
}

func (h *ShopHandler) UpdateShop(c echo.Context) error {
	shopID := c.Param("id")
	uuidShopID, err := uuid.Parse(shopID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid shop ID format",
		})
	}

	var req APIV1ShopsPostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	shop, err := h.shopRepo.FindByID(c.Request().Context(), uuidShopID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	if shop == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Shop not found",
		})
	}

	if req.Name != "" {
		name, err := model.NewShopName(req.Name)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid shop name",
			})
		}
		shop.Name = name
	}

	if req.PostCode != "" {
		postCode, err := model.NewPostCode(req.PostCode)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid post code",
			})
		}
		shop.PostCode = postCode
	}

	if req.Latitude != 0 || req.Longitude != 0 {
		if req.Latitude < -90 || req.Latitude > 90 || req.Longitude < -180 || req.Longitude > 180 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid latitude or longitude",
			})
		}

		if req.Latitude == 0 && req.Longitude == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Latitude and longitude cannot both be zero",
			})
		}

		if req.Latitude == 0 || req.Longitude == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Latitude and longitude must be provided together",
			})
		}

		shop.Latitude = req.Latitude
		shop.Longitude = req.Longitude
	}
	if req.Images != nil {
		images := make([]model.ImageFile, len(req.Images))
		for i, imgPath := range req.Images {
			images[i] = model.ImageFile{Path: imgPath}
		}
		shop.Images = images
	}
	if req.PaymentMethods != nil {
		shop.PaymentMethods = req.PaymentMethods
	}
	if req.Stations != nil {
		stationUUIDs := make([]uuid.UUID, len(req.Stations))
		for i, s := range req.Stations {
			u, err := uuid.Parse(s)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "Invalid station UUID: " + s,
				})
			}
			stationUUIDs[i] = u
		}
		shop.Stations = stationUUIDs
	}
	if req.Registerer != "" {
		userID, err := model.NewUserID(req.Registerer)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid registerer user ID",
			})
		}
		shop.Registerer = userID
	}

	shop.UpdatedAt = time.Now()

	if err := h.shopRepo.Save(c.Request().Context(), shop); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, FromModelToShop(shop))
}

func (h *ShopHandler) Delete(c echo.Context) error {
	shopID := c.Param("id")
	uuidShopID, err := uuid.Parse(shopID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid shop ID format",
		})
	}

	if err := h.shopRepo.Delete(c.Request().Context(), uuidShopID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Shop deleted successfully",
	})
}
func (h *ShopHandler) GetShops(c echo.Context) error {
	shops, err := h.shopRepo.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	responses := make([]*Shop, len(shops))
	for i, v := range shops {
		responses[i] = FromModelToShop(v)
	}

	return c.JSON(http.StatusCreated, responses)
}

func (h *ShopHandler) CreateShop(c echo.Context) error {
	var req APIV1ShopsPostRequest
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

	shopName, _ := model.NewShopName(req.Name)
	postCode, _ := model.NewPostCode(req.PostCode)
	shop, err := model.NewShop(shopName, req.Address, postCode, req.Latitude, req.Longitude, req.images, req.PaymentMethods, string(req.Registerer), stations)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	if err := h.shopRepo.Save(c.Request().Context(), shop); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, FromModelToShop(shop))

}
