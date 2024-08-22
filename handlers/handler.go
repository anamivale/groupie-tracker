package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

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
type Dates struct {
	Date []string `json:"dates"`
}

type LocationResponse struct {
	Locations []string `json:"locations"`
}
type Relation struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}

// Handler for the main page displaying artists
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, 404)
		return
	}
	if r.Method != http.MethodGet {
		Errors(w, 405)
		return
	}
	url := "https://groupietrackers.herokuapp.com/api/artists"
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		// log.Fatalf("Failed to make request: %v", err)
		StatusInternalServerError(w)

		return
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
		StatusInternalServerError(w)
		return

	}
	// Parse the JSON response into the struct
	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	temp, _ := template.ParseFiles("template/index.html")
	temp.Execute(w, artists)
}

// Handler for the locations page
func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/location" {
		Errors(w, 404)
		return

	}
	if r.Method != http.MethodGet {
		Errors(w, 405)
		return

	}
	// Get the artist ID from the query parameters
	artistID := r.URL.Query().Get("id")
	url := "https://groupietrackers.herokuapp.com/api/locations/" + artistID

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		// log.Fatalf("Failed to make request: %v", err)
		StatusInternalServerError(w)
		return

	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
		StatusInternalServerError(w)
		return

	}

	// Parse the JSON response into the struct
	var locationResponse LocationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Render the locations template
	temp, _ := template.ParseFiles("template/location.html")
	temp.Execute(w, locationResponse.Locations)
}

func HandleDates(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dates" {
		Errors(w, 404)
		return

	}
	if r.Method != http.MethodGet {
		Errors(w, 405)
		return

	}

	Id := r.URL.Query().Get("id")
	url := "https://groupietrackers.herokuapp.com/api/dates/" + Id
	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
		StatusInternalServerError(w)
		return

	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
		StatusInternalServerError(w)
		return

	}

	var dates Dates
	err = json.Unmarshal(body, &dates)
	for i, s := range dates.Date {
		dates.Date[i] = strings.ReplaceAll(s, "*", "")
	}
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	temp, _ := template.ParseFiles("template/date.html")
	temp.Execute(w, dates.Date)
}

func HandleRelation(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/relations" {
		Errors(w, 404)
		return

	}
	if r.Method != http.MethodGet {
		Errors(w, 405)
		return

	}

	Id := r.URL.Query().Get("id")

	url := "https://groupietrackers.herokuapp.com/api/relation/" + Id

	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Fail to get: %v", err.Error())
		StatusInternalServerError(w)
		return

	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Fail to read: %v", err.Error())
		StatusInternalServerError(w)
		return

	}
	var relation Relation

	err = json.Unmarshal(body, &relation)

	if err != nil {
		log.Fatalf("Fail to unmarshal: %v", err.Error())
	}

	temp, err := template.ParseFiles("template/relation.html")
	if err != nil {
		log.Fatalf("Fail to parse file: %v", err.Error())
	}
	err = temp.Execute(w, relation)
	if err != nil {
		log.Fatalf("Fail to execute file: %v", err.Error())
	}

}
