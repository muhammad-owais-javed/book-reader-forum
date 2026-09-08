package services

import (
	"context"
	"database/sql"
	"errors"
	"forum/internal/constants"
	"forum/internal/repository"
	"forum/internal/uuid"
	"time"
)


type SessionService struct {
	repo *repository.SessionRepository
}

func NewSessionService(repo *repository.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

// creates a session with a UUID, userID and expiration. Returns the UUID
func (s *SessionService) CreateSession(ctx context.Context, userID string) (string, error) {

	// ---- generate UUID -------
	sessionID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	// ---- define expiration -------
	expiresAt := time.Now().Add(constants.SessionExpiry * time.Minute).Format("2006-01-02 15:04:05")

	// ---- add session -------
	// _, err = tx.ExecContext(ctx, constants.AddSession, UUID, userID, expiresAt)
	err = s.repo.InsertSession(ctx, sessionID, userID, expiresAt)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

// validates that a given session id has a valid session in the db that has not expired, returns bool
func ValidateSession(ctx context.Context, tx *sql.Tx, sessionID string) (bool, error) {

	// var expiresAt string
	// err := tx.QueryRowContext(ctx, constants.GetExpiryTime, sessionID).Scan(&expiresAt)
	
	expiresAt, err := s.repo.GetSessionExpiry(ctx, sessionID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil 
		}
		return false, err 
	}

	expiry, err := time.Parse(time.RFC3339, expiresAt) // expiresAt is stored in YYYY-MM-DD HH:MM:SS format but Scan reads it to RCF3339
	if err != nil {
		return false, err
	}

	if expiry.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}
