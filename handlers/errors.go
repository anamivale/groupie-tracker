package handlers

import (
	"fmt"
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
