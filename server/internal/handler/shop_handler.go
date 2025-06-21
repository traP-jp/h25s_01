package handler

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"time"
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
