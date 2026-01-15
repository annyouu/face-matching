package entity

import "time"

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	PasswordHash  string `json:"password_hash"`
	Name string `json:"name"`
	ProfileImageURL string
	Status string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}