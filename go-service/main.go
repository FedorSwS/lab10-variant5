package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	City   string `json:"city" binding:"required"`
	Street string `json:"street" binding:"required"`
}

type User struct {
	ID      int     `json:"id" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required,email"`
	Address Address `json:"address" binding:"required"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/api/users", func(c *gin.Context) {
		users := []User{
			{ID: 1, Name: "Alice", Email: "alice@example.com", Address: Address{City: "Moscow", Street: "Tverskaya"}},
			{ID: 2, Name: "Bob", Email: "bob@example.com", Address: Address{City: "Saint Petersburg", Street: "Nevsky"}},
		}
		c.JSON(http.StatusOK, users)
	})

	router.POST("/api/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	})

	router.Run(":8080")
}
