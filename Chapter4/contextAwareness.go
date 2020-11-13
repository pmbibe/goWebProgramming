package main

import (
	"html/template"
	"net/http"
)

func process517(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template517.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}
