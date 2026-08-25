package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to the home page"))
}
func main() {

}
