package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Mix it well"))
}

func snippView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snipp view..."))
}

func snippCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snipp"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/snipp/view", snippView)
	mux.HandleFunc("/snipp/create", snippCreate)

	log.Print("Starting server on", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
