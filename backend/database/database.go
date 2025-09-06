package database

import (
	"log"
	"os"

	"github.com/supabase-community/supabase-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB             *gorm.DB
	SupabaseClient *supabase.Client
)

func Connect() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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
	supabaseKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
	var err error
	SupabaseClient, err = supabase.NewClient(supabaseURL, supabaseKey, nil)
	return err
}
