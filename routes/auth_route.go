package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/controllers"
)

type AuthRoute struct {
}

func NewAuthRoute(route *gin.Engine) {

	route.POST("/api/register", controllers.AuthController{}.Register)
	route.POST("/api/login", controllers.AuthController{}.Login)
}
