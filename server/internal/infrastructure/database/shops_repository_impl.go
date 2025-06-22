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

type ShopDto struct {
	ID         string    `db:"id"`
	Name       string    `db:"name"`
	PostCode   string    `db:"post_code"`
	Address    string    `db:"address"`
	Latitude   string    `db:"latitude"`
	Longitude  string    `db:"longitude"`
	Registerer string    `db:"registerer"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type ShopStationDto struct {
	ShopID    string `db:"shop_id"`
	StationID string `db:"station_id"`
}

type ShopPaymentMethodDto struct {
	ShopID        string `db:"shop_id"`
	PaymentMethod string `db:"payment_method"`
}

type ShopImageDto struct {
	ShopID  string `db:"shop_id"`
	ImageID string `db:"image_id"`
}

func (dto *ShopDto) ToModel(stations []uuid.UUID, images []model.ImageFile, paymentMethods []string) (*model.Shop, error) {
	id, err := uuid.Parse(dto.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse shop UUID: %w", err)
	}

	shopName, err := model.NewShopName(dto.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create ShopName: %w", err)
	}

	var postCode model.PostCode
	if dto.PostCode != "" {
		postCode, err = model.NewPostCode(dto.PostCode)
		if err != nil {
			return nil, fmt.Errorf("failed to create PostCode: %w", err)
		}
	}

	var registerer model.UserID
	if dto.Registerer != "" {
		registerer, err = model.NewUserID(dto.Registerer)
		if err != nil {
			return nil, fmt.Errorf("failed to create UserID: %w", err)
		}
	}

	var latitude, longitude float64
	if dto.Latitude != "" {
		if _, err := fmt.Sscanf(dto.Latitude, "%f", &latitude); err != nil {
			return nil, fmt.Errorf("failed to parse latitude: %w", err)
		}
	}
	if dto.Longitude != "" {
		if _, err := fmt.Sscanf(dto.Longitude, "%f", &longitude); err != nil {
			return nil, fmt.Errorf("failed to parse longitude: %w", err)
		}
	}

	return &model.Shop{
		ID:             id,
		Name:           shopName,
		PostCode:       postCode,
		Address:        dto.Address,
		Latitude:       latitude,
		Longitude:      longitude,
		Stations:       stations,
		Images:         images,
		PaymentMethods: paymentMethods,
		Registerer:     registerer,
		CreatedAt:      dto.CreatedAt,
		UpdatedAt:      dto.UpdatedAt,
	}, nil
}

func (dto *ShopDto) FromModel(shop *model.Shop) {
	dto.ID = shop.ID.String()
	dto.Name = string(shop.Name)
	dto.PostCode = string(shop.PostCode)
	dto.Address = shop.Address
	dto.Latitude = fmt.Sprintf("%f", shop.Latitude)
	dto.Longitude = fmt.Sprintf("%f", shop.Longitude)
	dto.Registerer = string(shop.Registerer)
	dto.CreatedAt = shop.CreatedAt
	dto.UpdatedAt = shop.UpdatedAt
}

type ShopRepositoryImpl struct {
	db *sqlx.DB
}

func NewShopRepository(db *sqlx.DB) repository.ShopRepository {
	return &ShopRepositoryImpl{
		db: db,
	}
}

func (r *ShopRepositoryImpl) Save(ctx context.Context, shop *model.Shop) error {
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

	// Save shop
	dto := &ShopDto{}
	dto.FromModel(shop)

	query := `
INSERT INTO shops (id, name, post_code, address, latitude, longitude, registerer, created_at, updated_at)
VALUES (:id, :name, :post_code, :address, :latitude, :longitude, :registerer, :created_at, :updated_at)
ON DUPLICATE KEY UPDATE
name = VALUES(name),
post_code = VALUES(post_code),
address = VALUES(address),
latitude = VALUES(latitude),
longitude = VALUES(longitude),
registerer = VALUES(registerer),
updated_at = VALUES(updated_at)
`

	_, err = tx.NamedExecContext(ctx, query, dto)
	if err != nil {
		return fmt.Errorf("failed to save shop: %w", err)
	}

	// Delete existing shop stations
	deleteStationsQuery := `DELETE FROM shop_stations WHERE shop_id = ?`
	_, err = tx.ExecContext(ctx, deleteStationsQuery, shop.ID.String())
	if err != nil {
		return fmt.Errorf("failed to delete existing shop stations: %w", err)
	}

	// Save shop stations
	if len(shop.Stations) > 0 {
		stationQuery := `INSERT INTO shop_stations (shop_id, station_id) VALUES (?, ?)`
		for _, stationID := range shop.Stations {
			_, err = tx.ExecContext(ctx, stationQuery, shop.ID.String(), stationID.String())
			if err != nil {
				return fmt.Errorf("failed to save shop station: %w", err)
			}
		}
	}

	// Delete existing shop payment methods
	deletePaymentMethodsQuery := `DELETE FROM shop_payment_methods WHERE shop_id = ?`
	_, err = tx.ExecContext(ctx, deletePaymentMethodsQuery, shop.ID.String())
	if err != nil {
		return fmt.Errorf("failed to delete existing shop payment methods: %w", err)
	}

	// Save shop payment methods
	if len(shop.PaymentMethods) > 0 {
		paymentMethodQuery := `INSERT INTO shop_payment_methods (shop_id, payment_method) VALUES (?, ?)`
		for _, paymentMethod := range shop.PaymentMethods {
			_, err = tx.ExecContext(ctx, paymentMethodQuery, shop.ID.String(), paymentMethod)
			if err != nil {
				return fmt.Errorf("failed to save shop payment method: %w", err)
			}
		}
	}

	// Delete existing shop images
	deleteImagesQuery := `DELETE FROM shop_images WHERE shop_id = ?`
	_, err = tx.ExecContext(ctx, deleteImagesQuery, shop.ID.String())
	if err != nil {
		return fmt.Errorf("failed to delete existing shop images: %w", err)
	}

	// Save shop images
	if len(shop.Images) > 0 {
		imageQuery := `INSERT INTO shop_images (shop_id, image_id) VALUES (?, ?)`
		for _, image := range shop.Images {
			_, err = tx.ExecContext(ctx, imageQuery, shop.ID.String(), image.ID.String())
			if err != nil {
				return fmt.Errorf("failed to save shop image: %w", err)
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *ShopRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*model.Shop, error) {
	query := `
SELECT *
FROM shops
WHERE id = ?
`

	var dto ShopDto
	err := r.db.GetContext(ctx, &dto, query, id.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("shop not found")
		}

		return nil, fmt.Errorf("failed to get shop: %w", err)
	}

	// Get shop stations
	stations, err := r.getShopStations(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get shop stations: %w", err)
	}

	// Get shop images
	images, err := r.getShopImages(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get shop images: %w", err)
	}

	// Get shop payment methods
	paymentMethods, err := r.getShopPaymentMethods(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get shop payment methods: %w", err)
	}

	shop, err := dto.ToModel(stations, images, paymentMethods)
	if err != nil {
		return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
	}

	return shop, nil
}

func (r *ShopRepositoryImpl) FindAll(ctx context.Context) ([]*model.Shop, error) {
	query := `
SELECT *
FROM shops
`

	var dtos []ShopDto
	err := r.db.SelectContext(ctx, &dtos, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get shops: %w", err)
	}

	shops := make([]*model.Shop, 0, len(dtos))
	for _, dto := range dtos {
		// Get shop details for each shop
		shopID, err := uuid.Parse(dto.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse shop ID: %w", err)
		}

		stations, err := r.getShopStations(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop stations: %w", err)
		}

		images, err := r.getShopImages(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop images: %w", err)
		}

		paymentMethods, err := r.getShopPaymentMethods(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop payment methods: %w", err)
		}

		shop, err := dto.ToModel(stations, images, paymentMethods)
		if err != nil {
			return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
		}
		shops = append(shops, shop)
	}

	return shops, nil
}

func (r *ShopRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
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

	// Delete shop stations
	deleteStationsQuery := `DELETE FROM shop_stations WHERE shop_id = ?`
	_, err = tx.ExecContext(ctx, deleteStationsQuery, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete shop stations: %w", err)
	}

	// Delete shop payment methods
	deletePaymentMethodsQuery := `DELETE FROM shop_payment_methods WHERE shop_id = ?`
	_, err = tx.ExecContext(ctx, deletePaymentMethodsQuery, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete shop payment methods: %w", err)
	}

	// Delete shop images
	deleteImagesQuery := `DELETE FROM shop_images WHERE shop_id = ?`
	_, err = tx.ExecContext(ctx, deleteImagesQuery, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete shop images: %w", err)
	}

	// Delete shop
	deleteShopQuery := `DELETE FROM shops WHERE id = ?`
	result, err := tx.ExecContext(ctx, deleteShopQuery, id.String())
	if err != nil {
		return fmt.Errorf("failed to delete shop: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("shop not found")
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *ShopRepositoryImpl) FindByStation(ctx context.Context, stationID uuid.UUID) ([]*model.Shop, error) {
	query := `
SELECT s.*
FROM shops s
INNER JOIN shop_stations ss ON s.id = ss.shop_id
WHERE ss.station_id = ?
`

	var dtos []ShopDto
	err := r.db.SelectContext(ctx, &dtos, query, stationID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get shops by station: %w", err)
	}

	shops := make([]*model.Shop, 0, len(dtos))
	for _, dto := range dtos {
		// Get shop details for each shop
		shopID, err := uuid.Parse(dto.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse shop ID: %w", err)
		}

		stations, err := r.getShopStations(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop stations: %w", err)
		}

		images, err := r.getShopImages(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop images: %w", err)
		}

		paymentMethods, err := r.getShopPaymentMethods(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop payment methods: %w", err)
		}

		shop, err := dto.ToModel(stations, images, paymentMethods)
		if err != nil {
			return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
		}
		shops = append(shops, shop)
	}

	return shops, nil
}

func (r *ShopRepositoryImpl) getShopStations(ctx context.Context, shopID uuid.UUID) ([]uuid.UUID, error) {
	query := `
SELECT station_id
FROM shop_stations
WHERE shop_id = ?
`

	var stationIDStrs []string
	err := r.db.SelectContext(ctx, &stationIDStrs, query, shopID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get shop station IDs: %w", err)
	}

	stationIDs := make([]uuid.UUID, 0, len(stationIDStrs))
	for _, stationIDStr := range stationIDStrs {
		stationID, err := uuid.Parse(stationIDStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse station ID: %w", err)
		}
		stationIDs = append(stationIDs, stationID)
	}

	return stationIDs, nil
}

func (r *ShopRepositoryImpl) getShopImages(ctx context.Context, shopID uuid.UUID) ([]model.ImageFile, error) {
	query := `
SELECT image_id
FROM shop_images
WHERE shop_id = ?
`

	var imageIDs []string
	err := r.db.SelectContext(ctx, &imageIDs, query, shopID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get shop image IDs: %w", err)
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

func (r *ShopRepositoryImpl) getShopPaymentMethods(ctx context.Context, shopID uuid.UUID) ([]string, error) {
	query := `
SELECT payment_method
FROM shop_payment_methods
WHERE shop_id = ?
`

	var paymentMethods []string
	err := r.db.SelectContext(ctx, &paymentMethods, query, shopID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get shop payment methods: %w", err)
	}

	return paymentMethods, nil
}

func (r *ShopRepositoryImpl) FindAllWithLimit(ctx context.Context, limit int, offset int) ([]*model.Shop, error) {
	query := `
	SELECT *
	FROM shops
	LIMIT ? OFFSET ?
`

	var dtos []ShopDto
	err := r.db.SelectContext(ctx, &dtos, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get shops with limit: %w", err)
	}

	shops := make([]*model.Shop, 0, len(dtos))
	for _, dto := range dtos {
		shopID, err := uuid.Parse(dto.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse shop ID: %w", err)
		}

		stations, err := r.getShopStations(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop stations: %w", err)
		}

		images, err := r.getShopImages(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop images: %w", err)
		}

		paymentMethods, err := r.getShopPaymentMethods(ctx, shopID)
		if err != nil {
			return nil, fmt.Errorf("failed to get shop payment methods: %w", err)
		}

		shop, err := dto.ToModel(stations, images, paymentMethods)
		if err != nil {
			return nil, fmt.Errorf("failed to convert DTO to model: %w", err)
		}
		shops = append(shops, shop)
	}

	return shops, nil
}
