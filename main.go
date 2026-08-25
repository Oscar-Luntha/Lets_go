package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to the Snippet home page"))
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Specific snippet view"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form for creating a snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fmt.Println("Server listening on port 4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
