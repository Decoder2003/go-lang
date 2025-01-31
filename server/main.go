package main

import (
	"fmt"
	"log"
	"net/http"
)

// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Golang Server!")
}

// About handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a simple local server built with Go.")
}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// start the server
	port := ":8000"
	fmt.Println("server is running on PORT" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
