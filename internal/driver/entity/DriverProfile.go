package entity

import (
	"github.com/google/uuid"
	"koala.com/internal/shared"
)

type DriverStatus string

const (
	PENDING  DriverStatus = "PENDING"
	REJECTED DriverStatus = "REJECTED"
	WORKING  DriverStatus = "WORKING"
	RETIRED  DriverStatus = "RETIRED"
)

type DriverProfile struct {
	shared.BaseEntity[int64]
	UserId                    *uuid.UUID   `db:"user_id"`
	FirstName                 *string      `db:"first_name"`
	LastName                  *string      `db:"last_name"`
	AvatarURL                 *string      `db:"avatar_url"`
	NationalIDNumber          *string      `db:"national_id_number"`
	DriverLicenseNumber       *string      `db:"driver_license_number"`
	VehicleRegistrationNumber *string      `db:"vehicle_registration_number"`
	Status                    DriverStatus `db:"status"`
}
