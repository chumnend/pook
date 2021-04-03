package book

import (
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

	UserID uint        `json:"user_id"`
	Tasks  []task.Task `json:"tasks"`
}

// NewBook returns a new Book struct
func NewBook() *Book {
	return &Book{}
}

// ListBooksByUserID returns a list of books with given UserID
func ListBooksByUserID(db *gorm.DB, id string) ([]Book, error) {
	var books []Book

	err := db.Where("user_id = ?", id).Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

// Create adds a Book to the DB
func (b *Book) Create(db *gorm.DB) error {
	return db.Create(&b).Error
}

// Get adds a Book to the DB
func (b *Book) Get(db *gorm.DB) error {
	return db.First(&b).Error
}

// Update adds a Book to the DB
func (b *Book) Update(db *gorm.DB) error {
	return db.Model(&b).Update("title", "body").Error
}

// Delete adds a Book to the DB
func (b *Book) Delete(db *gorm.DB) error {
	return db.Delete(&b).Error
}
