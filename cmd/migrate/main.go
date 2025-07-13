package main

import (
	"american-empire/backend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// For local development, run:
// go run cmd/migrate/main.go
// For production, run:
// go run cmd/migrate/main.go prod
func main() {
	env := "dev"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	var db *gorm.DB
	var err error

	if env == "prod" {
		dsn := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	} else {
		if err := os.MkdirAll("data", 0755); err != nil {
			log.Fatal("Failed to create data directory", err)
		}

		db, err = gorm.Open(sqlite.Open("data/american-empire.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(
		&models.Event{},
		&models.Source{},
		&models.Media{},
		&models.Tag{},
	)

	log.Printf("Migration completed successfully for %s environment!", env)
}
