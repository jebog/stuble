package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/helpers"
	"github.com/jebog/stuble/models"
	"net/http"
)

type MediaController struct {
}

// Get             godoc
// @Summary      Get reviews array
// @Description  Responds with the list of all reviews as JSON.
// @Tags         reviews
// @Produce      json
// @Success      200  {array}  models.Review
// @Router       /reviews [get]
func (controller MediaController) Get(context *gin.Context) {
	_, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var medias []models.Media

	if err := database.Database.Debug().Find(&medias).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": &medias})
}

// Create             godoc
// @Summary      Create review object
// @Description  Responds with the created review as JSON.
// @Tags         reviews
// @Produce      json
// @Param        review  body      models.Review  true  "Review JSON"
// @Success      200  {object}  models.Review
// @Router       /reviews/create [put]
func (controller MediaController) Create(context *gin.Context) {
	var input models.Media
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

// Update        godoc
// @Summary      Update review object
// @Description  Responds with the Updated review as JSON.
// @Tags         reviews
// @Produce      json
// @Param        review  body      models.Review  true  "Review JSON"
// @Success      200  {object}  models.Review
// @Router       /reviews/update/:id [patch]
func (controller MediaController) Update(context *gin.Context) {
	var media models.Media

	if err := context.ShouldBindJSON(&media); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, e := media.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"data": &media})
}

// Destroy             godoc
// @Summary      Destroy review array
// @Description  Responds with the created review as JSON.
// @Tags         reviews
// @Produce      json
// @Param		 ID	body		int		true	"Review ID"
// @Success      200  {object}  models.Review
// @Router       /reviews/delete [delete]
func (controller MediaController) Destroy(context *gin.Context) {
	var media models.Media

	if err := context.ShouldBindJSON(&media); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if e := media.Delete(user.ID); e != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
