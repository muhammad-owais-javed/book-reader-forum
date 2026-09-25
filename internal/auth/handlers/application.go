package handlers

import (
	"database/sql"
	"forum/internal/auth/services"

	forumHandlers "forum/internal/forum/handlers"
)

type Application struct {
	DB *sql.DB

	UserService       *services.UserService
	SessionService    *services.SessionService
	ResetTokenService *services.ResetTokenService
	Forum             *forumHandlers.ForumHandler
}
