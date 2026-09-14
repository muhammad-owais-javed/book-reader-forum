package handlers

import (
	"database/sql"
	"forum/internal/auth/services"
)

type Application struct {
	DB *sql.DB

	UserService *services.UserService
	SessionService *services.SessionService

}
