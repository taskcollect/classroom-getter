package gc

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"
)

func getClient(config *oauth2.Config, secret []byte) (*http.Client, error) {
	tok, err := tokenFromFile(secret)
	if err != nil {
		return nil, err
	}
	return config.Client(context.Background(), tok), nil
}

func tokenFromFile(secret []byte) (*oauth2.Token, error) {
	tok := &oauth2.Token{}
	err := json.Unmarshal(secret, tok)
	return tok, err
}
