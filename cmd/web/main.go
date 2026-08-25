package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// create file server for serving static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// register the file server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", getSnippetCreate)
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	log.Print("Starting serverr on port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
