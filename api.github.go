package main

import (
	"context"

	"github.com/google/go-github/v74/github"
)

func checkPAT(config appConfig) (*github.User, *github.Response, error) {
	ctx := context.Background()

	client := github.NewClient(nil).WithAuthToken(config.PAT)

	user, resp, err := client.Users.Get(ctx, "")

	return user, resp, err
}

func getPRs(config appConfig) ([]*github.PullRequest, *github.Response, error) {
	ctx := context.Background()

	client := github.NewClient(nil).WithAuthToken(config.PAT)

	prOpts := github.PullRequestListOptions{
		State: "open",
		Sort:  "created",
	}

	prs, resp, err := client.PullRequests.List(ctx, config.Owner, "cinexin-downloader", &prOpts)

	return prs, resp, err
}
