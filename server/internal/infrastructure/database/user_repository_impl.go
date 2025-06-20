package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"backend/internal/domain/model"
	"backend/internal/domain/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserDto struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (dto *UserDto) ToModel() (*model.User, error) {
	id, err := uuid.Parse(dto.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse UUID: %w", err)
	}

	return &model.User{
		Id:        id,
		Name:      dto.Name,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}, nil
}

func (dto *UserDto) FromModel(user *model.User) {
	dto.Id = user.Id.String()
	dto.Name = user.Name
	dto.Email = user.Email
	dto.CreatedAt = user.CreatedAt
	dto.UpdatedAt = user.UpdatedAt
}

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	dto := &UserDto{}
	dto.FromModel(user)

	query := `
		INSERT INTO users (id, name, email, created_at, updated_at)
		VALUES (:id, :name, :email, :created_at, :updated_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *UserRepositoryImpl) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE id = ?
	`

	var dto UserDto
	err := r.db.GetContext(ctx, &dto, query, id.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user, err := dto.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
