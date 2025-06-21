package handler

import "time"

type Shop struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name"`

	PostCode string `json:"post_code,omitempty"`

	Address string `json:"address"`

	Latitude string `json:"latitude,omitempty"`

	Longitude string `json:"longitude,omitempty"`

	Images []string `json:"images,omitempty"`

	PaymentMethods []string `json:"payment_methods,omitempty"`

	Registerer string `json:"registerer,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Review struct {
	ID string `json:"id,omitempty"`

	// レビュー投稿者のユーザーID
	Author string `json:"author"`

	// レビュー対象の店舗ID
	Shop string `json:"shop"`

	// 評価（0から3まで）
	Rating string `json:"rating"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Content string `json:"content,omitempty"`

	Images []string `json:"images,omitempty"`
}

type APIV1ShopsPostRequest struct {
	Name string `json:"name"`

	PostCode string `json:"post_code,omitempty"`

	Address string `json:"address"`

	Latitude string `json:"latitude,omitempty"`

	Longitude string `json:"longitude,omitempty"`

	Images []string `json:"images,omitempty"`

	PaymentMethods []string `json:"payment_methods,omitempty"`

	Registerer string `json:"registerer,omitempty"`
}

type APIV1ShopsIDImagesPost200Response struct {
	ImageURL string `json:"image_url,omitempty"`
}

type APIV1ShopsIDImagesDeleteRequest struct {
	ImageURL string `json:"image_url"`
}

type APIV1ReviewsPostRequest struct {
	Author string `json:"author"`

	Shop string `json:"shop"`

	Rating string `json:"rating"`

	Content string `json:"content,omitempty"`

	Images []string `json:"images,omitempty"`
}
