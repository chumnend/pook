package comments

import (
	"time"

	"github.com/google/uuid"
)

// Comment represents a comment in the Comments table
type Comment struct {
	Id        uuid.UUID `db:"id" json:"id"`
	BookId    uuid.UUID `db:"book_id" json:"bookId"`
	UserId    uuid.UUID `db:"user_id" json:"userId"`
	Comment   string    `db:"comment" json:"comment"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
