package models

import (
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/requests"
	"gorm.io/gorm"
)

type UserDetails struct {
	gorm.Model
	UserID        uint   `gorm:"omitempty;unique;not null"`
	FirstName     string `gorm:"size:255;not null;omitempty"`
	LastName      string `gorm:"size:255;not null;omitempty"`
	PhoneNumber   string `gorm:"size:255;not null;omitempty"`
	Description   string `gorm:"type:text;size:500;not null;omitempty"`
	ProfileImage  string `gorm:"size:255;omitempty"`
	RememberToken string `gorm:"size:255;omitempty"`
}

func (u *UserDetails) Get(*UserDetails) []UserDetails {
	var categories []UserDetails
	database.Database.Find(&categories)

	return categories
}

func (u *UserDetails) Save() (*UserDetails, error) {
	err := database.Database.Create(&u).Error
	if err != nil {
		return &UserDetails{}, err
	}
	return u, err
}

func (u *UserDetails) Update(param string, input *requests.UserDetailsInput) (*UserDetails, error) {

	err := database.Database.First(&u, "id = ? AND user_id = ?", param, u.UserID).Error

	if err != nil {
		return &UserDetails{}, err
	}

	u.Description = input.Description
	u.FirstName = input.FirstName
	u.LastName = input.LastName
	u.ProfileImage = input.ProfileImage

	database.Database.Save(&u)

	return u, nil
}

func (u *UserDetails) Delete(param string) error {

	if err := database.Database.Debug().First(&u, "id = ? AND user_id = ?", param, u.UserID).Error; err != nil {
		return err
	}

	database.Database.Debug().Delete(&u)
	return nil
}
