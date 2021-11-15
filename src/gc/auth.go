package gc

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func getClient(config *oauth2.Config, secret []byte) *http.Client {
	tok, err := tokenFromFile(secret)
	if err != nil {
		log.Printf("500: Auth error: %v", err)
	}
	return config.Client(context.Background(), tok)
}

func tokenFromFile(secret []byte) (*oauth2.Token, error) {
	tok := &oauth2.Token{}
	err := json.Unmarshal(secret, tok)
	return tok, err
}
