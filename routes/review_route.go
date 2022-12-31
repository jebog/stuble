package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/configs"
	"github.com/jebog/stuble/controllers"
	middlewares "github.com/jebog/stuble/midldlewares"
)

type ReviewRoute struct {
	Path string
}

func NewReviewRoute(route *gin.Engine) {

	h := &ReviewRoute{Path: configs.NewConfig().BasePath}
	r := route.Group(h.Path)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/reviews", controllers.ReviewController{}.Get)
	r.PUT("/reviews/create", controllers.ReviewController{}.Create)
	r.PATCH("/reviews/update/:id", controllers.ReviewController{}.Update)
	r.DELETE("/reviews/delete", controllers.ReviewController{}.Destroy)
}
