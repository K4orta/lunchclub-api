package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var fbClientID = os.Getenv("LCAPI_FB_APP_ID")
var fbClientSecret = os.Getenv("LCAPI_FB_SECRET")
var sessionSecret = os.Getenv("LCAPI_SESSION_SECRET")
var store = sessions.NewCookieStore([]byte(sessionSecret))

var conf = &oauth2.Config{
	ClientID:     fbClientID,
	ClientSecret: fbClientSecret,
	Scopes:       []string{"public_profile", "email"},
	Endpoint:     facebook.Endpoint,
}

// AccessToken is a stuct that holds a token value and an extires time.
type AccessToken struct {
	Token   string
	Expires int64
}

// GetFBToken is a function that exchanges a code for an access token.
func GetFBToken(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	tok, err := conf.Exchange(oauth2.NoContext, q.Get("code"))
	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + tok.AccessToken)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	session, _ := store.Get(req, "lc-session")
	session.Values["uname"] = "Adam"
	session.Values["pw"] = "1234"
	session.Save(req, w)

	fmt.Fprint(w, string(respBody))
}

// RedirectFBLogin forwards the user to the Facebook login page.
func RedirectFBLogin(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, FBLoginURL("http://localhost:8001/auth"), http.StatusTemporaryRedirect)
}

// ConfirmSession forwards the user to the Facebook login page.
func ConfirmSession(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "lc-session")
	fmt.Fprint(w, session.Values["uname"])
}

// FBLoginURL returns a URL for the FB login promt.
func FBLoginURL(redirectURL string) string {
	conf.RedirectURL = redirectURL
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	return url
}
