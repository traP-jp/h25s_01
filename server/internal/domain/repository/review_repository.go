package repository

import (
	"backend/internal/domain/model"
	"context"
	"github.com/google/uuid"
)

type ReviewRepository interface {
	Save(ctx context.Context, user *model.Review) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Review, error)
	FindRecentReviews(
		ctx context.Context,
		limit int,
		offset int,
		shopID uuid.UUID,
		authorID uuid.UUID,
	) ([]*model.Review, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
