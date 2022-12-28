package models

import (
	"github.com/jebog/stuble/database"
	"gorm.io/gorm"
)

type UserDetails struct {
	gorm.Model
	UserID        uint   `gorm:"omitempty;unique;not null" json:"user_id"`
	FirstName     string `gorm:"size:255;not null;omitempty" json:"first_name"`
	LastName      string `gorm:"size:255;not null;omitempty" json:"last_name"`
	PhoneNumber   string `gorm:"size:255;not null;omitempty" json:"phone_number"`
	Description   string `gorm:"type:text;size:500;not null;omitempty" json:"description"`
	ProfileImage  string `gorm:"size:255;omitempty" json:"profile_image"`
	RememberToken string `gorm:"size:255;omitempty" json:"remember_token"`
}

func (u *UserDetails) Get(*UserDetails) []UserDetails {
	var details []UserDetails
	database.Database.Find(&details)

	return details
}

func (u *UserDetails) Save() (*UserDetails, error) {
	err := database.Database.Create(&u).Error
	if err != nil {
		return &UserDetails{}, err
	}
	return u, err
}

func (u *UserDetails) Update() (*UserDetails, error) {

	if err := database.Database.Save(&u).Error; err != nil {
		return &UserDetails{}, err
	}

	return u, nil
}

func (u *UserDetails) Delete(id uint) error {

	if err := database.Database.Delete(&u, "user_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
