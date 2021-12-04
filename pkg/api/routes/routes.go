package routes

import (
	"gcp-golang/resources/projects"
	"html/template"
	"log"
	"net/http"
)

// tmpl is a map of all the templates used in the application.
var tmpl = template.Must(template.ParseGlob("pkg/api/templates/*.html"))

// CallRoutes is the main function that handles all the routes
func CallRoutes() {
	http.HandleFunc("/", Index)
}

// Index is the main page
func Index(w http.ResponseWriter, r *http.Request) {

	p1 := projects.AllProjects{}

	sliceProducts := p1

	err := tmpl.ExecuteTemplate(w, "index.html", sliceProducts)
	if err != nil {
		log.Printf("Error executing t: %v", err)
		return
	}
}
