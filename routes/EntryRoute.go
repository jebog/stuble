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

type EntryRoute struct {
	Path string
}

func NewEntryRoute(route *gin.Engine) {

	h := &EntryRoute{Path: helpers.NewConfig().BasePath}
	r := route.Group(h.Path)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/entries", controllers.UserController{}.Get)
	r.PUT("/entries/create", controllers.UserController{}.Create)
	r.POST("/entries/update", controllers.UserController{}.Get)
	r.GET("/entries/{id}", controllers.UserController{}.Destroy)
}
