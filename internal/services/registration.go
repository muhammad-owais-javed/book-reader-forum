package services

import (
	"context"
	"errors"
	"forum/internal/constants"
	"forum/internal/repositories"
	"forum/internal/uuid"

	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	Repository *repositories.UserRepository
}

var ErrEmailExists = errors.New("email already exists")
var ErrUsernameExists = errors.New("username already exists")

// Registers a new user into the database or returns and error if username has been taken or email has already been registered.
func (r *RegistrationService) Register(ctx context.Context, username, email, password string) (err error) {

	tx, err := r.Repository.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// ---- email check -------
	var emailExistsAlready bool
	err = tx.QueryRowContext(ctx, constants.CheckUniqueEmail, email).Scan(&emailExistsAlready)
	if err != nil {
		return err
	}

	if emailExistsAlready {
		return ErrEmailExists
	}

	// ---- username check -------
	var usernameExistsAlready bool
	err = tx.QueryRowContext(ctx, constants.CheckUniqueUserName, username).Scan(&usernameExistsAlready)
	if err != nil {
		return err
	}
	if usernameExistsAlready {
		return ErrUsernameExists
	}

	// ---- bcrypt hash password -------
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// ---- generate UUID -------
	UUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	// ---- create user -------
	_, err = tx.ExecContext(ctx, constants.CreateUser, UUID, username, email, hashedPassword)
	if err != nil {
		return
	}

	tx.Commit()

	return nil
}
