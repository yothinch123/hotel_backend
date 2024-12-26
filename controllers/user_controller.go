package controllers

import (
	"go_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
	var user []models.User

	users := models.GetAllUsers(&user)
	c.JSON(http.StatusOK, users)
}

// Get a single user by ID
func GetUserByID(c *gin.Context) {
	var userdb models.User
	id := c.Param("id")
	user := models.GetUser(&userdb, id)

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

	id := c.Param("id")
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := models.UpdateUser(user, id)
	c.JSON(res.StatusCode, res)
}

// Delete a user
func DeleteUser(c *gin.Context) {
	var userdb models.User
	id := c.Param("id")
	res := models.DeleteUser(&userdb, id)

	c.JSON(
		res.StatusCode,
		res,
	)
}
