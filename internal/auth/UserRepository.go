package auth

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(Db *sqlx.DB) *UserRepository {
	return &UserRepository{Db}
}

func (userRepository *UserRepository) Save(ctx context.Context, u User) error {
	_, err := userRepository.db.NamedExecContext(ctx, "INSERT INTO users (username, password, phone_number) VALUES (:username, :password, :phone_number);", &u)

	if err != nil {
		return fmt.Errorf("Insert user query: %w", err)
	}
	return nil
}

func (userRepository *UserRepository) FindByUsername(ctx context.Context, username string) (User, error) {
	user := User{}

	err := userRepository.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1;", username)
	if err != nil {
		return User{}, fmt.Errorf("Find by username: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepository) FindAll(ctx context.Context) ([]User, error) {
	users := []User{}

	err := userRepository.db.SelectContext(ctx, &users, "SELECT * FROM users;")
	if err != nil {
		return []User{}, fmt.Errorf("Insert user query: %w", err)
	}

	return users, nil
}

func (userRepository *UserRepository) updatePassword(ctx context.Context, u User) error {
	_, err := userRepository.db.NamedExecContext(ctx, "UPDATE table users SET password = :password WHERE username = :username", &u)

	if err != nil {
		return fmt.Errorf("Insert user query: %w", err)
	}
	return nil
}
