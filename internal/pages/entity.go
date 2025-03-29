package pages

import (
	"time"

	"github.com/google/uuid"
)

// Page represents a page in the Pages table
type Page struct {
	Id        uuid.UUID `db:"id" json:"id"`
	BookId    uuid.UUID `db:"book_id" json:"bookId"`
	ImageURL  string    `db:"image_url" json:"imageUrl"`
	Caption   string    `db:"caption" json:"caption"`
	PageOrder int       `db:"page_order" json:"pageOrder"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"created" json:"updatedAt"`
}
