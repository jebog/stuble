// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/controllers"
)

type AuthRoute struct {
}

func NewAuthRoute(route *gin.Engine) {

	route.POST("/register", controllers.AuthController{}.Register)
	route.POST("/login", controllers.AuthController{}.Login)
}