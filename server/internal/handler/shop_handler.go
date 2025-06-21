package handler

import "time"

type Shop struct {
	ID string `json:"id"`

	Name string `json:"name"`

	PostCode string `json:"post_code,omitempty"`

	Address string `json:"address,omitempty"`

	Latitude string `json:"latitude,omitempty"`

	Longitude string `json:"longitude,omitempty"`

	Images []string `json:"images,omitempty"`

	PaymentMethods []string `json:"payment_methods,omitempty"`

	// 関連する駅のID配列
	Stations []string `json:"stations,omitempty"`

	Registerer string `json:"registerer,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type APIV1ShopsPostRequest struct {
	Name string `json:"name"`

	PostCode string `json:"post_code,omitempty"`

	Address string `json:"address"`

	Latitude string `json:"latitude,omitempty"`

	Longitude string `json:"longitude,omitempty"`

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
