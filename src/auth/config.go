package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var TC_API_SCOPES = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/classroom.courses.readonly",
	"https://www.googleapis.com/auth/classroom.coursework.me.readonly",
	"https://www.googleapis.com/auth/classroom.courseworkmaterials.readonly",
	"https://www.googleapis.com/auth/classroom.guardianlinks.me.readonly",
	"https://www.googleapis.com/auth/classroom.push-notifications",
	"https://www.googleapis.com/auth/classroom.student-submissions.me.readonly",
	"https://www.googleapis.com/auth/classroom.announcements.readonly",
}

func GetOAuth2Config(secrets *OAuth2Secrets, scopes []string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     secrets.ClientID,
		ClientSecret: secrets.ClientSecret,
		RedirectURL:  "postmessage",
		Scopes:       TC_API_SCOPES,
		Endpoint:     google.Endpoint,
	}
}
