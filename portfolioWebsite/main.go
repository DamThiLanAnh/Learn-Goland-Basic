package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Route handler for home page (root URL)
	http.HandleFunc("/", home)
	// Route handler for projects page (root URL)
	http.HandleFunc("/projects", projects)
	// Route handler for skills page (root URL)
	http.HandleFunc("/skills", skills)
	// Route handler for about page (root URL)
	http.HandleFunc("/about", about)
	// Route handler for contact page (root URL)
	http.HandleFunc("/contact", contact)

	fmt.Println("🚀Server is running on http://localhost:8080")

	//Start the server on port 8080 and listen to incoming request
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")
}

func skills(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
