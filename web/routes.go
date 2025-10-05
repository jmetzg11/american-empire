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
	mux.HandleFunc("GET /event/{id}", app.eventHandler)

	return rateLimit(mux)
}
