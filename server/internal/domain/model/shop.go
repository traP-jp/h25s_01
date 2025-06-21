package model

import (
	"github.com/google/uuid"
	"regexp"
	"time"
)

type Shop struct {
	ID             uuid.UUID
	Name           ShopName
	postCode       PostCode
	latitude       float64
	longitude      float64
	images         []ImageFile
	paymentMethods []string
	registerer     UserID
	createdAt      time.Time
	updatedAt      time.Time
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
