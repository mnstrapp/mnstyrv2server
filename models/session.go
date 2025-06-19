package models

import "time"

type Session struct {
	ID         int       `json:"-" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`
	Token      string    `json:"token" db:"session_token"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	ArchivedAt time.Time `json:"-" db:"archived_at"`
	ExpiresAt  time.Time `json:"expires_at" db:"expires_at"`
}
