package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/helpers"
	"github.com/jebog/stuble/models"
	"net/http"
)

type UserController struct {
	apiVersion string
}

// Get             godoc
// @Summary      Get users array
// @Description  Responds with the list of all users as JSON.
// @Tags         users
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /users [get]
func (userController UserController) Get(context *gin.Context) {
	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Details})
}

// Create             godoc
// @Summary      Create user object
// @Description  Responds with the created user as JSON.
// @Tags         users
// @Produce      json
// @Param        user  body      models.User  true  "User JSON"
// @Success      200  {object}  models.User
// @Router       /users/create [put]
func (userController UserController) Create(context *gin.Context) {
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
// @Summary      Update user object
// @Description  Responds with the Updated user as JSON.
// @Tags         users
// @Produce      json
// @Param        user  body      models.User  true  "User JSON"
// @Success      200  {object}  models.User
// @Router       /users/update/:id [patch]
func (userController UserController) Update(context *gin.Context) {

}

// Destroy             godoc
// @Summary      Destroy user array
// @Description  Responds with the created user as JSON.
// @Tags         users
// @Produce      json
// @Param		 ID	body		int		true	"User ID"
// @Success      200  {object}  models.User
// @Router       /users/delete [delete]
func (userController UserController) Destroy(context *gin.Context) {

}
