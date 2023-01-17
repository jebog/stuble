package models

import (
	"github.com/jebog/stuble/database"
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	RomeName	   string    `json:"room_name,omitempty"`
	HomeType       string    `json:"home_type,omitempty"`
	RoomType       string    `json:"room_type,omitempty"`
	TotalOccupancy uint8     `gorm:"default:1" json:"total_occupancy,omitempty"`
	TotalBedrooms  uint8     `gorm:"default:1" json:"total_bedrooms,omitempty"`
	Summary        string    `gorm:"size:255" json:"summary,omitempty"`
	Address        string    `gorm:"size:255" json:"address,omitempty"`
	HasTV          *bool     `gorm:"default:false" json:"has_tv,omitempty"`
	HasKitchen     *bool     `gorm:"default:false" json:"has_kitchen,omitempty"`
	HasAirCon      *bool     `gorm:"default:false" json:"has_air_con,omitempty"`
	HasHeating     *bool     `gorm:"default:false" json:"has_heating,omitempty"`
	HasInternet    *bool     `gorm:"default:false" json:"has_internet,omitempty"`
	Price          float32   `json:"price,omitempty"`
	PublishedAt    time.Time `json:"published_at,omitempty"`
	Latitude       float32   `gorm:"default:0" json:"latitude,omitempty"`
	Longitude      float32   `gorm:"default:0" json:"longitude,omitempty"`
	UserID         uint
	User           User
	Reservation    []Reservation
}

// Eager loading of rooms when find users
func (u *Room) Preload(*Room) []Room {
	var users []User
	var rooms []Room
	database.Database.Preload(&u).Find(&users)

	return rooms
}


func (u *Room) Get(*Room) []Room {
	var rooms []Room
	database.Database.Find(&rooms)

	return rooms
}

func (u *Room) Save() (*Room, error) {
	err := database.Database.Create(&u).Error
	if err != nil {
		return &Room{}, err
	}
	return u, err
}

func (u *Room) Update() (*Room, error) {

	if err := database.Database.Save(&u).Error; err != nil {
		return &Room{}, err
	}

	return u, nil
}

func (u *Room) Delete(id uint) error {

	if err := database.Database.Delete(&u, "user_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
