package models

import (
	"time"
)

type Post struct {
	ID        string
	UserID    string
	Title     string
	AuthorName string
	Content   string
	CreatedAt time.Time
}
