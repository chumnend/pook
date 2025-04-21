package models

import (
	"time"

	"github.com/chumnend/pook/internal/db"
	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"userId" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func CreateBook(userID uuid.UUID, title string) error {
	id := uuid.New()
	createdAt := time.Now()

	_, err := db.DB.Exec(
		"INSERT INTO books (id, user_id, title, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		id, userID, title, createdAt, createdAt,
	)
	return err
}

func GetAllBooks() (*[]Book, error) {
	rows, err := db.DB.Query("SELECT id, user_id, title, created_at, updated_at FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.UserID, &book.Title, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return &books, nil
}

func GetBooksByUserID(userID uuid.UUID) (*[]Book, error) {
	rows, err := db.DB.Query(
		"SELECT id, user_id, title, created_at, updated_at FROM books WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.UserID, &book.Title, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return &books, nil
}

func GetBookByID(id uuid.UUID) (*Book, error) {
	var book Book
	err := db.DB.QueryRow(
		"SELECT id, user_id, title, created_at, updated_at FROM books WHERE id = $1",
		id,
	).Scan(&book.ID, &book.UserID, &book.Title, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func UpdateBookByID(id uuid.UUID, title string) error {
	updatedAt := time.Now()

	_, err := db.DB.Exec(
		"UPDATE books SET title = $1, updated_at = $2 WHERE id = $3",
		title, updatedAt, id,
	)
	return err
}

func DeleteBookByID(id uuid.UUID) error {
	_, err := db.DB.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
