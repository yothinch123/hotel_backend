package controllers

import (
	"go_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetRoomType(c *gin.Context) {
	var user []models.RoomType

	users := models.GetAllRoomTypes(&user)
	c.JSON(http.StatusOK, users)
}

// Get a single user by ID
func GetRoomTypeByID(c *gin.Context) {
	var userdb models.RoomType
	id := c.Param("id")
	user := models.GetRoomType(&userdb, id)

	c.JSON(http.StatusOK, user)
}

// Create a new user
func CreateRoomType(c *gin.Context) {
	var user *models.RoomType = &models.RoomType{}

	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := models.AddRoomType(user)
	c.JSON(http.StatusCreated, res)
}

// Update an existing user
func UpdateRoomType(c *gin.Context) {
	var user *models.RoomType = &models.RoomType{}

	id := c.Param("id")
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := models.UpdateRoomType(user, id)
	c.JSON(res.StatusCode, res)
}

// Delete a user
func DeleteRoomType(c *gin.Context) {
	var userdb models.RoomType
	id := c.Param("id")
	res := models.DeleteRoomType(&userdb, id)

	c.JSON(
		res.StatusCode,
		res,
	)
}
