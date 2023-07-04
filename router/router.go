package router

import (
	"diary-api/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ServeApplication() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("Hello Golang")
	})

	publicRoutes := router.Group("/auth")
	privateRoutes := router.Group("/api")
	privateRoutes.Use(middleware.JWTAuthMiddleware())

	authenticationApi(publicRoutes)
	books(privateRoutes)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}