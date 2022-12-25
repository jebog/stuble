package models

import (
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	ModelID   uint   `json:"model_id,omitempty" json:"model_id"`
	ModelType string `json:"model_type,omitempty" json:"model_type"`
	FileName  string `json:"file_name,omitempty" json:"file_name"`
	MimeType  string `json:"mime_type,omitempty" json:"mime_type"`
}

func (m Media) FindById(id uint) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
