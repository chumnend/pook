package models

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

// Book represents a book in the Books table
type Book struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Page represents a page in the Pages table
type Page struct {
	Id        uuid.UUID `json:"id"`
	BookId    uuid.UUID `json:"bookId"`
	ImageURL  string    `json:"imageUrl"`
	Caption   string    `json:"caption"`
	PageOrder int       `json:"pageOrder"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Comment represents a comment in the Comments table
type Comment struct {
	Id        uuid.UUID `json:"id"`
	BookId    uuid.UUID `json:"bookId"`
	UserId    uuid.UUID `json:"userId"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}

// Rating represents a rating in the Ratings table
type Rating struct {
	Id        uuid.UUID `json:"id"`
	BookId    uuid.UUID `json:"bookId"`
	UserId    uuid.UUID `json:"userId"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
}
