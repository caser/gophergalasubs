package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	_ "github.com/wader/disable_sendfile_vbox_linux"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user:email"},
		Endpoint:     githuboauth.Endpoint,
	}
	// TODO: random string for oauth2 API calls to protect against CSRF
	oauthStateString = "thisshouldberandomx"
	db               *sql.DB
)

func ensureDBTables() {
	time.Sleep(5 * time.Second)
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY NOT NULL, github_id INT, login TEXT, email TEXT, avatar_url TEXT, vote1 INT, vote2 INT, vote3 INT, vote4 INT, vote5 INT);")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	ensureDBTables()

	r := mux.NewRouter()
	r.HandleFunc("/login", handleGitHubLogin)
	r.HandleFunc("/github_oauth_cb", handleGitHubCallback)
	r.HandleFunc("/logout", handleLogout)
	r.HandleFunc("/repos", handleRepos).Methods("GET")
	r.HandleFunc("/user", handleUser).Methods("GET")
	r.HandleFunc("/user", handleUserUpdate).Methods("PATCH")
	r.HandleFunc("/vote/{owner}/{name}", handleVoteCreate).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	http.Handle("/", r)

	fmt.Println("Started running on http://127.0.0.1:8080")
	fmt.Println(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
