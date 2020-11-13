package main

import (
	"html/template"
	"net/http"
)

func process519(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template519.html")
	t.Execute(w, r.FormValue("comment"))
	//Unescaping HTML
	t.Execute(w, template.HTML(r.FormValue("comment")))
	//Turn off XSS Protection
	//w.Header().Set("X-XSS-Protection", "0")

}
func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}
