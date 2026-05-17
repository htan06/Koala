package entity

import "koala.com/internal/shared"

type Role struct {
	shared.BaseEntity[int64]
	RoleName string `db:"role_name"`
}
