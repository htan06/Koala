package request

type AddProfileDto struct {
	FirstName                 *string       `json:"first_name"`
	LastName                  *string       `json:"last_name"`
	AvatarURL                 *string       `json:"avatar_url"`
	NationalIDNumber          *string       `json:"national_id_number"`
	DriverLicenseNumber       *string       `json:"driver_license_number"`
	VehicleRegistrationNumber *string       `json:"vehicle_registration_number"`
}