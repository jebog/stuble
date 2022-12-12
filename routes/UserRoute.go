// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/controllers"
	middlewares "github.com/jebog/stuble/midldlewares"
)

type UserRoute struct {
	Path string
}

func NewUserRoute(route *gin.Engine) {

	h := &UserRoute{}

	r := route.Group(h.Path)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/users", controllers.UserController{}.Get)
	r.PUT("/users/create", controllers.UserController{}.Create)
	r.POST("/users/update", controllers.UserController{}.Update)
	r.GET("/users/{id}", controllers.UserController{}.Destroy)
}
