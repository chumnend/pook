package ratings

import (
	"time"

	"github.com/google/uuid"
)

// Rating represents a rating in the Ratings table
type Rating struct {
	Id        uuid.UUID `db:"id" json:"id"`
	BookId    uuid.UUID `db:"book_id" json:"bookId"`
	UserId    uuid.UUID `db:"user_id" json:"userId"`
	Rating    int       `db:"rating" json:"rating"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
