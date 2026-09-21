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
	Comments  []Comment
}

type Comment struct {
	ID        string
	PostID    string
	UserID    string
	Username  string
	Content   string
	CreatedAt time.Time
}

type PostReaction struct {
	UserID string
	PostID string
	IsLike bool
}