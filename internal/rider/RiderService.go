package rider

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type RiderService interface {
	GetProfile(ctx context.Context, userId uuid.UUID) (RiderProfile, error)
	AddProfile(ctx context.Context, profile RiderProfile) error
	UpdateProfile(ctx context.Context, profile RiderProfile) error
}

type RiderServiceImpl struct {
	riderProfileRepository RiderRepository
}

func NewRiderService(riderProfileRepository RiderRepository) *RiderServiceImpl {
	return &RiderServiceImpl{riderProfileRepository}
}

func (riderService *RiderServiceImpl) GetProfile(ctx context.Context, userId uuid.UUID) (RiderProfile, error) {
	profile, err := riderService.riderProfileRepository.GetProfileByUserId(ctx, userId)

	if err != nil {
		return RiderProfile{}, fmt.Errorf("Get rider profile: %w", err)
	}

	return profile, nil
}

func (riderService *RiderServiceImpl) AddProfile(ctx context.Context, profile RiderProfile) error {
	err := riderService.riderProfileRepository.AddProfileById(ctx, profile)

	if err != nil {
		return fmt.Errorf("Add rider profile: %w", err)
	}

	return nil
}

func (riderService *RiderServiceImpl) UpdateProfile(ctx context.Context, profile RiderProfile) error {
	err := riderService.riderProfileRepository.UpdateProfileById(ctx, profile)

	if err != nil {
		return fmt.Errorf("Update rider profile: %w", err)
	}

	return nil
}
