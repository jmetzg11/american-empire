package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	events, err := app.getMainPage()
	if err != nil {
		log.Printf("Error fetching events: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	app.render(w, http.StatusOK, "home.tmpl", events)
}

func (app *application) contribute(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "contribute.tmpl", nil)
}

func (app *application) eventDisplay(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	event, err := app.getEvent(id)
	if err != nil {
		log.Printf("Error fetching events: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	app.render(w, http.StatusOK, "event.tmpl", event)
}

func (app *application) eventEdit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
	app.render(w, http.StatusOK, "event_edit.tmpl", nil)
}
