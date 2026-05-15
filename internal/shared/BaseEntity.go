package shared

import (
	"time"
)

type BaseEntity[T comparable] struct {
	Id        T         `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}