package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Reservation Reservation
	Rating      int8
	Comment     string
}
