package main

import (
	"log"
	"sync"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var allRepos []Repo
var mut = new(sync.RWMutex)
var lastCached = time.Now().Add(-40 * time.Minute)

type Repo struct {
	ID              *int    `json:"id"`
	Name            *string `json:"name"`
	FullName        *string `json:"full_name"`
	HTMLURL         *string `json:"html_url"`
	StargazersCount *int    `json:"stargazers_count"`
	Description     *string `json:"description"`
}

func RepoCacheIsInDate() bool {
	cacheLife := time.Duration(20 * time.Minute)
	delta := time.Since(lastCached)
	return delta > cacheLife
}

func GetRepos() ([]Repo, error) {
	mut.RLock()
	if len(allRepos) > 0 && RepoCacheIsInDate() {
		log.Print("Using cache")
		res := allRepos
		mut.RUnlock()
		return res, nil
	}
	mut.RUnlock()

	mut.Lock()
	defer mut.Unlock()

	tc := oauth2.NewClient(oauth2.NoContext, nil)

	client := github.NewClient(tc)

	githubRepos := make([]github.Repository, 0)

	// list all repositories for the authenticated user
	opt := &github.RepositoryListByOrgOptions{
		Type:        "public",
		ListOptions: github.ListOptions{PerPage: 50, Page: 1},
	}

	// get all pages of results
	for {
		repos, resp, err := client.Repositories.ListByOrg("gophergala", opt)
		if err != nil {
			allRepos = make([]Repo, 0)
			return allRepos, err
		}
		githubRepos = append(githubRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}

	for _, githubRepo := range githubRepos {
		allRepos = append(allRepos, Repo{
			ID:              githubRepo.ID,
			Name:            githubRepo.Name,
			FullName:        githubRepo.FullName,
			HTMLURL:         githubRepo.HTMLURL,
			StargazersCount: githubRepo.StargazersCount,
			Description:     githubRepo.Description,
		})

	}
	return allRepos, nil
}
