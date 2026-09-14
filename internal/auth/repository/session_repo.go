package repository

import (
	"context"
	"database/sql"
	"forum/internal/constants"
)

type SessionRepository struct {
	DB *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{DB: db}
}

func (r *SessionRepository) InsertSession(ctx context.Context, sessionID, userID, expiresAt string) error {

	_, err := r.DB.ExecContext(ctx, constants.AddSession, sessionID, userID, expiresAt)
	
	return err

}


func (r *SessionRepository) GetSessionExpiry(ctx context.Context, sessionID string) (string, error) {

	var expiresAt string

	err := r.DB.QueryRowContext(ctx, constants.GetExpiryTime, sessionID).Scan(&expiresAt)
	
	return expiresAt, err

}


func (r *SessionRepository) GetSessionDetails(ctx context.Context, sessionID string) (string, string, error) {
	var userID, expiresAt string
	
	query := `SELECT user_id, expires_at FROM sessions WHERE id = ?`

	err := r.DB.QueryRowContext(ctx, query, sessionID).Scan(&userID, &expiresAt)
	
	return userID, expiresAt, err
}
