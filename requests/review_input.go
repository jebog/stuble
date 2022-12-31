package requests

type ReviewInput struct {
	ReservationID string `json:"reservation_id" binding:"required"`
	Rating        string `json:"rating" binding:"required"`
	Comment       string `json:"comment" binding:"required"`
}
