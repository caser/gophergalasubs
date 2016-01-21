package main

import (
	"log"

	"github.com/google/go-github/github"
)

type User struct {
	Id       *int
	GithubId *int
	Email    *string
	Login    *string
}

func UpsertUserFromGithubUser(u *github.User) (*User, error) {
	//TODO this should use a Postgres upsert
	var id, github_id *int
	var login, email *string

	rows, err := db.Query("SELECT id, github_id, login, email FROM users WHERE github_id = $1", u.ID)
	if err != nil {
		log.Println("HERE 2")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &github_id, &login, &email)
		if err != nil {
			log.Println("HERE 2")
			log.Printf("%#v", err)
			return nil, err
		}
		return &User{
			Id:       id,
			GithubId: github_id,
			Login:    login,
			Email:    email,
		}, nil
	}
	return CreateUserFromGithubUser(u)
}

func CreateUserFromGithubUser(u *github.User) (*User, error) {
	var lastInsertId *int
	err := db.QueryRow(
		"INSERT INTO users(email, login, github_id) VALUES($1, $2, $3) RETURNING id",
		u.Email,
		u.Login,
		u.ID,
	).Scan(lastInsertId)
	if err != nil {
		log.Println("HERE 5")
		log.Printf("%#v", err)
		return nil, err
	}
	return &User{
		Id:       lastInsertId,
		GithubId: u.ID,
		Login:    u.Login,
		Email:    u.Email,
	}, nil
}
