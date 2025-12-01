package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	tags := queryParams["tags"]
	if tags == nil {
		tags = []string{}
	}

	data, err := app.getMainPage(tags)
	if err != nil {
		log.Printf("Error fetching events: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) search(w http.ResponseWriter, r *http.Request) {
	tags, err := app.getSearchParams()
	if err != nil {
		log.Printf("Error fetching tags: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

func (app *application) books(w http.ResponseWriter, r *http.Request) {
	data, err := app.getBooks()
	if err != nil {
		log.Printf("Error fetching books: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	app.render(w, http.StatusOK, "books.html", data)
}

func (app *application) contribute(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "contribute.html", nil)
}

func (app *application) eventDisplay(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	event, err := app.getEvent(id)
	if err != nil {
		log.Printf("Error fetching events: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	app.render(w, http.StatusOK, "event.html", event)
}

func (app *application) eventEdit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
	app.render(w, http.StatusOK, "event_edit.html", nil)
}
