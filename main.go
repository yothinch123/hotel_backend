package main

import (
	"log"
	"net/http"
	"os"

	"go_backend/config"
	"go_backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	initDB()

	// Define routes for user API
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUserByID)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.Run(":8080")
}

func initDB() {
	err := godotenv.Load()

	dbHose := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHose + ":3306)/" + dbName + "?parseTime=true"
	config.DB, err = gorm.Open("mysql", dsn)

	config.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	// Test connection to the database
	log.Println("Database connection established")
}

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

	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := models.UpdateUser(user)
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
