package model

import "github.com/google/uuid"

type ImageFile struct {
	ID   uuid.UUID
	Path string
}

func NewImageFile(id uuid.UUID) *ImageFile {
	return &ImageFile{
		ID:   id,
		Path: "/images/" + id.String(),
	}
}

type UserID string

func NewUserID(id string) (UserID, error) {
	if id == "" {
		return "", ErrInvalidUserID
	}

	return UserID(id), nil
}
