package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/configs"
	"github.com/jebog/stuble/controllers"
	middlewares "github.com/jebog/stuble/midldlewares"
)

type RoomRoute struct {
	Path string
}

func NewRoomRoute(route *gin.Engine) {

	h := &RoomRoute{Path: configs.NewConfig().BasePath}
	r := route.Group(h.Path)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/rooms", controllers.RoomController{}.Get)
	r.PUT("/rooms/create", controllers.RoomController{}.Create)
	r.PATCH("/rooms/update/:id", controllers.RoomController{}.Update)
	r.DELETE("/rooms/delete", controllers.RoomController{}.Destroy)
}
