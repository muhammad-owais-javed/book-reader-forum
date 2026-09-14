package repository

import (
	"context"
	"database/sql"
	"forum/internal/forum/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *models.Post) error {

	query := `INSERT INTO posts (id, user_id, title, content) VALUES (?, ?, ?, ?)`

	_, err := r.DB.ExecContext(ctx, query, post.ID, post.UserID, post.Title, post.Content)
	
	return err
}

func (r *PostRepository) GetAllPosts(ctx context.Context) ([]*models.Post, error) {

	query := `SELECT id, user_id, title, content, created_at FROM posts ORDER BY created_at DESC`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*models.Post

	for rows.Next() {
		
		p := &models.Post{}
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt)
		
		if err != nil {
			return nil, err
		}
		
		posts = append(posts, p)

	}

	return posts, nil
}
