package main

import "net/http"

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", process)
	http.HandleFunc("/53", process53)
	http.HandleFunc("/55", process55)
	http.HandleFunc("/55Adv", process55)
	http.HandleFunc("/57", process57)
	http.HandleFunc("/59", process59)
	http.HandleFunc("/514", process514)
	http.HandleFunc("/517", process517)
	http.HandleFunc("/form", form)
	http.HandleFunc("/519", process519)
	http.HandleFunc("/525", process525)
	http.HandleFunc("/528", process528)
	http.HandleFunc("/529", process529)
	server.ListenAndServe()
}
