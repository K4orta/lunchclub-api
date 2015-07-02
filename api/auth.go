package api

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var fbClientID = os.Getenv("LCAPI_FB_APP_ID")
var fbClientSecret = os.Getenv("LCAPI_FB_SECRET")

// AccessToken is a stuct that holds a token value and an extires time
type AccessToken struct {
	Token   string
	Expires int64
}

// GetFBToken is a function that exchanges a code for an access token
func GetFBToken(w http.ResponseWriter, req *http.Request) {

}

func RedirectFBLogin(w http.ResponseWriter, req *http.Request) {

}

// FBLoginURL returns a URL for the FB login promt
func FBLoginURL(redirectURL string) string {
	conf := &oauth2.Config{
		ClientID:     fbClientID,
		ClientSecret: fbClientSecret,
		Scopes:       []string{"public_profile", "email"},
		RedirectURL:  redirectURL,
		Endpoint:     facebook.Endpoint,
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	return url
}
