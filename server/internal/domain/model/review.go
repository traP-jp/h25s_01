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

func NewReview(author UserID, shop uuid.UUID, rating Rating, content string, images []ImageFile) *Review {
	return &Review{
		ID:        uuid.New(),
		Author:    author,
		Shop:      shop,
		Rating:    rating,
		Content:   content,
		Images:    images,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Rating int

func NewRating(value int) (Rating, error) {
	if value < 0 || value > 3 {
		return 0, ErrInvalidRating
	}

	return Rating(value), nil
}
