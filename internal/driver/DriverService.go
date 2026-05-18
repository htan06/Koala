package driver

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"koala.com/internal/driver/dto/request"
	"koala.com/internal/driver/entity"
	"koala.com/internal/shared"
)

type DriverService interface {
	AddProfile(ctx context.Context, profileDto request.AddProfileDto) error
	UpdateProfile(ctx context.Context, profileDto request.UpdateProfileDto) error
	GetProfileByUserId(ctx context.Context, userId uuid.UUID) (entity.DriverProfile, error)
	GetListProfileByStatus(ctx context.Context, status entity.DriverStatus, limit uint, offset uint) ([]entity.DriverProfile, error)
}

type DriverServiceImpl struct {
	driverRepository DriverRepository
}

func NewDriverService(driverRepository DriverRepository) *DriverServiceImpl {
	return &DriverServiceImpl{driverRepository}
}

func (driverService *DriverServiceImpl) AddProfile(ctx context.Context, profileDto request.AddProfileDto) error {

	profile := entity.DriverProfile{
		FirstName:                 profileDto.FirstName,
		LastName:                  profileDto.LastName,
		AvatarURL:                 profileDto.AvatarURL,
		NationalIDNumber:          profileDto.NationalIDNumber,
		DriverLicenseNumber:       profileDto.DriverLicenseNumber,
		VehicleRegistrationNumber: profileDto.VehicleRegistrationNumber,
		Status:                    entity.PENDING,
	}

	err := driverService.driverRepository.AddProfile(ctx, profile)

	if err != nil {
		return fmt.Errorf("Add profile err %w", err)
	}

	return nil
}

func (driverService *DriverServiceImpl) UpdateProfile(ctx context.Context, profileDto request.UpdateProfileDto) error {

	profileUpdate := entity.DriverProfile{
		BaseEntity: shared.BaseEntity[int64]{Id: profileDto.Id},
		UserId:     profileDto.UserId,
		Status:     *profileDto.Status,
	}

	err := driverService.driverRepository.UpdateProfileById(ctx, profileUpdate)

	if err != nil {
		return fmt.Errorf("Update profile err: %w", err)
	}

	return nil
}

func (driverService *DriverServiceImpl) GetProfileByUserId(ctx context.Context, userId uuid.UUID) (entity.DriverProfile, error) {
	profile, err := driverService.driverRepository.GetProfileByUserId(ctx, userId)

	if err != nil {
		return entity.DriverProfile{}, fmt.Errorf("Get profile by id: %w", err)
	}

	return profile, nil
}

func (driverService *DriverServiceImpl) GetListProfileByStatus(ctx context.Context, status entity.DriverStatus, limit uint, offset uint) ([]entity.DriverProfile, error) {
	listProfile, err := driverService.driverRepository.FindAllProfileByStatus(ctx, status, limit, offset)

	if err != nil {
		return []entity.DriverProfile{}, fmt.Errorf("Get profile by id: %w", err)
	}

	return listProfile, nil
}
