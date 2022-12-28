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

// Get             godoc
// @Summary      Get reservations list
// @Description  Responds with the list of all or filtered reservations as JSON.
// @Tags         reservations
// @Produce      json
// @Success      200  {array}  models.Reservation
// @Router       /reservations [get]
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

// Create             godoc
// @Summary      Create reservation object
// @Description  Responds with the created reservation as JSON.
// @Tags         reservations
// @Produce      json
// @Param        reservation  body      models.Reservation  true  "Reservation JSON"
// @Success      200  {object}  models.Reservation
// @Router       /reservations/create [put]
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

// Update             godoc
// @Summary      Update Reservation object
// @Description  Responds with the Updated Reservation as JSON.
// @Tags         reservations
// @Produce      json
// @Param        reservation  body      models.Reservation  true  "Reservation JSON"
// @Success      200  {object}  models.Reservation
// @Router       /reservation/update/:id [patch]
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

// Destroy             godoc
// @Summary      Destroy reservation array
// @Description  Responds with the created reservation as JSON.
// @Tags         reservations
// @Produce      json
// @Param		 ID	body		int		true	"Reservation ID"
// @Success      200  {object}  models.Reservation
// @Router       /reservations/delete [delete]
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
