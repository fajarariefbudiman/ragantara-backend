package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type Config struct {
	Google_Client_Id       string
	Google_Client_Secret   string
	Facebook_Client_Id     string
	Facebook_Client_Secret string
}

var config Config
var GoogleOauthConfig *oauth2.Config
var FacebookOauthConfig *oauth2.Config

func Init() {

	config = Config{
		Google_Client_Id:       "843631712111-gp4gju9atodttmdf8bf4gsdg22adc8sf.apps.googleusercontent.com",
		Google_Client_Secret:   "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",
		Facebook_Client_Id:     os.Getenv("FACEBOOK_CLIENT_ID"),
		Facebook_Client_Secret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
	}
	// fmt.Println("Google Client ID:", config.Google_Client_Id)

	GoogleOauthConfig = &oauth2.Config{
		ClientID:     config.Google_Client_Id,
		ClientSecret: config.Google_Client_Secret,
		RedirectURL:  "http://localhost:1323/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	FacebookOauthConfig = &oauth2.Config{
		ClientID:     config.Facebook_Client_Id,
		ClientSecret: config.Facebook_Client_Secret,
		RedirectURL:  "http://localhost:1323/auth/facebook/callback",
		Scopes: []string{
			"email",
			"public_profile",
		},
		Endpoint: facebook.Endpoint,
	}
}
