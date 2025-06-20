package repository

import (
	"context"

	"backend/internal/domain/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
