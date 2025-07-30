package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", about)
	http.HandleFunc("/skills", skills)
	http.HandleFunc("/experience", experience)
	http.HandleFunc("/education", education)
	http.HandleFunc("/project", project)
	http.HandleFunc("/contact", contact)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Index page served successfully")
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func skills(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "skills.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Skills page served successfully")
}
func experience(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "experience.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Experience page served successfully")
}

func education(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "education.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Education page served successfully")
}
func project(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "project.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Projects page served successfully")

}

func contact(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Contact page served successfully")
}
