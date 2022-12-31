package requests

type ReservationInput struct {
	UserId    string `json:"user_id" binding:"required"`
	RoomId    string `json:"room_id" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Price     string `json:"price" binding:"required"`
	Total     string `json:"total" binding:"required"`
}
