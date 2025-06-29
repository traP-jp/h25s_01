package repository

import (
	"backend/internal/domain/model"
	"context"
	"github.com/google/uuid"
	"time"
)

type ReviewRepository interface {
	Save(ctx context.Context, user *model.Review) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Review, error)
	FindRecentReviews(
		ctx context.Context,
		after time.Time,
		before time.Time,
		limit int,
		offset int,
		shopID uuid.UUID,
		authorID model.UserID,
	) ([]*model.Review, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
