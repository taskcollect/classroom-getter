package auth

import (
	"errors"
	"os"
)

type OAuth2Secrets struct {
	ClientID     string
	ClientSecret string
}

func getenv_safe(name string) string {
	// fn to get env var at process startup, and panic
	// if it doesn't exist. should be used for required env vars
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	panic("environment variable " + name + " not specified")
}

func GetSecretsFromEnv() (*OAuth2Secrets, error) {
	id := os.Getenv("CLIENT_ID")
	if id == "" {
		return nil, errors.New("CLIENT_ID not set or empty in environment")
	}

	secret := os.Getenv("CLIENT_SECRET")
	if secret == "" {
		return nil, errors.New("CLIENT_SECRET not set or empty in environment")
	}

	return &OAuth2Secrets{
		ClientID:     id,
		ClientSecret: secret,
	}, nil
}
