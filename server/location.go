package server

import "github.com/jinzhu/gorm"

// Location represents a spot that can be booked
type Location struct {
	gorm.Model

	Address    string  `json:"address"`
	CostPerDay float64 `json:"costPerDay"`

	UserID uint `gorm:"foreignKey:Owner"`
}
