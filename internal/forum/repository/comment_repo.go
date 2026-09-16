package repository

import (
	"context"
	"database/sql"
	"forum/internal/forum/models"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(ctx context.Context, comment *models.Comment) error {
	query := `INSERT INTO comments (id, post_id, user_id, content) VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, comment.ID, comment.PostID, comment.UserID, comment.Content)

	return err
}