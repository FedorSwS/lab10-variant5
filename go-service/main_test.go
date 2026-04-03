package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/api/users", func(c *gin.Context) {
		users := []User{
			{ID: 1, Name: "Alice", Email: "alice@example.com", Address: Address{City: "Moscow", Street: "Tverskaya"}},
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
	return router
}

func TestComplexJSONStructure(t *testing.T) {
	router := setupTestRouter()
	userJSON := `{"id": 1, "name": "Test User", "email": "test@example.com", "address": {"city": "Moscow", "street": "Lenina"}}`
	req, _ := http.NewRequest(http.MethodPost, "/api/users", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var created User
	json.Unmarshal(w.Body.Bytes(), &created)
	assert.Equal(t, "Test User", created.Name)
	assert.Equal(t, "Moscow", created.Address.City)
}

func TestNestedJSONValidation(t *testing.T) {
	router := setupTestRouter()
	userJSON := `{"id": 2, "name": "Invalid"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/users", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
