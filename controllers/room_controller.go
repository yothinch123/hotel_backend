package controllers

import (
	"go_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetRoom(c *gin.Context) {
	var user []models.Room

	users := models.GetAllRooms(&user)
	c.JSON(http.StatusOK, users)
}

// Get a single user by ID
func GetRoomByID(c *gin.Context) {
	var userdb models.Room
	id := c.Param("id")
	user := models.GetRoom(&userdb, id)

	c.JSON(http.StatusOK, user)
}

// Create a new user
func CreateRoom(c *gin.Context) {
	var user *models.Room = &models.Room{}

	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := models.AddRoom(user)
	c.JSON(http.StatusCreated, res)
}

// Update an existing user
func UpdateRoom(c *gin.Context) {
	var user *models.Room = &models.Room{}

	id := c.Param("id")
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := models.UpdateRoom(user, id)
	c.JSON(res.StatusCode, res)
}

// Delete a user
func DeleteRoom(c *gin.Context) {
	var userdb models.Room
	id := c.Param("id")
	res := models.DeleteRoom(&userdb, id)

	c.JSON(
		res.StatusCode,
		res,
	)
}
