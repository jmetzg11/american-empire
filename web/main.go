package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type application struct {}

func main() {
	app := &application{}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server starting on :8080")
	log.Fatal(srv.ListenAndServe())
}

