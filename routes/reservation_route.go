package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/controllers"
	"github.com/jebog/stuble/helpers"
	middlewares "github.com/jebog/stuble/midldlewares"
)

type ReservationRoute struct {
	Path string
}

func NewReservationRoute(route *gin.Engine) {

	h := &ReservationRoute{Path: helpers.NewConfig().BasePath}
	r := route.Group(h.Path)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/reservations", controllers.ReservationController{}.Get)
	r.PUT("/reservations/create", controllers.ReservationController{}.Create)
	r.PATCH("/reservations/update/:id", controllers.ReservationController{}.Update)
	r.DELETE("/reservations/delete", controllers.ReservationController{}.Destroy)
}
