package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type application struct {
	db            *sql.DB
	templateCache map[string]*template.Template
}

func main() {
	prod := flag.Bool("prod", false, "Use production environment")
	flag.Parse()

	db, err := connectDB(*prod)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatalf("Failed to create template cache: %v", err)
	}

	app := &application{
		db:            db,
		templateCache: templateCache,
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



