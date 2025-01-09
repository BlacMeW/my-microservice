package routes

import (
	"my-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
}
