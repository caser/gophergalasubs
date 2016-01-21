package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/index.html")
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

func handleRepos(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if len(token) == 0 {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		var allRepos []github.Repository
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(oauth2.NoContext, ts)

		client := github.NewClient(tc)

		// list all repositories for the authenticated user
		opt := &github.RepositoryListByOrgOptions{
			Type:        "public",
			ListOptions: github.ListOptions{PerPage: 50, Page: 1},
		}

		// get all pages of results
		for {
			repos, resp, err := client.Repositories.ListByOrg("github", opt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			allRepos = append(allRepos, repos...)
			if resp.NextPage == 0 {
				break
			}
			opt.ListOptions.Page = resp.NextPage
		}

		js, err := json.Marshal(allRepos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if len(token) == 0 {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		user, err := authenticateUser(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func authenticateUser(token string) (*github.User, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	user, _, err := client.Users.Get("")
	if err != nil {
		return nil, err
	}

	_, err = UpsertUserFromGithubUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
