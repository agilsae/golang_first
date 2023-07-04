package router

import (
	"diary-api/middleware"
	"fmt"
	"net/http"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
)

func ServeApplication() {
	router := gin.Default()

	router.Use(healthcheck.New(healthcheck.Config{
		HeaderName:   "X-Custom-Header",
		HeaderValue:  "customValue",
		ResponseCode: http.StatusTeapot,
		ResponseText: "teapot",
	}))

	router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("Hello Golang")
	})

	host := router.Group("/sae")

	publicRoutes := host.Group("/auth")
	privateRoutes := host.Group("/api")
	privateRoutes.Use(middleware.JWTAuthMiddleware())

	healthCheck(host)
	authenticationApi(publicRoutes)
	books(privateRoutes)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}