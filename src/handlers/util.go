package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/buger/jsonparser"
	"golang.org/x/oauth2"
)

type BaseHandler struct {
	Config *oauth2.Config
}

func NewBaseHandler(config *oauth2.Config) *BaseHandler {
	return &BaseHandler{
		Config: config,
	}
}

// make sure that the method of the request is what we expect
func EnsureMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed, expected " + method))
		return false
	}
	return true
}

func RequestLogger(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		mux.ServeHTTP(w, r)

		log.Printf(
			"%s %s from [ %s ] done in [ %v ]",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

func ReadToken(body []byte) (*oauth2.Token, error) {
	/*
		request format:

		{
			"token": {
				"access": "accesstoken123"
				"refresh": "refreshtoken456"
				"expires": 12345678
			}
		}
	*/

	// get stuff from json
	access, err := jsonparser.GetString(body, "token", "access")
	if err != nil {
		return nil, err
	}

	refresh, err := jsonparser.GetString(body, "token", "refresh")
	if err != nil {
		return nil, err
	}

	timestamp, err := jsonparser.GetInt(body, "token", "expires")
	if err != nil {
		return nil, err
	}

	// parse time
	expiry := time.Unix(timestamp, 0).UTC()

	// construct token struct, get pointer
	token := &oauth2.Token{
		AccessToken:  access,
		RefreshToken: refresh,
		Expiry:       expiry,
	}

	// hand pointer back
	return token, nil
}
