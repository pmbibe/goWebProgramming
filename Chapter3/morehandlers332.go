package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}
type WordHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
func (h *WordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL)
}
func MoreHandlers() {
	var hello HelloHandler
	var word WordHandler
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	mux.Handle("/world", &hello)
	mux.Handle("/hello", &word)
	server.ListenAndServe()

}
