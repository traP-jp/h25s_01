package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ReviewDto struct {
	ID        string    `db:"id"`
	Author    string    `db:"author"`
	ShopID    string    `db:"shop_id"`
	Rating    int       `db:"rating"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ReviewImageDto struct {
	ReviewID string `db:"review_id"`
	ImageID  string `db:"image_id"`
}

func (dto *ReviewDto) ToModel(images []model.ImageFile) (*model.Review, error) {
	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse review UUID: %w", err)
	}

	shopID, err := uuid.Parse(dto.ShopID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse shop UUID: %w", err)
	}

	userID, err := model.NewUserID(dto.Author)
	if err != nil {
		return nil, fmt.Errorf("failed to create UserID: %w", err)
	}

	rating, err := model.NewRating(dto.Rating)
	if err != nil {
		return nil, fmt.Errorf("failed to create Rating: %w", err)
	}

	return &model.Review{
		ID:        id,
		Author:    userID,
		Shop:      shopID,
		Rating:    rating,
		Content:   dto.Content,
		Images:    images,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}, nil
}

func (dto *ReviewDto) FromModel(review *model.Review) {
	dto.ID = review.ID.String()
	dto.Author = string(review.Author)
	dto.ShopID = review.Shop.String()
	dto.Rating = int(review.Rating)
	dto.Content = review.Content
	dto.CreatedAt = review.CreatedAt
	dto.UpdatedAt = review.UpdatedAt
}

type ReviewRepositoryImpl struct {
	db *sqlx.DB
}

func NewReviewRepository(db *sqlx.DB) repository.ReviewRepository {
	return &ReviewRepositoryImpl{
		db: db,
	}
}

func (r *ReviewRepositoryImpl) Save(ctx context.Context, review *model.Review) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			fmt.Printf("failed to rollback transaction: %v\n", err)
		}
	}()

	// Save review
	dto := &ReviewDto{}
	dto.FromModel(review)

	query := `
		INSERT INTO reviews (id, author, shop_id, rating, content, created_at, updated_at)
		VALUES (:id, :author, :shop_id, :rating, :content, :created_at, :updated_at)
		ON DUPLICATE KEY UPDATE
			author = VALUES(author),
			shop_id = VALUES(shop_id),
			rating = VALUES(rating),
			content = VALUES(content),
			updated_at = VALUES(updated_at)
	`

	_, err = tx.NamedExecContext(ctx, query, dto)
	if err != nil {
		return fmt.Errorf("failed to save review: %w", err)
	}

	// Delete existing review images
	deleteImagesQuery := `DELETE FROM review_images WHERE review_id = ?`
	_, err = tx.ExecContext(ctx, deleteImagesQuery, review.ID.String())
	if err != nil {
		return fmt.Errorf("failed to delete existing review images: %w", err)
	}

	// Save review images
	if len(review.Images) > 0 {
		imageQuery := `INSERT INTO review_images (review_id, image_id) VALUES (?, ?)`
		for _, image := range review.Images {
			_, err = tx.ExecContext(ctx, imageQuery, review.ID.String(), image.ID.String())
			if err != nil {
				return fmt.Errorf("failed to save review image: %w", err)
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *ReviewRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*model.Review, error) {
	query := `
		SELECT *
		FROM reviews
		WHERE id = ?
	`

	var dto ReviewDto
	err := r.db.GetContext(ctx, &dto, query, id.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("review not found")
		}

		return nil, fmt.Errorf("failed to get review: %w", err)
	}

	// Get review images
	images, err := r.getReviewImages(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get review images: %w", err)
	}

	review, err := dto.ToModel(images)
	if err != nil {
		return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
	}

	return review, nil
}

func (r *ReviewRepositoryImpl) FindRecentReviews(
	ctx context.Context,
	after time.Time,
	before time.Time,
	limit int,
	offset int,
	shopID uuid.UUID,
	authorID model.UserID,
) ([]*model.Review, error) {
	conditions := []string{}
	args := []interface{}{}

	if !after.IsZero() {
		conditions = append(conditions, "created_at >= ?")
		args = append(args, after)
	}

	if !before.IsZero() {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, before)
	}

	if shopID != uuid.Nil {
		conditions = append(conditions, "shop_id = ?")
		args = append(args, shopID.String())
	}

	if authorID != "" {
		conditions = append(conditions, "author = ?")
		args = append(args, string(authorID))
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	query := fmt.Sprintf(`
		SELECT *
		FROM reviews
		%s
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, whereClause)

	args = append(args, limit, offset)

	var dtos []ReviewDto
	err := r.db.SelectContext(ctx, &dtos, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get reviews: %w", err)
	}

	reviews := make([]*model.Review, 0, len(dtos))
	for _, dto := range dtos {
		// Get review images for each review
		reviewID, err := uuid.Parse(dto.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse review ID: %w", err)
		}

		images, err := r.getReviewImages(ctx, reviewID)
		if err != nil {
			return nil, fmt.Errorf("failed to get review images: %w", err)
		}

		review, err := dto.ToModel(images)
		if err != nil {
			return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *ReviewRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			fmt.Printf("failed to rollback transaction: %v\n", err)
		}
	}()

	// Delete review images first (foreign key constraint)
	deleteImagesQuery := `DELETE FROM review_images WHERE review_id = ?`
	_, err = tx.ExecContext(ctx, deleteImagesQuery, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete review images: %w", err)
	}

	// Delete review
	deleteReviewQuery := `DELETE FROM reviews WHERE id = ?`
	result, err := tx.ExecContext(ctx, deleteReviewQuery, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete review: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("review not found")
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *ReviewRepositoryImpl) getReviewImages(ctx context.Context, reviewID uuid.UUID) ([]model.ImageFile, error) {
	query := `
		SELECT image_id
		FROM review_images
		WHERE review_id = ?
	`

	var imageIDs []string
	err := r.db.SelectContext(ctx, &imageIDs, query, reviewID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get review image IDs: %w", err)
	}

	images := make([]model.ImageFile, 0, len(imageIDs))
	for _, imageIDStr := range imageIDs {
		imageID, err := uuid.Parse(imageIDStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse image ID: %w", err)
		}
		images = append(images, *model.NewImageFile(imageID))
	}

	return images, nil
}
