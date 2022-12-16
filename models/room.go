package models

import (
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	HomeType       string    `json:"home_type"`
	RoomType       string    `json:"room_type"`
	TotalOccupancy uint8     `json:"total_occupancy"`
	TotalBedrooms  uint8     `json:"total_bedrooms"`
	Summary        string    `gorm:"type:text,size:255" json:"summary"`
	Address        string    `gorm:"type:text,size:255" json:"address"`
	HasTV          bool      `json:"has_tv"`
	HasKitchen     bool      `json:"has_kitchen"`
	HasAirCon      bool      `json:"has_air_con"`
	HasHeating     bool      `json:"has_heating"`
	HasInternet    bool      `gorm:"type:text,size:255" json:"has_internet"`
	Price          float32   `json:"price"`
	PublishedAt    time.Time `json:"published_at"`
	Latitude       float32   `json:"latitude"`
	Longitude      float32   `json:"longitude"`
	User           User      `gorm:"embedded"`
}

func NewRoom() *Room {
	return &Room{}
}
