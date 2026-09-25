package repository

import (
	"context"
	"database/sql"
	"forum/internal/constants"
)

type ResetTokenRepository struct {
	DB *sql.DB
}

func NewResetTokenRepository(db *sql.DB) *ResetTokenRepository {
	return &ResetTokenRepository{DB: db}
}

func (r *ResetTokenRepository) InsertResetToken(ctx context.Context, resetTokenID, userID, expiresAt string) error {

	_, err := r.DB.ExecContext(ctx, constants.AddResetToken, resetTokenID, userID, expiresAt)

	return err

}

func (r *ResetTokenRepository) GetResetTokenExpiry(ctx context.Context, resetTokenID string) (string, error) {

	var expiresAt string

	err := r.DB.QueryRowContext(ctx, constants.GetResetTokenExpiryTime, resetTokenID).Scan(&expiresAt)

	return expiresAt, err

}

func (r *ResetTokenRepository) GetResetTokenDetails(ctx context.Context, resetTokenID string) (string, string, error) {
	var userID, expiresAt string

	err := r.DB.QueryRowContext(ctx, constants.GetResetTokenDetails, resetTokenID).Scan(&userID, &expiresAt)

	return userID, expiresAt, err
}

func (r *ResetTokenRepository) DeleteResetToken(ctx context.Context, resetTokenID string) error {

	_, err := r.DB.ExecContext(ctx, constants.DeleteResetToken, resetTokenID)

	return err

}
