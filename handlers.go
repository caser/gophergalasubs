package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	//just some test code
	_, err := r.Cookie("token")
	if err == nil {
		http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
		return
	}

	t := template.New("index.html")
	t, err = t.ParseFiles("./static/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// /login
func handleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// /github_oauth_cb. Called by github after authorization is granted
func handleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Printf("token is: %#v\n", token)

	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "token", Value: token.AccessToken, Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	var repos []github.Repository
	var err error

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: cookie.Value},
		)
		tc := oauth2.NewClient(oauth2.NoContext, ts)

		client := github.NewClient(tc)

		// list all repositories for the authenticated user
		opt := &github.RepositoryListByOrgOptions{
			Type:        "public",
			ListOptions: github.ListOptions{PerPage: 20, Page: 1},
		}
		repos, _, err = client.Repositories.ListByOrg("gophergala", opt)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	t := template.New("dashboard.html")
	t, err = t.ParseFiles("./static/html/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = t.Execute(w, repos)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func handleLeaderboard(w http.ResponseWriter, r *http.Request) {
	t := template.New("leaderboard.html")
	t, err := t.ParseFiles("./static/html/leaderboard.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
