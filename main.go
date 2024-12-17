package main

import (
	"fmt"
	"groupie-tracker/handlers"
	"groupie-tracker/models"
	"html/template"
	"net/http"
)

func main() {

	models.Tmpl = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", handlers.HomeHandler)

	fmt.Printf("Server is running on http://localhost:10000")

	http.ListenAndServe(":10000", nil)

}
