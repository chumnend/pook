package users

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the User table
type User struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
}
