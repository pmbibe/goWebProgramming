package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
func process514(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("template514.html").Funcs(funcMap)
	t, _ = t.ParseFiles("template514.html")
	t.Execute(w, time.Now())
}
