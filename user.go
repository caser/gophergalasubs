package main

import (
	"errors"
	"log"

	"github.com/google/go-github/github"
)

type User struct {
	Id        *int    `json:"id,omitempty"`
	GithubId  *int    `json:"github_id,omitempty"`
	Email     *string `json:"email,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Login     *string `json:"login,omitempty"`
	Vote1     *int    `json:"vote1,omitempty"`
	Vote2     *int    `json:"vote2,omitempty"`
	Vote3     *int    `json:"vote3,omitempty"`
	Vote4     *int    `json:"vote4,omitempty"`
	Vote5     *int    `json:"vote5,omitempty"`
}

func UpsertUserFromGithubUser(u *github.User) (*User, error) {
	//TODO this should use a Postgres upsert
	user, err := GetUserByGithubId(u.ID)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	return CreateUserFromGithubUser(u)
}

func GetUserByGithubId(github_id *int) (*User, error) {
	var id, vote1, vote2, vote3, vote4, vote5 *int
	var login, email, avatar_url *string

	rows, err := db.Query("SELECT id, login, email, avatar_url, vote1, vote2, vote3, vote4, vote5 FROM users WHERE github_id = $1", github_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &login, &email, &avatar_url, &vote1, &vote2, &vote3, &vote4, &vote5)
		if err != nil {
			return nil, err
		}
		return &User{
			Id:        id,
			GithubId:  github_id,
			Login:     login,
			Email:     email,
			AvatarUrl: avatar_url,
			Vote1:     vote1,
			Vote2:     vote2,
			Vote3:     vote3,
			Vote4:     vote4,
			Vote5:     vote5,
		}, nil
	}
	return nil, nil
}

func CreateUserFromGithubUser(u *github.User) (*User, error) {
	var lastInsertId *int
	err := db.QueryRow(
		"INSERT INTO users(email, login, github_id, avatar_url) VALUES($1, $2, $3, $4) RETURNING id",
		u.Email,
		u.Login,
		u.ID,
		u.AvatarURL,
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

func (u *User) Vote(id int) error {
	if u.Vote5 != nil {
		return errors.New("5 votes already submitted")
	}
	if u.Vote1 == nil {
		u.Vote1 = &id
		return u.Save()
	} else if *u.Vote1 == id {
		return errors.New("Duplicate vote")
	}
	if u.Vote2 == nil {
		u.Vote2 = &id
		return u.Save()
	} else if *u.Vote2 == id {
		return errors.New("Duplicate vote")
	}
	if u.Vote3 == nil {
		u.Vote3 = &id
		return u.Save()
	} else if *u.Vote3 == id {
		return errors.New("Duplicate vote")
	}
	if u.Vote4 == nil {
		u.Vote4 = &id
		return u.Save()
	} else if *u.Vote4 == id {
		return errors.New("Duplicate vote")
	}
	if u.Vote5 == nil {
		u.Vote5 = &id
		return u.Save()
	} else if *u.Vote5 == id {
		return errors.New("Duplicate vote")
	}
	return nil
}

func (u *User) Save() error {
	_, err := db.Query(
		"UPDATE users SET (email, login, github_id, vote1, vote2, vote3, vote4, vote5) = ($1, $2, $3, $4, $5, $6, $7, $8) WHERE id = $9",
		u.Email,
		u.Login,
		u.GithubId,
		u.Vote1,
		u.Vote2,
		u.Vote3,
		u.Vote4,
		u.Vote5,
		u.Id,
	)
	return err
}
