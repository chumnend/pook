package server

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Booking represents the scheduled time for a user to reserve a location
type Booking struct {
	gorm.Model

	Arrival   time.Time `json:"arrivalDate"`
	Departure time.Time `json:"departureDate"`
	TotalCost float64   `json:"totalCost"`

	UserID     uint
	LocationID uint
}
