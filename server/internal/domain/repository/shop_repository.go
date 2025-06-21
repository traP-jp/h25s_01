package repository

import (
	"backend/internal/domain/model"
	"context"
	"io"
	"github.com/google/uuid"
)

type ShopRepository interface {
	Save(ctx context.Context, user *model.Shop) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Shop, error)
	FindAll(ctx context.Context) ([]*model.Shop, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByStation(ctx context.Context, id uuid.UUID) ([]*model.Shop, error)
    UploadImage(ctx context.Context, contentType string, src io.Reader) (uuid.UUID, error)
}
