package handlers

import (
	"database/sql"
	"forum/internal/services"
)

type Application struct {
	DB *sql.DB

	UserService *services.UserService
}
