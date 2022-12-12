package controllers

import "github.com/gin-gonic/gin"

type BaseController interface {
	Get(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Destroy(context *gin.Context)
}
