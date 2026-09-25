package services

import (
	"context"
	"database/sql"
	"errors"
	"forum/internal/auth/repository"
	"forum/internal/constants"
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
	err = s.repo.InsertSession(ctx, sessionID, userID, expiresAt)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

// validates that a given session id has a valid session in the db that has not expired, returns bool
func (s *SessionService) ValidateSession(ctx context.Context, sessionID string) (string, error) {

	userID, expiresAt, err := s.repo.GetSessionDetails(ctx, sessionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil // Session not found
		}
		return "", err // Database error
	}

	expiry, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {

		expiry, err = time.Parse("2006-01-02 15:04:05", expiresAt)
		if err != nil {
			return "", err
		}
	}

	if expiry.Before(time.Now()) {
		return "", nil // Session expired
	}

	return userID, nil
}

// deletes the session with a given UUID
func (s *SessionService) DeleteSession(ctx context.Context, sessionID string) (err error) {

	// ---- delete session -------
	err = s.repo.DeleteSession(ctx, sessionID)
	if err != nil {
		return err
	}

	return nil
}
