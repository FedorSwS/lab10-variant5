package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/api/users", func(c *gin.Context) {
		users := []string{"Alice", "Bob", "Charlie"}
		c.JSON(http.StatusOK, users)
	})

	router.POST("/api/users", func(c *gin.Context) {
		var user map[string]interface{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	})

	router.Run(":8080")
}
