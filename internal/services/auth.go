package services

import (
	"context"
	"forum/internal/repositories"
)

type AuthService struct {
	Repository *repositories.UserRepository
}

// Authenticates a user based on email and password, returning the user id of the user
func (r *AuthService) Authenticate(ctx context.Context, email, password string) (userID int, err error) {

	// to be continued ...

	return 0, nil
}
