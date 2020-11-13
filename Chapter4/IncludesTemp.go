package main

import (
	"html/template"
	"net/http"
)

func process59(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template59.html", "template510.html")
	t.Execute(w, "HelloWorld!")
}
