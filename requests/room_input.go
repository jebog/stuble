package requests

import "time"

type RoomInput struct {
	HomeType       string    `json:"home_type" binding:"required"`
	RoomType       string    `json:"room_type" binding:"required"`
	TotalOccupancy string    `json:"total_occupancy" binding:"required"`
	TotalBedrooms  string    `json:"total_bedrooms" binding:"required"`
	Summary        string    `json:"summary" binding:"required"`
	Address        string    `json:"address" binding:"required"`
	HasTv          bool      `json:"has_tv" binding:"required"`
	HasKitchen     bool      `json:"has_kitchen" binding:"required"`
	HasAirCon      bool      `json:"has_air_con" binding:"required"`
	HasHeating     bool      `json:"has_heating" binding:"required"`
	HasInternet    bool      `json:"has_internet" binding:"required"`
	Price          bool      `json:"price" binding:"required"`
	PublishedAt    time.Time `json:"published_at"`
	Latitude       float32   `json:"latitude" binding:"required"`
	Longitude      float32   `json:"longitude" binding:"required"`
	UserID         uint      `json:"user_id" binding:"required"`
}
