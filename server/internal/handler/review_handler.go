package handler

import "time"

type Review struct {
	ID string `json:"id"`

	// レビュー投稿者のユーザーID
	Author string `json:"author"`

	// レビュー対象の店舗ID
	Shop string `json:"shop"`

	// 評価（0から3まで）
	Rating string `json:"rating,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Content string `json:"content"`

	Images []string `json:"images,omitempty"`
}

type APIV1ReviewsPostRequest struct {
	Author string `json:"author"`

	Shop string `json:"shop"`

	Rating string `json:"rating"`

	Content string `json:"content,omitempty"`

	Images []string `json:"images,omitempty"`
}
