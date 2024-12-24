package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"go_backend/models"

	"github.com/gin-gonic/gin"
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

var db *sql.DB

func initDB() {
	err := godotenv.Load()

	dbHose := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHose + ":3306)/" + dbName
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// Test connection to the database
	if err := db.Ping(); err != nil {
		log.Fatalf("Error testing database connection: %s", err)
	}
	log.Println("Database connection established")
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
