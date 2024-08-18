package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"net/http"
)

// Define a struct for the API response

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/location", handlers.LocationHandler)
	http.HandleFunc("/dates", handlers.HandleDates)
	http.HandleFunc("/relations", handlers.HandleRelation)
	go func() {
        if err := http.ListenAndServe(":3000", nil); err != nil {
            fmt.Println("Server failed:", err)
        }
    }()

    // Your main program logic here
    fmt.Println("Server started on port 3000")
	fmt.Println("http://localhost:3000")

    // Prevent the main function from exiting immediately
    select {}  // This will block forever
	
}
