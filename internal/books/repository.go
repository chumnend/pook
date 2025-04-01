package books

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type bookRepository struct {
	db *sql.DB
}

// NewBookRepository returns a BookRepository struct using SQL
func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (repo *bookRepository) FindAll() ([]Book, error) {
	var books []Book
	rows, err := repo.db.Query("SELECT id, user_id, title, created_at, updated_at FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Id, &book.UserId, &book.Title, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (repo *bookRepository) FindAllByUserID(userId uuid.UUID) ([]Book, error) {
	var books []Book
	rows, err := repo.db.Query("SELECT id, user_id, title, created_at, updated_at FROM books WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Id, &book.UserId, &book.Title, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (repo *bookRepository) FindByID(id uuid.UUID) (*Book, error) {
	var book Book
	err := repo.db.QueryRow("SELECT id, user_id, title, created_at, updated_at FROM books WHERE id = ?", id).
		Scan(&book.Id, &book.UserId, &book.Title, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}

	return &book, nil
}

func (repo *bookRepository) Create(book *Book) error {
	_, err := repo.db.Exec("INSERT INTO books (title, user_id) VALUES (?, ?)", book.Title, book.UserId)
	return err
}

func (repo *bookRepository) Save(book *Book) error {
	_, err := repo.db.Exec("UPDATE books SET title = ?, user_id = ? WHERE id = ?", book.Title, book.UserId, book.Id)
	return err
}

func (repo *bookRepository) Delete(book *Book) error {
	_, err := repo.db.Exec("DELETE FROM books WHERE id = ?", book.Id)
	return err
}
