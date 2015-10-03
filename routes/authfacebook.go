package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/oauth2"

	"github.com/julienschmidt/httprouter"
)

var oauthConfig = &oauth2.Config{ //setup
	ClientID:     "910962655664722",
	ClientSecret: "cdd845180e7223f34f4e2617360414cf",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://www.facebook.com/dialog/oauth",
		TokenURL: "https://graph.facebook.com/oauth/access_token",
	},
	RedirectURL: "http://localhost:8080/auth/facebook/callback",
}

// AuthFacebook is the /auth/facebook route
func AuthFacebook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title string
	}{
		"tvt.io",
	}
	t.Execute(w, context)
}

// AuthFacebookCallback is the /auth/facebook/callback route
func AuthFacebookCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	code := r.FormValue("code")
	token, err := oauthConfig.Exchange(oauth2.NoContext, code)

	client := oauthConfig.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://graph.facebook.com/v2.4/me")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	// Store session
	fmt.Println(string(body))

	// TODO Store credentials
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}