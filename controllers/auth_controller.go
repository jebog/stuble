package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/helpers"
	"github.com/jebog/stuble/models"
	"github.com/jebog/stuble/requests"
	"net/http"
	"os"
)

type AuthController struct {
}

func (authController AuthController) Register(context *gin.Context) {
	var input requests.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func (authController AuthController) Login(context *gin.Context) {
	var input requests.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keycloak := helpers.NewConfig().Keycloak

	token, err := keycloak.Login(context,
		os.Getenv("KEYCLOAK_CLIENT_ID"),
		os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		os.Getenv("KEYCLOAK_REALM"),
		input.Username,
		input.Password,
	)

	fmt.Print(token)

	if err != nil {
		panic("Something wrong with the credentials or url")
	}
}

/**
func (authController AuthController) Login(context *gin.Context) {
	var input requests.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

*/
