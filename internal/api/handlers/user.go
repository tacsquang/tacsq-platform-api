package handlers

import (
	"net/http"
	"sync"

	"go-api-backend/internal/models"

	"github.com/gin-gonic/gin"
)

var (
	// simple in-memory store
	usersMu sync.Mutex
	users   = []*models.User{}
)

// GetUsers returns the list of users
func GetUsers(c *gin.Context) {
	usersMu.Lock()
	defer usersMu.Unlock()

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// CreateUser creates a user from JSON body
func CreateUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersMu.Lock()
	defer usersMu.Unlock()

	// assign a simple ID
	u.ID = string(len(users) + 1)
	users = append(users, &u)

	c.JSON(http.StatusCreated, gin.H{"user": u})
}
