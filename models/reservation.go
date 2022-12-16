package models

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	User      User      `gorm:"embedded"`
	Room      Room      `gorm:"embedded"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Price     float32   `json:"price"`
	Total     float32   `json:"total"`
}

func NewReservation(model gorm.Model) *Reservation {
	return &Reservation{Model: model}
}
