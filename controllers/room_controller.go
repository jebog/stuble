package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/helpers"
	"github.com/jebog/stuble/models"
	"net/http"
)

type RoomController struct {
}

// Get             godoc
// @Summary      Get rooms array
// @Description  Responds with the list of all rooms as JSON.
// @Tags         rooms
// @Produce      json
// @Success      200  {array}  models.Room
// @Router       /rooms [get]
func (controller RoomController) Get(context *gin.Context) {
	_, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var rooms []models.Room

	if err := database.Database.Find(&rooms).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": &rooms})
}

// Create             godoc
// @Summary      Create room object
// @Description  Responds with the created room as JSON.
// @Tags         rooms
// @Produce      json
// @Param        room  body      models.Room  true  "Room JSON"
// @Success      200  {object}  models.Room
// @Router       /rooms/create [put]
func (controller RoomController) Create(context *gin.Context) {
	var input models.Room
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
// @Summary      Update room object
// @Description  Responds with the Updated room as JSON.
// @Tags         rooms
// @Produce      json
// @Param        room  body      models.Room  true  "Room JSON"
// @Success      200  {object}  models.Room
// @Router       /rooms/update/:id [patch]
func (controller RoomController) Update(context *gin.Context) {
	var room models.Room

	if err := context.ShouldBindJSON(&room); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room.UserID = user.ID

	if _, e := room.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"data": &room})
}

// Destroy             godoc
// @Summary      Destroy room array
// @Description  Responds with the created room as JSON.
// @Tags         rooms
// @Produce      json
// @Param		 ID	body		int		true	"Room ID"
// @Success      200  {object}  models.Room
// @Router       /rooms/delete [delete]
func (controller RoomController) Destroy(context *gin.Context) {
	var room models.Room

	if err := context.ShouldBindJSON(&room); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if e := room.Delete(user.ID); e != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
