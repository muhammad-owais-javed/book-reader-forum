package services

import (
	"context"
	"forum/internal/auth/repository"
	"forum/internal/constants"
	"forum/internal/uuid"
	"time"
)

type ResetTokenService struct {
	ResetTokenService    *ResetTokenService
	ResetTokenRepository *repository.ResetTokenRepository
	UserService          *UserService
	UserRepository       *repository.UserRepository
}

func NewResetTokenService(resetTokenRepo *repository.ResetTokenRepository, userRepo *repository.UserRepository) *ResetTokenService {
	return &ResetTokenService{ResetTokenRepository: resetTokenRepo, UserRepository: userRepo}
}

// creates a reset token and stores it with user ID and expiry time in reset_tokens table
func (s *ResetTokenService) CreateResetToken(ctx context.Context, email string) error {

	// ---- generate reset token with UUID v4 syntax -------
	resetToken, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	// ---- define expiration -------
	expiresAt := time.Now().Add(constants.ResetTokenExpiry * time.Minute).Format("2006-01-02 15:04:05")

	// ---- get user ID by email -------
	userID, err := s.UserRepository.GetUserIDByEmail(ctx, email)
	if err != nil {
		return err
	}

	// ---- generate reset token -------
	err = s.ResetTokenRepository.InsertResetToken(ctx, resetToken, userID, expiresAt)
	if err != nil {
		return err
	}

	return err
}
