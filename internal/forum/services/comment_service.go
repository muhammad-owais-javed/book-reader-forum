package services

import (
	"context"
	"errors"
	"strings"
	"forum/internal/uuid"
	"forum/internal/forum/models"
	"forum/internal/forum/repository"
)

type CommentService struct {
	repo *repository.CommentRepository
}

func NewCommentService(repo *repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(ctx context.Context, userID string, postID string, content string) error {

	content = strings.TrimSpace(content)

	if content == "" {
		return errors.New("comment cannot be empty")
	}

	commentID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	comment := &models.Comment{
		ID:      commentID,
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}

	return s.repo.CreateComment(ctx, comment)
}

func (s *CommentService) GetCommentsByPostID(ctx context.Context, postID string) ([]models.Comment, error) {
	
	return s.repo.GetCommentsByPostID(ctx, postID)
}