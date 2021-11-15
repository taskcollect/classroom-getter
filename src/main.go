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
		log.Fatal("400: Failed to read request body.")
	}

	secret, _, _, err := jsonparser.Get(req, "secret")
	if err != nil {
		log.Fatalf("400: Failed to get secret from JSON request body. %v", err)
	}

	tasks := gc.GetTasks(secret)
	w.Write(tasks)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting test server on port 2000.")
	http.ListenAndServe(":2000", nil)
}
