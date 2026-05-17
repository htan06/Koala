package shared

type UserStatus string

const (
	ACTIVE UserStatus = "ACTIVE"
	PENDING UserStatus = "PENDING"
	LOCKED UserStatus = "LOCKED"
)