package rider

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RiderRepository interface{
	GetProfileByUserId(ctx context.Context, id uuid.UUID) (RiderProfile, error)
	AddProfileById(ctx context.Context, profile RiderProfile) error
	UpdateProfileById(ctx context.Context, profile RiderProfile) error
}

type RiderRepositoryImpl struct {
	db *sqlx.DB
}

func NewRiderRepository(db *sqlx.DB) *RiderRepositoryImpl {
	return &RiderRepositoryImpl{db}
}

func (riderRepository *RiderRepositoryImpl) GetProfileByUserId(ctx context.Context, id uuid.UUID) (RiderProfile, error) {
	query := "SELECT * FROM rider_profiles WHERE user_id = $1;"

	profile := RiderProfile{}

	err := riderRepository.db.GetContext(ctx, &profile, query, id)

	if err != nil {
		return RiderProfile{}, fmt.Errorf("Err Get profile by id %w", err)
	}

	return profile, nil
}

func (riderRepository *RiderRepositoryImpl) AddProfileById(ctx context.Context, profile RiderProfile) error {
	query := "INSERT INTO rider_profiles (user_id, first_name, last_name, avartar_url) VALUES (:user_id, :first_name, :last_name, :avartar_url) ON CONFLICT (user_id) DO NOTHING;"

	_, err := riderRepository.db.NamedExecContext(ctx, query, &profile)

	if err != nil {
		return fmt.Errorf("Insert into query err: %w", err)
	}

	return nil
}

func (riderRepository *RiderRepositoryImpl) UpdateProfileById(ctx context.Context, profile RiderProfile) error {
	query := "UPDATE rider_profiles SET first_name = :first_name, last_name = :last_name, avatar_url = :avatar_url WHERE user_id = :user_id;"

	_, err := riderRepository.db.NamedExecContext(ctx, query, &profile)

	if err != nil {
		return fmt.Errorf("Update query err: %w", err)
	}

	return nil
}