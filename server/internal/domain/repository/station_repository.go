package repository

import (
	"backend/internal/domain/model"
	"context"
	"github.com/google/uuid"
)

type StationRepository interface {
	Save(ctx context.Context, user *model.Station) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Station, error)
	FindAll(ctx context.Context) ([]*model.Station, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
