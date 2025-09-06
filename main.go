package main

import (
	"american-empire/backend/database"
	"american-empire/backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Using environment variables.")

	}

	database.Connect()
	fmt.Println("Database connected successfully!")

	if err := database.InitSupabase(); err != nil {
		log.Fatal("Failed to initialize Supabase:", err)
	}

	r := gin.Default()
	if os.Getenv("GIN_MODE") != "release" {
		r.Static("/photos", "./data/photos")
	}
	routes.SetupAPIRoutes(r)

	r.Run(":8080")
}
