package main

import (
	"fmt"
	"net/http"
)

func mux() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	GenerratingCert()
	server.ListenAndServeTLS("cert.pem", "key.pem")
}

type MyHandler struct{}

// That’s because ServeMux (which is what DefaultServeMux is an instance of) has a
// method named ServeHTTP with the same signature! In other words, a ServeMux is also
// an instance of the Handler struct. DefaultServeMux is an instance of ServeMux, so it is
// also an instance of the Handler struct. It’s a special type of handler, though, because
// the only thing it does is redirect your requests to different handlers depending on the
// URL that’s provided

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func ServeHTTPEx() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
