package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process529(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("usingBlockAcction.html", "redHello.html")
	} else {
		t, _ = template.ParseFiles("usingBlockAcction.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}
