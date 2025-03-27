package pages

import (
	"time"

	"github.com/google/uuid"
)

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
