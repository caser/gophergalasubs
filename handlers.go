package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
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

func handleLogout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "token", Value: "", Expires: time.Now().Add(-time.Minute)}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func handleRepos(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if len(token) == 0 {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		repos, err := GetRepos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(repos)
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

func handleUserUpdate(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if len(token) == 0 {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		user, err := authenticateUser(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		dummyUser := User{}
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&dummyUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		user.Vote1 = dummyUser.Vote1
		user.Vote2 = dummyUser.Vote2
		user.Vote3 = dummyUser.Vote3
		user.Vote4 = dummyUser.Vote4
		user.Vote5 = dummyUser.Vote5

		err = user.Save()
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

func handleVoteCreate(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	vars := mux.Vars(r)
	owner := vars["owner"]
	name := vars["name"]

	if len(token) == 0 {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		user, err := authenticateUser(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(oauth2.NoContext, ts)

		client := github.NewClient(tc)

		log.Printf("OWNER: %s, NAME: %s", owner, name)

		repo, _, err := client.Repositories.Get(owner, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = user.Vote(*repo.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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

func authenticateUser(token string) (*User, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	user, _, err := client.Users.Get("")
	if err != nil {
		return nil, err
	}

	u, err := UpsertUserFromGithubUser(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}
