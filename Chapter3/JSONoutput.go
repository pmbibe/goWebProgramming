package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func JSONExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
func JSONoutput() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/json", JSONExample)
	server.ListenAndServe()
}
