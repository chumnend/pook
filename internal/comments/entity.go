package comments

import (
	"time"

	"github.com/google/uuid"
)

// Comment represents a comment in the Comments table
type Comment struct {
	Id        uuid.UUID `json:"id"`
	BookId    uuid.UUID `json:"bookId"`
	UserId    uuid.UUID `json:"userId"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}
