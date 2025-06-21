package model

import (
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Shop struct {
	ID             uuid.UUID
	Name           ShopName
	Stations       []uuid.UUID
	Address        string
	PostCode       PostCode
	Latitude       float64
	Longitude      float64
	Images         []ImageFile
	PaymentMethods []string
	Registerer     UserID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewShop(
	name ShopName,
	address string,
	postCode PostCode,
	latitude, longitude float64,
	images []ImageFile,
	paymentMethods []string,
	registerer UserID,
	stations []uuid.UUID,
) (*Shop, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Shop{
		ID:             id,
		Name:           name,
		Address:        address,
		PostCode:       postCode,
		Latitude:       latitude,
		Longitude:      longitude,
		Images:         images,
		PaymentMethods: paymentMethods,
		Registerer:     registerer,
		Stations:       stations,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

type ShopName string

func NewShopName(name string) (ShopName, error) {
	if name == "" {
		return "", ErrInvalidShopName
	}

	return ShopName(name), nil
}

type PostCode string

func NewPostCode(code string) (PostCode, error) {
	matched, err := regexp.MatchString("^\\d{3}-\\d{4}$", code)
	if err != nil || !matched {
		return "", ErrInvalidPostCode
	}

	return PostCode(code), nil
}
