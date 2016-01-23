package main

import (
	"sync"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var allRepos []github.Repository
var mut = new(sync.RWMutex)
var lastCached = time.Now().Add(-40 * time.Minute)

func RepoCacheIsInDate() bool {
	cacheLife := time.Duration(20 * time.Minute)
	delta := time.Since(lastCached)
	return delta > cacheLife
}

func GetRepos() ([]github.Repository, error) {
	mut.RLock()
	if len(allRepos) > 0 && RepoCacheIsInDate() {
		res := allRepos
		mut.RUnlock()
		return res, nil
	}
	mut.RUnlock()

	mut.Lock()
	defer mut.Unlock()

	tc := oauth2.NewClient(oauth2.NoContext, nil)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	opt := &github.RepositoryListByOrgOptions{
		Type:        "public",
		ListOptions: github.ListOptions{PerPage: 50, Page: 1},
	}

	// get all pages of results
	for {
		repos, resp, err := client.Repositories.ListByOrg("gophergala", opt)
		if err != nil {
			allRepos = make([]github.Repository, 0)
			return allRepos, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
	return allRepos, nil
}
