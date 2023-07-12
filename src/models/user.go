package models

import (
	"github.com/jebog/stuble/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	Password        string `gorm:"size:255;not null;omitempty" json:"-"`
	Email           string `gorm:"size:255;not null;unique;"`
	EmailVerifiedAt time.Time
	Reservations    *[]Reservation
	Rooms           *[]Room
	Details         *UserDetails
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := database.Database.Where("email=?", email).
		Preload("Details").
		Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Rooms").
		Preload("Details").
		Preload("Reservations").
		Where("ID=?", id).Find(&user).Error

	if err != nil {
		return User{}, err
	}
	return user, nil
}
