package main

import (
	"net/http"

	"go_backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define routes for user API
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUserByID)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.Run(":8080")
}

// Get all users
func GetUsers(c *gin.Context) {
	users := models.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// Get a single user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user := models.GetUser(id)

	c.JSON(http.StatusOK, user)
}

// Create a new user
func CreateUser(c *gin.Context) {
	var user *models.User = &models.User{}

	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := models.AddUser(user)
	c.JSON(http.StatusCreated, res)
}

// Update an existing user
func UpdateUser(c *gin.Context) {
	var user *models.User = &models.User{}

	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := models.UpdateUser(user)
	c.JSON(res.StatusCode, res)
}

// Delete a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	res := models.DeleteUser(id)

	c.JSON(
		res.StatusCode,
		res,
	)
}
