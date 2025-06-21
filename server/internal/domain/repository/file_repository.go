package repository

import (
	"github.com/google/uuid"
	"io"
)

type FileRepository interface {
	UploadImage(reviewID uuid.UUID, contentType string, reader io.Reader) (uuid.UUID, error)
	DeleteImage(fileID uuid.UUID) error
}
