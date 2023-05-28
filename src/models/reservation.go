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

func (u *Reservation) Get(*Reservation) []Reservation {
	var r []Reservation
	database.Database.Find(&r)

	return r
}

func (u *Reservation) Save() (*Reservation, error) {
	err := database.Database.Create(&u).Error
	if err != nil {
		return &Reservation{}, err
	}
	return u, err
}

func (u *Reservation) Update() (*Reservation, error) {

	if err := database.Database.Save(&u).Error; err != nil {
		return &Reservation{}, err
	}

	return u, nil
}

func (u *Reservation) Delete(id uint) error {

	if err := database.Database.Delete(&u, "user_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
