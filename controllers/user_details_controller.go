package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/helpers"
	"github.com/jebog/stuble/models"
	"github.com/jebog/stuble/requests"
	"net/http"
)

type UserDetailsController struct {
	apiVersion string
}

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

func (userController UserDetailsController) Update(context *gin.Context) {

	var input requests.UserDetailsInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userDetails models.UserDetails
	userDetails.UserID = user.ID

	if _, e := userDetails.Update(context.Param("id"), &input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"data": &userDetails})
}

func (userController UserDetailsController) Destroy(context *gin.Context) {
	user, err := helpers.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userDetails models.UserDetails
	userDetails.UserID = user.ID

	if e := userDetails.Delete(context.Param("id")); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
