package book

import (
	"errors"
	"time"

	"github.com/chumnend/pook/internal/task"
	"github.com/jinzhu/gorm"
)

// Book struct declaration
type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	UserID uint
	Tasks  []task.Task
}

// NewBook returns a new Book struct
func NewBook() *Book {
	return &Book{}
}

// Create adds a Book to the DB
func (t *Book) Create(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Get adds a Book to the DB
func (t *Book) Get(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Update adds a Book to the DB
func (t *Book) Update(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Delete adds a Book to the DB
func (t *Book) Delete(db *gorm.DB) error {
	return errors.New("Not implemented")
}
