package main

import (
	"io/fs"
	"net/http"
)


func (app *application) routes() http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.homeHandler)

	staticFS, _ := fs.Sub(Files, "ui/static")
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	return mux
}

