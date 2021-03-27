package task

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Task struct declaration
type Task struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	UserID uint
	BookID uint
}

// NewTask returns a new Task struct
func NewTask() *Task {
	return &Task{}
}

// Create adds a task to the DB
func (t *Task) Create(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Get adds a task to the DB
func (t *Task) Get(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Update adds a task to the DB
func (t *Task) Update(db *gorm.DB) error {
	return errors.New("Not implemented")
}

// Delete adds a task to the DB
func (t *Task) Delete(db *gorm.DB) error {
	return errors.New("Not implemented")
}
