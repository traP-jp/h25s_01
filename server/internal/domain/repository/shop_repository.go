package repository

import (
	"backend/internal/domain/model"
	"context"
	"github.com/google/uuid"
)

type ShopRepository interface {
	Save(ctx context.Context, user *model.Shop) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Shop, error)
	FindAllWithLimit(ctx context.Context, limit, offset int) ([]*model.Shop, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByStation(ctx context.Context, id uuid.UUID) ([]*model.Shop, error)
}
