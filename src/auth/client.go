package auth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/classroom/v1"
	"google.golang.org/api/option"
)

func GetClient(config *oauth2.Config, token *oauth2.Token) (*http.Client, error) {
	return config.Client(context.Background(), token), nil
}

func GetService(client *http.Client) (*classroom.Service, error) {
	return classroom.NewService(
		context.Background(),
		option.WithHTTPClient(client),
	)
}
