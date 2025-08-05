package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v74/github"
)

func checkPAT(config appConfig) (*github.User, *github.Response) {
	ctx := context.Background()

	client := github.NewClient(nil).WithAuthToken(config.PAT)

	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Println(errorText("Error fetching user: " + err.Error()))
	}

	return user, resp
}
