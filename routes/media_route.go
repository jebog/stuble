package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/configs"
	"github.com/jebog/stuble/controllers"
	middlewares "github.com/jebog/stuble/midldlewares"
)

type MediaRoute struct {
	Path string
}

func NewMediaRoute(route *gin.Engine) {

	h := &MediaRoute{Path: configs.NewConfig().BasePath}
	r := route.Group(h.Path)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/medias", controllers.MediaController{}.Get)
	r.PUT("/medias/create", controllers.MediaController{}.Create)
	r.PATCH("/medias/update/:id", controllers.MediaController{}.Update)
	r.DELETE("/medias/delete", controllers.MediaController{}.Destroy)
}
