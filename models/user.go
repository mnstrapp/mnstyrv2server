package models

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mnstrapp/mnstrv2server/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         string    `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password_hash"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	ArchivedAt time.Time `json:"archived_at,omitempty" db:"archived_at,omitempty"`
}

func NewUser(email, password string) (*User, error) {
	id := uuid.New().String()
	return &User{
		ID:         id,
		Email:      email,
		Password:   password,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		ArchivedAt: time.Time{},
	}, nil
}

func FromJSON(data []byte) (*User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (u *User) HashPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *User) Create() error {
	db := database.Connection()
	defer db.Close(context.Background())

	hashedPassword, err := u.HashPassword()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (id, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = db.Exec(context.Background(), query, u.ID, u.Email, hashedPassword, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
