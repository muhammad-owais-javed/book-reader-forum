package services

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailDoesntExist = errors.New("email doesnt exist")
var ErrWrongPassword = errors.New("wrong password")

// Authenticates a user based on email and password, returning the user id of the user
func (s *UserService) Authenticate(ctx context.Context, email, password string) (string, error) {

	// var realHashedPassword string
	// err = tx.QueryRowContext(ctx, constants.GetUserIDAndPasswordByEmail, email).Scan(&userID, &realHashedPassword)
	userID, realHashedPassword, err := s.repo.GetUserCredentialsByEmail(ctx, email)
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
