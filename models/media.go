package models

import (
	"github.com/jebog/stuble/database"
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	ModelID   uint   `json:"model_id,omitempty" json:"model_id"`
	ModelType string `json:"model_type,omitempty" json:"model_type"`
	FileName  string `json:"file_name,omitempty" json:"file_name"`
	MimeType  string `json:"mime_type,omitempty" json:"mime_type"`
}

func (u *Media) Get(*Media) []Media {
	var reviews []Media
	database.Database.Find(&reviews)

	return reviews
}

func (u *Media) Save() (*Media, error) {
	err := database.Database.Create(&u).Error
	if err != nil {
		return &Media{}, err
	}
	return u, err
}

func (u *Media) Update() (*Media, error) {

	if err := database.Database.Save(&u).Error; err != nil {
		return &Media{}, err
	}

	return u, nil
}

func (u *Media) Delete(id uint) error {

	if err := database.Database.Delete(&u, "user_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
