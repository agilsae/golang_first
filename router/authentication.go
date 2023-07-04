package router

import (
	"diary-api/controller"

	"github.com/gin-gonic/gin"
)

func authenticationApi(router *gin.RouterGroup) {
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}