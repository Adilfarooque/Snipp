package main

import (
	"log"
	"net/http"
)

const port = ":8080"

//Define a home handler func which writes a byte slice containing
func home(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("Hello from Snipp"))
}

func main() {
	//User the http.NewServeMux() function to initialize a new servemux , then
	// register the home func as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)

	log.Print("Starting server on",port)
	if err := http.ListenAndServe(port,mux); err != nil{
		log.Fatal(err)
	}
}
