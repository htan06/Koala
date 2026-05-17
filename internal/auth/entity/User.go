package entity

import (
	"github.com/google/uuid"
	"koala.com/internal/shared"
)

type User struct {
	shared.BaseEntity[uuid.UUID]
	Username    *string     `db:"username"`
	Password    *string     `db:"password"`
	PhoneNumber *string     `db:"phone_number"`
	Email       *string     `db:"email"`
	Status      *shared.UserStatus `db:"status"`
}