package main

import (
	"html/template"
	"net/http"
)

func process525(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("multipleTemplates.html")
	t.ExecuteTemplate(w, "layout", "")
}
