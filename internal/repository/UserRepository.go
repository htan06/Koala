package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"koala.com/internal/entity"
)

type UserRepository struct {
	Db *sqlx.DB
}

func NewUserRepository(Db *sqlx.DB) *UserRepository {
	return &UserRepository{Db}
}

func (u *UserRepository) FindAll() []entity.User {
	users := []entity.User{}

	err := u.Db.Select(&users, "SELECT * FROM users;")
	if err != nil {
		log.Fatal(err)
	}

	return users
}