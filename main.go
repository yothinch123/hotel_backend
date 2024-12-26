package main

import (
	"fmt"
	"log"
	"os"

	"go_backend/config"
	"go_backend/models"
	"go_backend/routers"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	initDB()
	r := routers.SetupRouter()
	r.Run(":8080")
}

func initDB() {
	err := godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)
	config.DB, err = gorm.Open("mysql", dsn)

	config.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	// Test connection to the database
	log.Println("Database connection established")
}
