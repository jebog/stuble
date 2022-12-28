package models

import (
	"github.com/jebog/stuble/database"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Reservation   Reservation
	ReservationID uint
	Rating        uint8
	Comment       string
}

func (u *Review) Get(*Review) []Review {
	var reviews []Review
	database.Database.Find(&reviews)

	return reviews
}

func (u *Review) Save() (*Review, error) {
	err := database.Database.Create(&u).Error
	if err != nil {
		return &Review{}, err
	}
	return u, err
}

func (u *Review) Update() (*Review, error) {

	if err := database.Database.Save(&u).Error; err != nil {
		return &Review{}, err
	}

	return u, nil
}

func (u *Review) Delete(id uint) error {

	if err := database.Database.Delete(&u, "user_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
