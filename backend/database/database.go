package database

import (
	"os"

	"github.com/supabase-community/supabase-go"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"log"
)

var DB *gorm.DB
var SupabaseClient *supabase.Client

func Connect() error {
	var err error

	if os.Getenv("GIN_MODE") == "release" {
		dsn := os.Getenv("DATABASE_URL")
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	} else {
		if err := os.MkdirAll("data", 0755); err != nil {
			log.Fatal("Failed to create data directory", err)
		}

		DB, err = gorm.Open(sqlite.Open("data/american-empire.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	return nil
}

func InitSupabase() error {
	if os.Getenv("GIN_MODE") != "release" {
		return nil
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")
	var err error
	SupabaseClient, err = supabase.NewClient(supabaseURL, supabaseKey, nil)
	return err
}
