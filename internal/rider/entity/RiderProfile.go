package entity

import (
	"time"

	"github.com/google/uuid"
)

type RiderProfile struct {
	UserID    uuid.UUID `db:"user_id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	AvatarUrl string 	`db:"avatar_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
