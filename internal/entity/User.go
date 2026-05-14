package entity

import "github.com/google/uuid"

type UserStatus string

const (
	ACTIVE UserStatus = "active"
	PENDING UserStatus = "pending"
	LOCKED UserStatus = "locked"
)

type User struct {
	BaseEntity[uuid.UUID]
	Username    string     `db:"username"`
	Password    string     `db:"password"`
	PhoneNumber string     `db:"phone_number"`
	Email       string     `db:"email"`
	Status      UserStatus `db:"status"`
}

func (u *User) ToString() string {
	return "| " + u.Id.String() + " | " + u.Username + " | " + string(u.Status) + " |\n"
}