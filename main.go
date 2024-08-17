package main

import (
	"groupie-tracker/handlers"
	"net/http"
)

// Define a struct for the API response

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/location", handlers.LocationHandler)
	http.HandleFunc("/dates", handlers.HandleDates)
	http.HandleFunc("/relations", handlers.HandleRelation)
	http.ListenAndServe(":3000", nil)
	
}
