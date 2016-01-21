package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	_ "github.com/wader/disable_sendfile_vbox_linux"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	// You must register the app at https://github.com/settings/applications
	// Set callback to http://127.0.0.1:7000/github_oauth_cb
	// Set ClientId and ClientSecret to
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		// select level of access you want https://developer.github.com/v3/oauth/#scopes
		Scopes:   []string{"user:email", "repo"},
		Endpoint: githuboauth.Endpoint,
	}
	// TODO: random string for oauth2 API calls to protect against CSRF
	oauthStateString = "thisshouldberandomx"
	db               *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@postgres/development?sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/login", handleGitHubLogin)
	r.HandleFunc("/github_oauth_cb", handleGitHubCallback)
	r.HandleFunc("/repos", handleRepos)
	r.HandleFunc("/user", handleUser)
	r.HandleFunc("/vote/{vote_id}", handleVote)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend")))
	http.Handle("/", r)

	fmt.Println("Started running on http://127.0.0.1:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
