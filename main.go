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
	// r.PUT("/users/:id", UpdateUser)
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
// func UpdateUser(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, user := range users {
// 		if strconv.Itoa(user.ID) == id {
// 			var updatedUser User

// 			if err := c.BindJSON(&updatedUser); err != nil {
// 				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			updatedUser.ID = user.ID
// 			users[i] = updatedUser

// 			c.JSON(http.StatusOK, updatedUser)
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
// }

// Delete a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	res := models.DeleteUser(id)

	c.JSON(
		res.StatusCode,
		res,
	)
}
