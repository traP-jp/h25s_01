package handler

import (
	"backend/internal/domain/repository"
	"time"
)

type ReviewHandler struct {
	reviewRepo repository.ReviewRepository
}

func NewReviewHandler(reviewRepo repository.ReviewRepository) *ReviewHandler {
	return &ReviewHandler{
		reviewRepo: reviewRepo,
	}
}

type Review struct {
	ID string `json:"id"`

	// レビュー投稿者のユーザーID
	Author string `json:"author"`

	// レビュー対象の店舗ID
	Shop string `json:"shop"`

	// 評価（0から3まで）
	Rating int32 `json:"rating,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Content string `json:"content,omitempty"`

	Images []string `json:"images,omitempty"`
}

type APIV1ReviewsPostRequest struct {
	Author string `json:"author"`

	Shop string `json:"shop"`

	Rating int32 `json:"rating"`

	Content string `json:"content,omitempty"`

	Images []string `json:"images,omitempty"`
}
