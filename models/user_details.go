package models

import (
	"gorm.io/gorm"
	"time"
)

type UserDetails struct {
	gorm.Model
	UserID          uint
	FirstName       string `gorm:"size:255;not null" json:"first_name"`
	LastName        string `gorm:"size:255;not null" json:"last_name"`
	PhoneNumber     string `gorm:"size:255;not null" json:"phone_number"`
	Description     string `gorm:"type:text;size:500;not null" json:"description"`
	ProfileImage    string `gorm:"size:255;not null" json:"profile_image"`
	Email           string `gorm:"size:255;not null;unique" json:"email"`
	EmailVerifiedAt time.Time
	RememberToken   string `gorm:"size:255"`
}
