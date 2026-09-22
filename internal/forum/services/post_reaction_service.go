package services

import (
	"context"
	"forum/internal/forum/models"
	"forum/internal/forum/repository"
)

type PostReactionService struct {
	repo *repository.PostReactionRepository
}

func NewPostReactionService(repo *repository.PostReactionRepository) *PostReactionService {
	return &PostReactionService{repo: repo}
}

func (s *PostReactionService) ToggleReaction(ctx context.Context, userID string, postID string, isLike bool) error {
	existingReaction, err := s.repo.GetReaction(ctx, userID, postID)
	if err != nil {
		return err
	}
	// no reaction yet
	if existingReaction == nil {
		reaction := &models.PostReaction{UserID: userID, PostID: postID, IsLike: isLike}
		return s.repo.CreateReaction(ctx, reaction)
	}
	// user pressed on reaction again
	if existingReaction.IsLike == isLike {
		return s.repo.DeleteReaction(ctx, userID, postID)
	}
	// user switched from like to dislike or other way around
	return s.repo.UpdateReaction(ctx, userID, postID, isLike)
}

func (s *PostReactionService) GetReactionCounts(ctx context.Context, postID string) (int, int, error) {
	return s.repo.GetReactionCounts(ctx, postID)
}

func (s *PostReactionService) GetReaction(ctx context.Context,userID string,postID string) (*models.PostReaction, error) {
	return s.repo.GetReaction(ctx, userID, postID)
}