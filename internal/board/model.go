package board

import (
	"time"

	"github.com/chumnend/pook/internal/task"
	"github.com/jinzhu/gorm"
)

// Board struct declaration
type Board struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	UserID uint        `json:"user_id"`
	Tasks  []task.Task `json:"tasks"`
}

// NewBoard returns a new Board struct
func NewBoard() *Board {
	return &Board{}
}

// ListBoardsByUserID returns a list of boards with given UserID
func ListBoardsByUserID(db *gorm.DB, id string) ([]Board, error) {
	var boards []Board

	err := db.Where("user_id = ?", id).Find(&boards).Error
	if err != nil {
		return boards, err
	}

	return boards, nil
}

// Create adds a Board to the DB
func (b *Board) Create(db *gorm.DB) error {
	return db.Create(&b).Error
}

// Get adds a Board to the DB
func (b *Board) Get(db *gorm.DB) error {
	return db.First(&b).Error
}

// Update adds a Board to the DB
func (b *Board) Update(db *gorm.DB) error {
	return db.Model(&b).Update("title", "body").Error
}

// Delete adds a Board to the DB
func (b *Board) Delete(db *gorm.DB) error {
	return db.Delete(&b).Error
}
