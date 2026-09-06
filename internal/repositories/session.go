package repositories

import "database/sql"

type SessionRepository struct {
	DB *sql.DB
}
