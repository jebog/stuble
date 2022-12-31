package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/configs"
	"github.com/jebog/stuble/controllers"
	"github.com/jebog/stuble/midldlewares"
)

type UserDetailsRoute struct {
	Path string
}

func NewUserDetailsRoute(route *gin.Engine) {

	h := &UserRoute{Path: configs.NewConfig().BasePath}

	r := route.Group(h.Path)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/users/details", controllers.UserDetailsController{}.Get)
	r.PUT("/users/details/create", controllers.UserDetailsController{}.Create)
	r.PATCH("/users/details/update/:id", controllers.UserDetailsController{}.Update)
	r.DELETE("/users/details/delete/:id", controllers.UserDetailsController{}.Destroy)
}
