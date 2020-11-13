package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process528(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("multipleTemplates.html", "redHello.html")
	} else {
		t, _ = template.ParseFiles("multipleTemplates.html", "blueHello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}
