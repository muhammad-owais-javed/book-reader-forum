package repository

import (
	"context"
	"database/sql"
	"forum/internal/constants"
	"forum/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {

	var exists bool
	err := r.DB.QueryRowContext(ctx, constants.CheckUniqueEmail, email).Scan(&exists)
	return exists, err

}

func (r *UserRepository) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	var exists bool
	err := r.DB.QueryRowContext(ctx, constants.CheckUniqueUserName, username).Scan(&exists)
	return exists, err
}