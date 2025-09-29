package main 

import (
	"net/http"
	"fmt"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func (app *application) routes() http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)

	return mux
}

