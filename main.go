package main

import (
	"fmt"
	"good-guys/backend/database"
	"good-guys/backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Using environment variables.")
	}

	plainPassword := os.Getenv("ADMIN_PASSWORD")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	os.Setenv("ADMIN_PASSWORD_HASHED", string(hashedPassword))

	database.Connect()
	fmt.Println("Database connected successfully!")
	r := gin.Default()

	r.Static("/photos", "./data/photos")
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	// Start the server
	r.Run(":3000")
}
