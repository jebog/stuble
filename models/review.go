package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Reservation Reservation
	Rating      int8
	Comment     string
}

func NewReview(reservation Reservation, rating int8, comment string) *Review {
	return &Review{Reservation: reservation, Rating: rating, Comment: comment}
}
