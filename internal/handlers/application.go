package handlers

import (
	"database/sql"
)

type Application struct {
	DB *sql.DB
}
