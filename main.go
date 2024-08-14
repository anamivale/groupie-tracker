package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
type Concerts struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	// Parse the JSON response into the struct
	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Access the members of the first artist (e.g., Queen)
	if len(artists) > 0 {
		for i, artist := range artists{
			fmt.Printf("Members of %s:\n", artist.Name)
			for _, member := range artist.Members {
				fmt.Println(member)
			}
			
			
			fmt.Printf("first Album: %s\n", artist.FirstAlbum)
			fmt.Printf("Band creation date: %d\n", artist.CreationDate)
			if i==3{
				break
			}
			println()


		}
		
		

	}
}
