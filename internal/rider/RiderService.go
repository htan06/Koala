package rider

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"koala.com/internal/auth"
	"koala.com/internal/rider/dto/request"
	"koala.com/internal/rider/entity"
	"koala.com/internal/shared"
)

type RiderService interface {
	Register(ctx context.Context, registerDto shared.RegisterDto) error
	GetProfile(ctx context.Context, userId uuid.UUID) (entity.RiderProfile, error)
	AddProfile(ctx context.Context, userId uuid.UUID, profileDto request.ProfileDto) error
	UpdateProfile(ctx context.Context, userId uuid.UUID, profileDto request.ProfileDto) error
}

type RiderServiceImpl struct {
	riderProfileRepository RiderRepository
	authService auth.AuthService
}

func NewRiderService(riderProfileRepository RiderRepository, authService auth.AuthService) *RiderServiceImpl {
	return &RiderServiceImpl{riderProfileRepository, authService}
}

func (riderService *RiderServiceImpl) Register(ctx context.Context, registerDto shared.RegisterDto) error {
	err := riderService.authService.Register(ctx, registerDto, "RIDER", shared.ACTIVE)

	if err != nil {
		return fmt.Errorf("Register rider: %w", err)
	}

	return nil
}

func (riderService *RiderServiceImpl) GetProfile(ctx context.Context, userId uuid.UUID) (entity.RiderProfile, error) {
	profile, err := riderService.riderProfileRepository.GetProfileByUserId(ctx, userId)

	if err != nil {
		return entity.RiderProfile{}, fmt.Errorf("Get rider profile: %w", err)
	}

	return profile, nil
}

func (riderService *RiderServiceImpl) AddProfile(ctx context.Context, userId uuid.UUID, profileDto request.ProfileDto) error {
	profile := entity.RiderProfile{
		UserID: &userId,
		FirstName: &profileDto.FirstName,
		LastName: &profileDto.LastName,
		AvatarUrl: &profileDto.AvatarUrl,
	}

	err := riderService.riderProfileRepository.AddProfileById(ctx, profile)

	if err != nil {
		return fmt.Errorf("Add rider profile: %w", err)
	}

	return nil
}

func (riderService *RiderServiceImpl) UpdateProfile(ctx context.Context, userId uuid.UUID, profileDto request.ProfileDto) error {
	profile := entity.RiderProfile{
		UserID: &userId,
		FirstName: &profileDto.FirstName,
		LastName: &profileDto.LastName,
		AvatarUrl: &profileDto.AvatarUrl,
	}
	
	err := riderService.riderProfileRepository.UpdateProfileById(ctx, profile)

	if err != nil {
		return fmt.Errorf("Update rider profile: %w", err)
	}

	return nil
}
