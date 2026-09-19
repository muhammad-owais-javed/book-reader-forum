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

func (r *CommentRepository) GetCommentsByPostID(ctx context.Context, postID string) ([]models.Comment, error) {
	query := `SELECT c.id, c.post_id, c.user_id, u.username, c.content, c.created_at FROM comments c JOIN users u ON c.user_id = u.id WHERE c.post_id = ? ORDER BY c.created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Username, &comment.Content, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}