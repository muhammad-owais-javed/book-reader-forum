package services

import (
	"context"
	"forum/internal/forum/models"
	"forum/internal/forum/repository"
	"forum/internal/uuid"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}


func (s *PostService) CreatePost(ctx context.Context, userID, title, content string) error {

	postID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	post := &models.Post{
		ID:      postID,
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	return s.repo.CreatePost(ctx, post)
}

func (s *PostService) GetAllPosts(ctx context.Context) ([]*models.Post, error) {

	return s.repo.GetAllPosts(ctx)
	
}
