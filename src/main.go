package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting test server on port 2000.")
	http.ListenAndServe(":2000", nil)
}
