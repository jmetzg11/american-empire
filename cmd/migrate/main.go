package main

import (
	"american-empire/backend/database"
	"american-empire/backend/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	if err := database.Connect(): err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.DB.AutoMigrate(
		&models.Event{},
		&models.Source{},
		&models.Media{},
		&models.Tag{},
	)

	log.Println("Migration completed successfully!")
}
