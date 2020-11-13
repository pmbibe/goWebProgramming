package main

import (
	"html/template"
	"net/http"
)

func process55(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template55.html")
	daysOfWeek := []string{"Mon", "Tue"}
	t.Execute(w, daysOfWeek)
}
