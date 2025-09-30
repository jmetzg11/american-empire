package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	db *sql.DB
}

func main() {
	// Parse command line flags
	prod := flag.Bool("prod", false, "Use production environment")
	flag.Parse()

	// Connect to database
	db, err := connectDB(*prod)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app := &application{
		db: db,
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server starting on :8080")
	log.Fatal(srv.ListenAndServe())
}

func connectDB(prod bool) (*sql.DB, error) {
	var dsn string
	if prod {
		if err := godotenv.Load("../.env"); err != nil {
			return nil, fmt.Errorf("failed to load .env file: %w", err)
		}
		dsn = os.Getenv("DATABASE_URL")
		if dsn == "" {
			return nil, fmt.Errorf("DATABASE_URL environment variable not set")
		}
		fmt.Println("Connecting to production database")
	} else {
		dsn = "postgresql://admin:admin@localhost:5432/american_empire?sslmode=disable"
		fmt.Println("Connecting to local development database")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	fmt.Println("Database connection established")

	return db, nil
}

