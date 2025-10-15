package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-api-backend/internal/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserEndpoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.SetupRoutes()

	// Test GET /users
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Additional assertions can be added here based on the expected response
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.SetupRoutes()

	// Test POST /users
	body := `{"username":"alice","email":"alice@example.com"}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	// Additional assertions can be added here based on the expected response
}
