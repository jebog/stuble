// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/controllers"
)

type AuthRoute struct {
	Path string
}

func NewAuthRoute(route *gin.Engine) {

	h := &AuthRoute{}
	r := route.Group(h.Path)
	r.POST("/register", controllers.AuthController{}.Register)
	r.POST("/login", controllers.AuthController{}.Login)
}
