package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template



func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w , "home.html", nil )
}


func main() {


 tpl = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", homeHandler)

	fmt.Printf("Server is running on http://localhost:10000")

	 http.ListenAndServe(":10000", nil)

}
