package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"taxi-service/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a valid JWT token
func createValidToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "12345",
		"exp":     time.Now().Add(time.Minute * 5).Unix(), // Token expiration
	})

	tokenString, _ := token.SignedString(middlewares.JWTSecret)
	return tokenString
}

func TestJWTAuthMiddleware_ValidToken(t *testing.T) {
	// Create a valid token
	tokenString := createValidToken()

	// Set up the Gin engine with the JWT middleware
	r := gin.Default()
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/api/v1/me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	// Create a new HTTP request with the Bearer token
	req, _ := http.NewRequest("GET", "/api/v1/me", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	// Create a response recorder to capture the middleware's response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check if the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Check if the response body contains the success message
	assert.Contains(t, w.Body.String(), "Success")
}

func TestJWTAuthMiddleware_InvalidToken(t *testing.T) {
	// Set up the Gin engine with the JWT middleware
	r := gin.Default()
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/api/v1/me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	// Create a new HTTP request with an invalid Bearer token
	req, _ := http.NewRequest("GET", "/api/v1/me", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")

	// Create a response recorder to capture the middleware's response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check if the response status code is 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Check if the response body contains the error message
	assert.Contains(t, w.Body.String(), "Invalid token")
}

func TestJWTAuthMiddleware_MissingToken(t *testing.T) {
	// Set up the Gin engine with the JWT middleware
	r := gin.Default()
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/api/v1/me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	// Create a new HTTP request without a Bearer token
	req, _ := http.NewRequest("GET", "/api/v1/me", nil)

	// Create a response recorder to capture the middleware's response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check if the response status code is 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Check if the response body contains the error message
	assert.Contains(t, w.Body.String(), "Missing Authorization Header")
}
