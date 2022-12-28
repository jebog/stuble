package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/helpers"
	"github.com/jebog/stuble/models"
	"net/http"
)

type UserDetailsController struct {
	apiVersion string
}

// Get             godoc
// @Summary      Get details array
// @Description  Responds with the list of all details as JSON.
// @Tags         details
// @Produce      json
// @Success      200  {array}  models.UserDetails
// @Router       /users/details [get]
func (userController UserDetailsController) Get(context *gin.Context) {
	//TODO Add Filters
	//user, err := helpers.CurrentUser(context)

	var details []models.UserDetails

	if err := database.Database.Find(&details).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": details})
}

// Create             godoc
// @Summary      Create detail object
// @Description  Responds with the created detail as JSON.
// @Tags         details
// @Produce      json
// @Param        detail  body      models.UserDetails  true  "UserDetails JSON"
// @Success      200  {object}  models.UserDetails
// @Router       /users/details/create [put]
func (userController UserDetailsController) Create(context *gin.Context) {
	var input models.UserDetails

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
// @Summary      Update detail object
// @Description  Responds with the Updated detail as JSON.
// @Tags         details
// @Produce      json
// @Param        detail  body      models.UserDetails  true  "UserDetails JSON"
// @Success      200  {object}  models.UserDetails
// @Router       /user/details/update/:id [patch]
func (userController UserDetailsController) Update(context *gin.Context) {

	var userDetails models.UserDetails

	if err := context.ShouldBindJSON(&userDetails); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDetails.UserID = user.ID

	if _, e := userDetails.Update(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"data": &userDetails})
}

// Destroy             godoc
// @Summary      Destroy detail array
// @Description  Responds with the created detail as JSON.
// @Tags         details
// @Produce      json
// @Param		 ID	body		int		true	"UserDetails ID"
// @Success      200  {object}  models.UserDetails
// @Router       /users/details/delete [delete]
func (userController UserDetailsController) Destroy(context *gin.Context) {
	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userDetails models.UserDetails
	userDetails.UserID = user.ID

	if e := userDetails.Delete(user.ID); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
