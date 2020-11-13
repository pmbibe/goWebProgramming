package main

import (
	"fmt"
	"net/http"
)

func WriteHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
}
func HeadervsWriteHeader() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/writeHeader", WriteHeaderExample)
	http.HandleFunc("/", headerExample)
	server.ListenAndServe()

}
