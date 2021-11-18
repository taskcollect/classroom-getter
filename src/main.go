package main

import (
	"io/ioutil"
	"log"
	"main/gc"
	"net/http"

	"github.com/buger/jsonparser"
)

func handler(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		log.Printf("400: Failed to read request body.")
	}

	secret, _, _, err := jsonparser.Get(req, "secret")
	if err != nil {
		w.WriteHeader(400)
		log.Printf("400: Failed to get secret from JSON request body. %v", err)
	}

	tasks, err := gc.GetTasks(secret)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("500: Failed to get tasks. %v", err)
	}

	w.Write(tasks)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting test server on port 2000.")
	http.ListenAndServe(":2000", nil)
}
