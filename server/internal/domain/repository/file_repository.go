package repository

import (
	"io"

	"github.com/google/uuid"
)

type FileRepository interface {
	UploadImage(contentType string, reader io.Reader) (uuid.UUID, error)
	DeleteImage(fileID uuid.UUID) error
}
