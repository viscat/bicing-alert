package app

import (
	"io/ioutil"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"context"
	"path/filepath"
)

type Credentials struct {
	Cid string `json:"clientId"`
	Csecret string `json:"secret"`
}

type GoogleUser struct {
	Sub string `json:"sub"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Profile string `json:"profile"`
	Picture string `json:"picture"`
	Email string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Gender string `json:"gender"`
}

type GoogleAuth struct {
	credentials Credentials
	conf *oauth2.Config
	state string
}


func NewGoogleAuth(redirectUrl string) GoogleAuth{

	credentials := loadCredentials()
	conf := &oauth2.Config{
		ClientID:     credentials.Cid,
		ClientSecret: credentials.Csecret,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}

	return GoogleAuth{credentials: credentials, conf: conf}
}

func loadCredentials() Credentials {
	var credentials Credentials
	path, err := filepath.Abs("cred.json")
	if err != nil {
		panic(err)
	}
	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(fileContents, &credentials)

	return credentials
}

func (g GoogleAuth) randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func (g GoogleAuth) getLoginURL(state string) string {
	return g.conf.AuthCodeURL(state)
}

func (g GoogleAuth) getState(r *http.Request) string {
	cookie, err := r.Cookie("state")
	if err != nil {
		return ""
	}

	return cookie.Value
}

func (g GoogleAuth) setState(w http.ResponseWriter, state string) {
	http.SetCookie(w, &http.Cookie{Name: "state", Value: state})
}



func (g GoogleAuth) AuthHandler(w http.ResponseWriter, r *http.Request) {

	retrievedState := g.getState(r)

	if retrievedState != r.URL.Query().Get("state") {
		Unauthorized(w)
		return
	}

	tok, err := g.conf.Exchange(context.Background(), r.URL.Query().Get("code"))
	if err != nil {
		BadRequest(w)
		return
	}

	client := g.conf.Client(context.Background(), tok)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		BadRequest(w)
		return
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	googleUser := &GoogleUser{}
	json.Unmarshal(data, googleUser)

	addUserSession(w, *googleUser)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (g GoogleAuth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	state := g.randToken()
	g.setState(w, state)

	w.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + g.getLoginURL(state) + "'><button>Login with Google!</button> </a> </body></html>"))
}
