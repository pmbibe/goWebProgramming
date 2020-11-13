package main

import (
	"html/template"
	"net/http"
)

func process57(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template57Adv.html")
	t.Execute(w, "hello")
}
