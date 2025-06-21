package model

import (
	"github.com/google/uuid"
	"regexp"
	"time"
)

type Shop struct {
	ID             uuid.UUID
	Name           ShopName
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
	postCode PostCode,
	latitude, longitude float64,
	images []ImageFile,
	paymentMethods []string,
	registerer UserID,
) (*Shop, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Shop{
		ID:             id,
		Name:           name,
		PostCode:       postCode,
		Latitude:       latitude,
		Longitude:      longitude,
		Images:         images,
		PaymentMethods: paymentMethods,
		Registerer:     registerer,
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
