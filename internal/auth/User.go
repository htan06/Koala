package auth

import (
	"github.com/google/uuid"
	"koala.com/internal/shared"
)
type UserStatus string

const (
	ACTIVE UserStatus = "active"
	PENDING UserStatus = "pending"
	LOCKED UserStatus = "locked"
)

type User struct {
	shared.BaseEntity[uuid.UUID]
	Username    string     `db:"username"`
	Password    string     `db:"password"`
	PhoneNumber string     `db:"phone_number"`
	Email       string     `db:"email"`
	Status      UserStatus `db:"status"`
}