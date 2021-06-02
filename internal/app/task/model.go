package task

import (
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

	UserID  uint `json:"user_id"`
	BoardID uint `json:"board_id"`
}

// NewTask returns a new Task struct
func NewTask() *Task {
	return &Task{}
}

// ListTasksByBoardID returns a list of tasks with given boardID
func ListTasksByBoardID(db *gorm.DB, id string) ([]Task, error) {
	var tasks []Task

	err := db.Where("board_id = ?", id).Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

// Create adds a task to the DB
func (t *Task) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Get adds a task to the DB
func (t *Task) Get(db *gorm.DB) error {
	return db.First(&t).Error
}

// Update adds a task to the DB
func (t *Task) Update(db *gorm.DB) error {
	return db.Model(&t).Update("title", "body").Error
}

// Delete adds a task to the DB
func (t *Task) Delete(db *gorm.DB) error {
	return db.Delete(&t).Error
}
