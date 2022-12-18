// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/controllers"
	"github.com/jebog/stuble/helpers"
	middlewares "github.com/jebog/stuble/midldlewares"
)

type RoomRoute struct {
	Path string
}

func NewRoomRoute(route *gin.Engine) {

	h := &RoomRoute{Path: helpers.NewConfig().BasePath}
	r := route.Group(h.Path)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/rooms", controllers.RoomController{}.Get)
	r.PUT("/rooms/create", controllers.RoomController{}.Create)
	r.POST("/rooms/update", controllers.RoomController{}.Get)
	r.GET("/rooms/:id", controllers.RoomController{}.Destroy)
}
