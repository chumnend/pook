package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the User table
type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
}

// Book represents a book in the Books table
type Book struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Page represents a page in the Pages table
type Page struct {
	ID        uuid.UUID `json:"id"`
	BookID    uuid.UUID `json:"bookId"`
	ImageURL  string    `json:"imageUrl"`
	Caption   string    `json:"caption"`
	PageOrder int       `json:"pageOrder"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Comment represents a comment in the Comments table
type Comment struct {
	ID        uuid.UUID `json:"id"`
	BookID    uuid.UUID `json:"book_id"`
	UserID    uuid.UUID `json:"user_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}

// Rating represents a rating in the Ratings table
type Rating struct {
	ID        uuid.UUID `json:"id"`
	BookID    uuid.UUID `json:"book_id"`
	UserID    uuid.UUID `json:"user_id"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
}
