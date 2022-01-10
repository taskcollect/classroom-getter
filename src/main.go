package main

import (
	"log"
	"main/auth"
	"main/handlers"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

type ServerConfig struct {
	BindAddr string
	OAuth2   *oauth2.Config
}

// server config, values here will get overriden by env
var config = ServerConfig{
	BindAddr: "0.0.0.0:2000",
	OAuth2:   nil,
}

func makeMux() *http.ServeMux {
	mux := http.NewServeMux()

	handler := handlers.NewBaseHandler(config.OAuth2)

	mux.HandleFunc("/v1/tasks", handler.ActiveTasks)

	return mux
}

func configure(c *ServerConfig) {
	bindAddr, exists := os.LookupEnv("BIND_ADDR")
	if exists {
		if bindAddr == "" {
			log.Fatalln("(cfg) empty bind address supplied, cannot bind")
		}
		c.BindAddr = bindAddr
	} else {
		log.Printf("(cfg) no bind address supplied, defaulting to '%s'", c.BindAddr)
	}

	secrets, err := auth.GetSecretsFromEnv()
	if err != nil {
		log.Fatalln("(cfg) error in credential init:", err.Error())
	}

	config.OAuth2 = auth.GetOAuth2Config(secrets, auth.TC_API_SCOPES)
}

func main() {
	log.Println("Initializing config from environment variables...")

	configure(&config)

	log.Printf("Starting server binded to %s...", config.BindAddr)

	mux := makeMux()
	http.ListenAndServe(config.BindAddr, handlers.RequestLogger(mux))

	log.Println("Server exited. Cleaning up...")
}
