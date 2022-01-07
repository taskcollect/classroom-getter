package auth

import "os"

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

func GetFromEnv() *OAuth2Secrets {
	return &OAuth2Secrets{
		ClientID:     getenv_safe("CLIENT_ID"),
		ClientSecret: getenv_safe("CLIENT_SECRET"),
	}
}
