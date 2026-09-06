package services

import "forum/internal/repositories"

type SessionService struct {
	Repository *repositories.SessionRepository
}

// creates a session with a UUID, userID and expiration
