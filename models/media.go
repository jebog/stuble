package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	ModelID   uint   `json:"model_id,omitempty"`
	ModelType string `json:"model_type,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
}

func (m Media) Save(model *Media) {
	//TODO implement me
	panic("implement me")
}

func (m Media) Update(model *gorm.Model) {
	//TODO implement me
	panic("implement me")
}

func (m Media) Delete(model *gorm.Model) {
	//TODO implement me
	panic("implement me")
}
