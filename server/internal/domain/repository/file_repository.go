package repository

import (
	"context"
	"io"

	"github.com/google/uuid"
)

type FileRepository interface {
	UploadImage(ctx context.Context, contentType string, reader io.Reader) (uuid.UUID, error)
	DeleteImage(ctx context.Context, fileID uuid.UUID) error
	GetImage(ctx context.Context, fileID uuid.UUID) (io.ReadCloser, string, error)
}
