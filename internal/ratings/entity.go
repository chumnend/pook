package ratings

import (
	"time"

	"github.com/google/uuid"
)

// Rating represents a rating in the Ratings table
type Rating struct {
	Id        uuid.UUID `json:"id"`
	BookId    uuid.UUID `json:"bookId"`
	UserId    uuid.UUID `json:"userId"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
}
