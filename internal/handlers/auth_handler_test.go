package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestLoginHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/login", LoginHandler)

	loginPayload := `{"username":"john", "password":"setup123"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response JSON
	var jsonResponse map[string]interface{}
	err := json.Unmarshal(responseData, &jsonResponse)
	assert.NoError(t, err, "Response body should be valid JSON")

	// Assert that the token is present
	token, tokenExists := jsonResponse["token"]
	assert.True(t, tokenExists, "Token should be present in the response")

	// Assert that the token is a non-empty string
	assert.IsType(t, "", token, "Token should be a string")
	assert.NotEmpty(t, token, "Token should not be empty")
}
