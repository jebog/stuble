package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Reservation   Reservation
	ReservationID uint
	Rating        uint8
	Comment       string
}

func NewReview(reservation uint, rating uint8, comment string) *Review {
	return &Review{ReservationID: reservation, Rating: rating, Comment: comment}
}

func (r Review) Save(model *gorm.Model) {
	//TODO implement me
	panic("implement me")
}

func (r Review) Update(model *gorm.Model) {
	//TODO implement me
	panic("implement me")
}

func (r Review) Delete(model *gorm.Model) {
	//TODO implement me
	panic("implement me")
}
