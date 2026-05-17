package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"koala.com/internal/auth/entity"
)

type UserRepository interface {
	Save(ctx context.Context, u entity.User, r entity.Role) error
	FindById(ctx context.Context, id uuid.UUID) (entity.User, error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	UpdatePassword(ctx context.Context, u entity.User) error
}

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(Db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{Db}
}

func (userRepository *UserRepositoryImpl) Save(ctx context.Context, u entity.User, r entity.Role) error {
	insetUserQuery := "INSERT INTO users (username, password, phone_number) VALUES ($1, $2, $3) RETURNING id;"
	insertUserHasRole := "INSERT INTO user_has_role (user_id, role_id) VALUES ($1, (SELECT id from roles WHERE role_name = $2));"

	tx, err := userRepository.db.BeginTxx(ctx, nil)

	if err != nil {
		return fmt.Errorf("insert into user: %w", err)
	}

	defer tx.Rollback()

	user := entity.User{}
	
	errInsertUser := tx.GetContext(ctx, &user, insetUserQuery, u.Username, u.Password, u.PhoneNumber)

	if errInsertUser != nil {
		return fmt.Errorf("Insert into user: %w", err)
	}

	rows, errInsertUserHasRole := tx.ExecContext(ctx, insertUserHasRole, user.Id, r.RoleName)

	if errInsertUserHasRole != nil {
		return fmt.Errorf("Insert user has role: %w", errInsertUserHasRole)
	}

	if affectd, err := rows.RowsAffected(); affectd == 0 || err != nil {
		return errors.New("Insert user has role: Can not insert role for user")
	}

	errCommit := tx.Commit()

	if errCommit != nil {
		return fmt.Errorf("Insert into user and insert user has role: %w", errCommit)
	}

	return nil
}

func (userRepository *UserRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (entity.User, error) {
	user := entity.User{}

	err := userRepository.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1;", id)
	if err != nil {
		return entity.User{}, fmt.Errorf("Find by id: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	user := entity.User{}

	err := userRepository.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1;", username)
	if err != nil {
		return entity.User{}, fmt.Errorf("Find by username: %w", err)
	}

	return user, nil
}

func (userRepository *UserRepositoryImpl) FindAll(ctx context.Context) ([]entity.User, error) {
	users := []entity.User{}

	err := userRepository.db.SelectContext(ctx, &users, "SELECT * FROM users;")
	if err != nil {
		return []entity.User{}, fmt.Errorf("Insert user query: %w", err)
	}

	return users, nil
}

func (userRepository *UserRepositoryImpl) UpdatePassword(ctx context.Context, u entity.User) error {
	_, err := userRepository.db.NamedExecContext(ctx, "UPDATE users SET password = :password, updated_at = CURRENT_DATE WHERE username = :username", &u)

	if err != nil {
		return fmt.Errorf("Insert user query: %w", err)
	}
	return nil
}
