package entities

import (
	"time"
)

// User representa la entidad de dominio para un usuario
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // No se serializa en JSON
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// NewUser crea una nueva instancia de User
func NewUser(username, password, email string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
	}
}