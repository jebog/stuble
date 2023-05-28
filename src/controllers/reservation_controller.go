package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/helpers"
	"github.com/jebog/stuble/models"
	"net/http"
)

type ReservationController struct {
}

func (controller ReservationController) Get(context *gin.Context) {
	_, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reservations []models.Reservation

	if err := database.Database.Debug().Find(&reservations).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": &reservations})
}

func (controller ReservationController) Create(context *gin.Context) {
	var input models.Reservation
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func (controller ReservationController) Update(context *gin.Context) {
	var reservation models.Reservation

	if err := context.ShouldBindJSON(&reservation); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation.UserID = user.ID

	if _, e := reservation.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"data": &reservation})
}

func (controller ReservationController) Destroy(context *gin.Context) {
	var reservation models.Reservation

	if err := context.ShouldBindJSON(&reservation); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if e := reservation.Delete(user.ID); e != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
