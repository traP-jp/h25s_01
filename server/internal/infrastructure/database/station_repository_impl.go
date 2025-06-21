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

type StationDto struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (dto *StationDto) ToModel() (*model.Station, error) {
	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse UUID: %w", err)
	}

	return &model.Station{
		ID:        id,
		Name:      dto.Name,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}, nil
}

func (dto *StationDto) FromModel(station *model.Station) {
	dto.ID = station.ID.String()
	dto.Name = station.Name
	dto.CreatedAt = station.CreatedAt
	dto.UpdatedAt = station.UpdatedAt
}

type StationRepositoryImpl struct {
	db *sqlx.DB
}

func NewStationRepository(db *sqlx.DB) repository.StationRepository {
	return &StationRepositoryImpl{
		db: db,
	}
}

func (r *StationRepositoryImpl) Save(ctx context.Context, station *model.Station) error {
	dto := &StationDto{}
	dto.FromModel(station)

	query := `
		INSERT INTO stations (id, name, created_at, updated_at)
		VALUES (:id, :name, :created_at, :updated_at)
		ON DUPLICATE KEY UPDATE
		name = VALUES(name),
		updated_at = VALUES(updated_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		return fmt.Errorf("failed to save station: %w", err)
	}

	return nil
}

func (r *StationRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*model.Station, error) {
	query := `
		SELECT *
		FROM stations
		WHERE id = ?
	`

	var dto StationDto
	err := r.db.GetContext(ctx, &dto, query, id.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("station not found")
		}

		return nil, fmt.Errorf("failed to get station: %w", err)
	}

	station, err := dto.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
	}

	return station, nil
}

func (r *StationRepositoryImpl) FindAll(ctx context.Context) ([]*model.Station, error) {
	query := `
		SELECT *
		FROM stations
	`

	var dtos []StationDto
	err := r.db.SelectContext(ctx, &dtos, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get stations: %w", err)
	}

	stations := make([]*model.Station, 0, len(dtos))
	for _, dto := range dtos {
		station, err := dto.ToModel()
		if err != nil {
			return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
		}
		stations = append(stations, station)
	}

	return stations, nil
}

func (r *StationRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM stations WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete station: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("station not found")
	}

	return nil
}
