package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/google/go-github/v74/github"
)

func checkPAT(config appConfig) (*github.User, error) {
	ctx := context.Background()

	client := github.NewClient(nil).WithAuthToken(config.PAT)

	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Println(errorText("Error loading user details: " + err.Error()))
	}
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
		repo := strings.Split(repo, "/")
		repo_prs, _, err := client.PullRequests.List(ctx, repo[0], repo[1], &prOpts)

		if err == nil {
			prs = append(prs, repo_prs...)
		}
	}

	sort.Slice(prs, func(i, j int) bool {
		return prs[i].GetCreatedAt().After(prs[j].GetCreatedAt().Time)
	})

	return prs, err
}
