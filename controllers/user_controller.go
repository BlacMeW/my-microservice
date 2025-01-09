package controllers

import (
	"my-microservice/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := services.GetUsers()
	c.JSON(http.StatusOK, users)
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
