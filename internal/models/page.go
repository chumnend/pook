package models

import (
	"time"

	"github.com/chumnend/pook/internal/db"
	"github.com/google/uuid"
)

type Page struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BookID    uuid.UUID `json:"bookId" db:"book_id"`
	ImageURL  string    `json:"imageUrl" db:"image_url"`
	Caption   string    `json:"caption" db:"caption"`
	PageOrder int       `json:"pageOrder" db:"page_order"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func CreatePage(bookId uuid.UUID, imageUrl string, caption string, pageOrder int) error {
	id := uuid.New()
	createdAt := time.Now()

	_, err := db.DB.Exec(
		"INSERT INTO pages (id, book_id, image_url, caption, page_order, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		id, bookId, imageUrl, caption, pageOrder, createdAt, createdAt,
	)

	return err
}

func GetPagesByBookID(bookId uuid.UUID) (*[]Page, error) {
	rows, err := db.DB.Query(
		"SELECT id, book_id, image_url, caption, page_order, created_at, updated_at FROM pages WHERE book_id = $1 ORDER BY page_order",
		bookId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []Page
	for rows.Next() {
		var page Page
		err := rows.Scan(
			&page.ID,
			&page.BookID,
			&page.ImageURL,
			&page.Caption,
			&page.PageOrder,
			&page.CreatedAt,
			&page.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}

	return &pages, nil
}

func GetPageByBookID(bookId uuid.UUID, pageOrder int) (*Page, error) {
	var page Page
	err := db.DB.QueryRow(
		"SELECT id, book_id, image_url, caption, page_order, created_at, updated_at FROM pages WHERE book_id = $1 AND page_order = $2",
		bookId, pageOrder,
	).Scan(
		&page.ID,
		&page.BookID,
		&page.ImageURL,
		&page.Caption,
		&page.PageOrder,
		&page.CreatedAt,
		&page.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func UpdatePage(id uuid.UUID, imageUrl string, caption string, pageOrder int) error {
	updatedAt := time.Now()
	_, err := db.DB.Exec(
		"UPDATE pages SET image_url = $1, caption = $2, page_order = $3, updated_at = $4 WHERE id = $5",
		imageUrl, caption, pageOrder, updatedAt, id,
	)
	return err
}

func DeletePage(id uuid.UUID) error {
	_, err := db.DB.Exec(
		"DELETE FROM pages WHERE id = $1",
		id,
	)
	return err
}
