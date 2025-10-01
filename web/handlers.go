package main 

import (
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
