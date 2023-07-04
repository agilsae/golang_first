package router

import (
	"diary-api/controller"
	"github.com/gin-gonic/gin"
)

func books(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	v1.POST("/entry", controller.AddEntry)
	v1.GET("/entry", controller.GetAllEntries)
}