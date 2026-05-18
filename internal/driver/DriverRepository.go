package driver

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"koala.com/internal/driver/entity"
)

type DriverRepository interface {
	AddProfile(ctx context.Context, profile entity.DriverProfile) error
	GetProfileByUserId(ctx context.Context, id uuid.UUID) (entity.DriverProfile, error)
	UpdateProfileById(ctx context.Context, profile entity.DriverProfile) error
	FindAllProfileByStatus(ctx context.Context, status entity.DriverStatus, limit uint, offset uint) ([]entity.DriverProfile, error)
}

type DriverRepositoryImpl struct {
	db *sqlx.DB
}

func NewDriverRepository(db *sqlx.DB) *DriverRepositoryImpl {
	return &DriverRepositoryImpl{db}
}

func (driverRepository *DriverRepositoryImpl) AddProfile(ctx context.Context, profile entity.DriverProfile) error {
	query := "INSERT INTO driver_profiles (user_id, first_name, last_name, avatar_url, national_id_number, driver_license_number, vehicle_registration_number, status) " + 
				"VALUES (:user_id, :first_name, :last_name, :avatar_url, :national_id_number, :driver_license_number, :vehicle_registration_number, :status);"

	res, err := driverRepository.db.NamedExecContext(ctx, query, profile)

	if err != nil {
		return fmt.Errorf("Insert into driver profile: %w", err)
	}

	if affected, err := res.RowsAffected(); affected == 0 || err != nil {
		return errors.New("Can not insert profile")
	}

	return nil
}

func (driverRepository *DriverRepositoryImpl) GetProfileByUserId(ctx context.Context, id uuid.UUID) (entity.DriverProfile, error) {
	query := "SELECT * FROM driver_profiles WHERE user_id = $1;"

	profile := entity.DriverProfile{}

	err := driverRepository.db.GetContext(ctx, &profile, query, id)

	if err != nil {
		return entity.DriverProfile{}, fmt.Errorf("Get profile by id %w", err)
	}

	return profile, nil
}

func (driverRepository *DriverRepositoryImpl) UpdateProfileById(ctx context.Context, profile entity.DriverProfile) error {
	query := "UPDATE driver_profiles SET user_id = :user_id, status = :status, updated_at = CURRENT_TIMESTAMP WHERE id = :id;"

	res, err := driverRepository.db.NamedExecContext(ctx, query, profile)

	if err != nil {
		return fmt.Errorf("Upadate profile by id: %w", err)
	}

	if affected, err := res.RowsAffected(); err != nil || affected == 0 {
		return errors.New("Can not update profile")
	}

	return nil
}

func (driverRepository *DriverRepositoryImpl) FindAllProfileByStatus(ctx context.Context, status entity.DriverStatus, limit uint, offset uint) ([]entity.DriverProfile, error) {
	query := "SELECT * FROM driver_profiles WHERE status = $1 ORDER BY created_at LIMIT $2 OFFSET $3;"

	listProfile := []entity.DriverProfile{}
	err := driverRepository.db.SelectContext(ctx, &listProfile, query, status, limit, offset)

	if err != nil {
		return []entity.DriverProfile{}, fmt.Errorf("Get list profile by status: %w", err)
	}

	return listProfile, nil
}