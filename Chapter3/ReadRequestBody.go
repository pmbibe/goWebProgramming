package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))

}
func ReadRequestBody() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
