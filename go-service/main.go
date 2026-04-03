package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Printf("Go Gin server starting on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Graceful shutdown initiated...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server stopped")
}
