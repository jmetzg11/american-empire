package main

import (
	"american-empire/backend/database"
	"american-empire/backend/routes"
	"fmt"
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

	if err := database.InitSupabase(); err != nil {
		log.Fatal("Failed to initialize Supabase:", err)
	}

	r := gin.Default()
	r.Static("/photos", "./data/photos")
	routes.SetupStaticRoutes(r)
	routes.SetupAPIRoutes(r)

	r.Run(":3000")
}
