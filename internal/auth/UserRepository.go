package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Save(ctx context.Context, u User) error
	FindById(ctx context.Context, id uuid.UUID) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	UpdatePassword(ctx context.Context, u User) error
}

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(Db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{Db}
}

func (userRepository *UserRepositoryImpl) Save(ctx context.Context, u User) error {
	_, err := userRepository.db.NamedExecContext(ctx, "INSERT INTO users (username, password, phone_number) VALUES (:username, :password, :phone_number);", &u)

	if err != nil {
		return fmt.Errorf("Insert user query: %w", err)
	}
	return nil
}

func (userRepository *UserRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (User, error) {
	user := User{}

	err := userRepository.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1;", id)
	if err != nil {
		return User{}, fmt.Errorf("Find by id: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (User, error) {
	user := User{}

	err := userRepository.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1;", username)
	if err != nil {
		return User{}, fmt.Errorf("Find by username: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepositoryImpl) FindAll(ctx context.Context) ([]User, error) {
	users := []User{}

	err := userRepository.db.SelectContext(ctx, &users, "SELECT * FROM users;")
	if err != nil {
		return []User{}, fmt.Errorf("Insert user query: %w", err)
	}

	return users, nil
}

func (userRepository *UserRepositoryImpl) UpdatePassword(ctx context.Context, u User) error {
	_, err := userRepository.db.NamedExecContext(ctx, "UPDATE users SET password = :password WHERE username = :username", &u)

	if err != nil {
		return fmt.Errorf("Insert user query: %w", err)
	}
	return nil
}
