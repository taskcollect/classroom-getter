package main

import (
	"log"
	"main/auth"
	"main/call"
	"main/fetch"
	"time"

	"golang.org/x/oauth2"
)

func main() {
	secrets := auth.GetFromEnv()
	config := auth.GetOAuth2Config(secrets, auth.TC_API_SCOPES)

	token := &oauth2.Token{
		AccessToken:  "get this from gauthman",
		RefreshToken: "get this from gauthman",
		Expiry:       time.Unix(1641532911, 0).UTC(),
	}

	client, err := auth.GetClient(config, token)

	if err != nil {
		log.Fatalf("Failed to get client: %v", err)
	}

	srv, err := auth.GetService(client)
	if err != nil {
		log.Fatalf("Failed to get service: %v", err)
	}

	courses, err := call.ListCourses(srv)
	if err != nil {
		log.Fatalf("Failed to get courses: %v", err)
	}

	start := time.Now()
	assignments, err := fetch.FetchAllRelevant(srv, courses)
	if err != nil {
		log.Fatalf("Failed to fetch assignments: %v", err)
	}
	elapsed := time.Since(start)
	log.Printf("Fetched %d assignments in %s", len(assignments), elapsed)

	for _, assignment := range assignments {
		log.Printf("%v", assignment.Work.Title)
	}
}
