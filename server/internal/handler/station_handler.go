package handler

import "time"

type Station struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type APIV1StationsPostRequest struct {
	Name string `json:"name"`
}
