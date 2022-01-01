package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to main page.\n")
}

func HandleLog(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseGlob("/home/iojhanbal/P1/static/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
	log.Println("Entrando")

}
