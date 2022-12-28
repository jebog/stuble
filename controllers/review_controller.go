package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/helpers"
	"github.com/jebog/stuble/models"
	"net/http"
)

type ReviewController struct {
}

// Get             godoc
// @Summary      Get reviews array
// @Description  Responds with the list of all reviews as JSON.
// @Tags         reviews
// @Produce      json
// @Success      200  {array}  models.Review
// @Router       /reviews [get]
func (controller ReviewController) Get(context *gin.Context) {
	_, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reviews []models.Review

	if err := database.Database.Debug().Find(&reviews).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": &reviews})
}

// Create             godoc
// @Summary      Create review object
// @Description  Responds with the created review as JSON.
// @Tags         reviews
// @Produce      json
// @Param        review  body      models.Review  true  "Review JSON"
// @Success      200  {object}  models.Review
// @Router       /reviews/create [put]
func (controller ReviewController) Create(context *gin.Context) {
	var input models.Review
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

// Update             godoc
// @Summary      Update review object
// @Description  Responds with the Updated review as JSON.
// @Tags         reviews
// @Produce      json
// @Param        review  body      models.Review  true  "Review JSON"
// @Success      200  {object}  models.Review
// @Router       /reviews/update/:id [patch]
func (controller ReviewController) Update(context *gin.Context) {
	var review models.Review

	if err := context.ShouldBindJSON(&review); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, e := review.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"data": &review})
}

// Destroy             godoc
// @Summary      Destroy review array
// @Description  Responds with the created review as JSON.
// @Tags         reviews
// @Produce      json
// @Param		 ID	body		int		true	"Review ID"
// @Success      200  {object}  models.Review
// @Router       /reviews/delete [delete]
func (controller ReviewController) Destroy(context *gin.Context) {
	var review models.Review

	if err := context.ShouldBindJSON(&review); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if e := review.Delete(user.ID); e != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
