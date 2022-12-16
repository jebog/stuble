package models

import (
	"github.com/jebog/stuble/database"
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	HomeType       string    `json:"home_type"`
	RoomType       string    `json:"room_type"`
	TotalOccupancy uint8     `json:"total_occupancy"`
	TotalBedrooms  uint8     `json:"total_bedrooms"`
	Summary        string    `gorm:"size:255" json:"summary"`
	Address        string    `gorm:"size:255" json:"address"`
	HasTV          bool      `json:"has_tv"`
	HasKitchen     bool      `json:"has_kitchen"`
	HasAirCon      bool      `json:"has_air_con"`
	HasHeating     bool      `json:"has_heating"`
	HasInternet    bool      `json:"has_internet"`
	Price          float32   `json:"price"`
	PublishedAt    time.Time `json:"published_at"`
	Latitude       float32   `json:"latitude"`
	Longitude      float32   `json:"longitude"`
	UserID         uint
	User           User
}

func NewRoom() *Room {
	return &Room{}
}

func (model *Room) Save() (*Room, error) {
	err := database.Database.Create(&model).Error
	if err != nil {
		return &Room{}, err
	}
	return model, nil
}

func (model2 *Room) Update(model *gorm.Model) {
	//TODO implement me
	panic("implement me")
}

func (model2 *Room) Delete(model *gorm.Model) {
	//TODO implement me
	panic("implement me")
}
