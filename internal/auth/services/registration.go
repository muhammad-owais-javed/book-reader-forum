package services

import (
	"context"
	// "database/sql"
	"errors"
	// "forum/internal/constants"
	"forum/internal/uuid"
	"forum/internal/auth/repository"
	"forum/internal/auth/models"
	"golang.org/x/crypto/bcrypt"
)


type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

var ErrEmailExists = errors.New("email already exists")
var ErrUsernameExists = errors.New("username already exists")

// Registers a new user into the database or returns and error if username has been taken or email has already been registered.
func (s *UserService) Register(ctx context.Context, username, email, password string) error {

	// ---- email check -------
	// var emailExistsAlready bool
	// err = tx.QueryRowContext(ctx, constants.CheckUniqueEmail, email).Scan(&emailExistsAlready)
	emailExists, err := s.repo.CheckEmailExists(ctx, email)
	if err != nil {
		return err
	}

	if emailExists == true {
		return ErrEmailExists
	}

	// ---- username check -------
	// var usernameExistsAlready bool
	// err = tx.QueryRowContext(ctx, constants.CheckUniqueUserName, username).Scan(&usernameExistsAlready)
	usernameExists, err := s.repo.CheckUsernameExists(ctx, username)
	if err != nil {
		return err
	}
	if usernameExists == true {
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

	user := &models.User{
		ID:           UUID,
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	// ---- create user -------
	// _, err = tx.ExecContext(ctx, constants.CreateUser, UUID, username, email, hashedPassword)
	// if err != nil {
	// 	return
	// }

	return s.repo.CreateUser(ctx, user)
}
