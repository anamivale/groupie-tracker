package main

import (
	"groupie-tracker/handlers"
	"net/http"
)

// Define a struct for the API response

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/location", handlers.LocationHandler)
	
	http.ListenAndServe(":3000", nil)
	
}
