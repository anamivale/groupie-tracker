package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

// Define a struct for the API response
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func Handler(w http.ResponseWriter, r *http.Request){

	url := "https://groupietrackers.herokuapp.com/api/artists"
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	// Parse the JSON response into the struct
	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	temp, _ := template.ParseFiles("index.html")
	temp.Execute(w, artists)
}


func main() {

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}