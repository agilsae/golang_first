package router

import (
	"diary-api/healthcheck"

	"github.com/gin-gonic/gin"
)

func healthCheck(router *gin.RouterGroup){
	router.GET("/healthz", healthcheck.DefaultCheck)
}