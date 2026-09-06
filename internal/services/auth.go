package services

import (
	"context"
	"database/sql"
	"errors"
	"forum/internal/constants"
	"forum/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repository *repositories.UserRepository
}

var ErrEmailDoesntExist = errors.New("email doesnt exist")
var ErrWrongPassword = errors.New("wrong password")

// Authenticates a user based on email and password, returning the user id of the user
func (r *AuthService) Authenticate(ctx context.Context, email, password string) (userID string, err error) {

	tx, err := r.Repository.DB.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	var realHashedPassword string
	err = tx.QueryRowContext(ctx, constants.GetUserIDAndPasswordByEmail, email).Scan(&userID, &realHashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrEmailDoesntExist
		}
		return "", err
	}

	// compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(realHashedPassword), []byte(password))
	if err != nil {
		return "", ErrWrongPassword
	}

	return userID, nil
}
