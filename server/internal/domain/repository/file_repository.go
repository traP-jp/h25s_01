package repository

import (
	"github.com/google/uuid"
	"io"
)

type FileRepository interface {
	UploadImage(reviewID uuid.UUID, content string, reader io.Reader) (uuid.UUID, error)
	DeleteImage(fileID uuid.UUID) error
}
