package services

import (
	"context"
	"database/sql"
	"errors"
	"forum/internal/constants"
	"forum/internal/uuid"
	"time"
)

// creates a session with a UUID, userID and expiration. Returns the UUID
func CreateSession(ctx context.Context, tx *sql.Tx, userID string) (string, error) {

	// ---- generate UUID -------
	UUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	// ---- define expiration -------
	expiresAt := time.Now().Add(constants.SessionExpiry * time.Minute).Format("2006-01-02 15:04:05")

	// ---- add session -------
	_, err = tx.ExecContext(ctx, constants.AddSession, UUID, userID, expiresAt)
	if err != nil {
		return "", err
	}

	return UUID, nil
}

// validates that a given session id has a valid session in the db that has not expired, returns bool
func ValidateSession(ctx context.Context, tx *sql.Tx, sessionID string) (bool, error) {

	var expiresAt string

	err := tx.QueryRowContext(ctx, constants.GetExpiryTime, sessionID).Scan(&expiresAt)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
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
