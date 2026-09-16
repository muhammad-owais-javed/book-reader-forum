package models

import (
	"time"
)

type Post struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
}

type Comment struct {
	ID        string
	PostID    string
	UserID    string
	Content   string
	CreatedAt time.Time
}