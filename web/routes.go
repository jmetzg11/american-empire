package main

import (
	"io/fs"
	"net/http"
)

func (app *application) routes(prod bool) http.Handler {
	mux := http.NewServeMux()

	staticFS, _ := fs.Sub(Files, "ui/static")
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// Serve photos locally in development only
	if !prod {
		photosFS := http.Dir("../data/photos")
		mux.Handle("GET /photos/", http.StripPrefix("/photos/", http.FileServer(photosFS)))
	}

	mux.HandleFunc("GET /{$}", app.homeHandler)
	mux.HandleFunc("GET /search", app.search)
	mux.HandleFunc("GET /books", app.books)
	mux.HandleFunc("GET /contribute", app.contribute)
	mux.HandleFunc("GET /event/{id}", app.eventDisplay)
	mux.HandleFunc("GET /event/edit/{id}", app.eventEdit)

	return rateLimit(mux)
}
