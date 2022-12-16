package models

import (
	"github.com/jebog/stuble/database"
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	User      User
	UserID    uint `json:"user_id,omitempty"`
	Room      Room
	RoomID    uint      `json:"room_id,omitempty"`
	StartDate time.Time `json:"start_date" json:"start_date"`
	EndDate   time.Time `json:"end_date" json:"end_date"`
	Price     float32   `json:"price" json:"price,omitempty"`
	Total     float32   `json:"total" json:"total,omitempty"`
}

func NewReservation(model gorm.Model) *Reservation {
	return &Reservation{Model: model}
}

func (r *Reservation) Save() (*Reservation, error) {
	err := database.Database.Create(&r).Error
	if err != nil {
		return &Reservation{}, err
	}
	return r, nil
}

func (r *Reservation) Update() (*Reservation, error) {
	err := database.Database.Updates(&r).Error
	if err != nil {
		return &Reservation{}, err
	}
	return r, nil
}

func (r *Reservation) Delete() (*Reservation, error) {
	err := database.Database.Delete(&r).Error
	if err != nil {
		return &Reservation{}, err
	}
	return r, nil
}
