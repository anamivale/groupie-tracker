package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func StatusInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := os.Stat("template/500.html"); os.IsNotExist(err) {
		fmt.Println("template not found")
		return
	}
	tmp, err := template.ParseFiles("template/500.html")
	if err != nil {
		http.Error(w, "Internal Server Error1", http.StatusInternalServerError)
	}
	tmp.Execute(w, nil)

}

func Errors(w http.ResponseWriter, statusCode int) {
	var msg string

	switch statusCode {
	case 405:
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg = "The request method you are making is not allowed."
	case 404:
		w.WriteHeader(http.StatusNotFound)
		msg = "The file you are trying to access is not found."
	case 400:
		w.WriteHeader(http.StatusBadRequest)
		msg = "Bad request, please correct your request."
	default:
		w.WriteHeader(http.StatusInternalServerError)
		msg = "An unknown error occurred."
	}

	if _, err := os.Stat("template/errors.html"); os.IsNotExist(err) {
		fmt.Println("Template not found")
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	tmp, err := template.ParseFiles("template/errors.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(w, msg)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
