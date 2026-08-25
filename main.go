package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to the home page"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fmt.Println("Server listening on port 4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
