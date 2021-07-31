package main

import (
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			log.Fatal(w, err.Error())
			return
		}
		t.ExecuteTemplate(w, "index", nil)
		return

	}

}

func donateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		t, err := template.ParseFiles("templates/donate.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			log.Fatal(w, err.Error())
			return
		}
		t.ExecuteTemplate(w, "donate", nil)
		return

	}

}
