package model

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID        uuid.UUID
	Author    UserID
	Shop      uuid.UUID
	Rating    Rating
	Content   string
	Images    []ImageFile
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewReview(author UserID, shop uuid.UUID, rating Rating, content string, images []ImageFile) (*Review, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Review{
		ID:        id,
		Author:    author,
		Shop:      shop,
		Rating:    rating,
		Content:   content,
		Images:    images,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

type Rating int

func NewRating(value int) (Rating, error) {
	if value < 0 || value > 3 {
		return 0, ErrInvalidRating
	}

	return Rating(value), nil
}
