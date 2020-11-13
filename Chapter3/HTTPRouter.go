package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func helloWW(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
	fmt.Fprintln(w, r.Header)
}
func HTTPRouter() {
	mux := httprouter.New()
	mux.GET("/hello/:name", helloWW)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
