package response

import (
	"github.com/google/uuid"
	"koala.com/internal/driver/entity"
)

type DriverProfileDto struct {
	Id                        *int64              `json:"id"`
	UserId                    *uuid.UUID          `json:"user_id"`
	FirstName                 *string             `json:"first_name"`
	LastName                  *string             `json:"last_name"`
	AvatarURL                 *string             `json:"avatar_url"`
	NationalIDNumber          *string             `json:"national_id_number"`
	DriverLicenseNumber       *string             `json:"driver_license_number"`
	VehicleRegistrationNumber *string             `json:"vehicle_registration_number"`
	Status                    entity.DriverStatus `json:"status"`
}
