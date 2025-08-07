package main

import (
	"context"

	"github.com/google/go-github/v74/github"
)

func checkPAT(config appConfig) (*github.User, error) {
	ctx := context.Background()

	client := github.NewClient(nil).WithAuthToken(config.PAT)

	user, _, err := client.Users.Get(ctx, "")

	return user, err
}

func getPRs(repo_list []string) ([]*github.PullRequest, error) {
	var prs []*github.PullRequest

	ctx := context.Background()

	client := github.NewClient(nil).WithAuthToken(config.PAT)

	prOpts := github.PullRequestListOptions{
		State: "open",
		Sort:  "created",
	}

	for _, repo := range repo_list {
		repo_prs, _, err := client.PullRequests.List(ctx, config.Owner, repo, &prOpts)
		if err == nil {
			prs = append(prs, repo_prs...)
		}
	}

	return prs, err
}
