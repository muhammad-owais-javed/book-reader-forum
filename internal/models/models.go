package models

import (
    "time"
)

type User struct {
    ID           string
    Username     string
    Email        string
    PasswordHash string
}

type Session struct {
    ID        string
    UserID    string
    ExpiresAt time.Time
}