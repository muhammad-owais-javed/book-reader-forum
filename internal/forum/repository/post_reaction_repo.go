package repository

import (
	"context"
	"database/sql"
	"forum/internal/forum/models"
)

type PostReactionRepository struct {
	db *sql.DB
}

func NewPostReactionRepository(db *sql.DB) *PostReactionRepository {
	return &PostReactionRepository{db: db}
}

func (r *PostReactionRepository) GetReaction(ctx context.Context, userID string, postID string) (*models.PostReaction, error) {
	query := `SELECT user_id, post_id, is_like FROM post_reactions WHERE user_id = ? AND post_id = ?`
	var reaction models.PostReaction
	err := r.db.QueryRowContext(ctx, query, userID, postID,).Scan(&reaction.UserID, &reaction.PostID, &reaction.IsLike)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &reaction, nil
}

func (r *PostReactionRepository) CreateReaction(ctx context.Context, reaction *models.PostReaction) error {
	query := `INSERT INTO post_reactions (user_id, post_id, is_like) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, reaction.UserID, reaction.PostID, reaction.IsLike)
	return err
}

func (r *PostReactionRepository) UpdateReaction(ctx context.Context, userID string, postID string, isLike bool) error {
	query := `UPDATE post_reactions SET is_like = ? WHERE user_id = ? AND post_id = ?`
	_, err := r.db.ExecContext(ctx, query, isLike, userID, postID)
	return err
}

func (r *PostReactionRepository) DeleteReaction(ctx context.Context, userID string, postID string) error {
	query := `DELETE FROM post_reactions WHERE user_id = ? AND post_id = ?`
	_, err := r.db.ExecContext(ctx, query, userID, postID)
	return err
}

func (r *PostReactionRepository) GetReactionCounts(ctx context.Context, postID string) (int, int, error) {
	query := `SELECT COUNT(CASE WHEN is_like = 1 THEN 1 END), COUNT(CASE WHEN is_like = 0 THEN 1 END) FROM post_reactions WHERE post_id = ?`
	var likeCount int
	var dislikeCount int
	err := r.db.QueryRowContext(ctx, query, postID).Scan(&likeCount, &dislikeCount)
	if err != nil {
		return 0, 0, err
	}
	return likeCount, dislikeCount, nil
}