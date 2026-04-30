package main

import (
	"fmt"

	"github.com/google/go-github/v74/github"
)

var config appConfig

func main() {
	enterAltScreen()

	errLoadingConfig := startSpinner("Loading local config...", func() error {
		err := loadConfig()
		if err != nil {
			return err
		}
		return nil
	})
	if errLoadingConfig != nil {
		return
	}

	if config.Sound != "" {
		errInitAudio := startSpinner("Loading audio speaker...", func() error {
			err := initAudio()
			if err != nil {
				return err
			}
			return nil
		})
		if errInitAudio != nil {
			return
		}
	}

	errCheckingConfig := startSpinner("Loading user data...", func() error {
		_, err := checkPAT(config)
		return err
	})
	if errCheckingConfig != nil {
		return
	}

	var prs []*github.PullRequest
	errGettingPRs := startSpinner("Retrieving PR information...", func() error {
		var err error
		prs, err = getPRs(config.RepositoryList)
		if err != nil {
			fmt.Println(errorText("Error loading user details: "))
			return err
		}

		if len(prs) <= 0 {
			fmt.Println(errorText("Error loading user details: repository_list is empty"))
			return fmt.Errorf("error loading user details: repository_list is empty")
		}

		return nil
	})
	if errGettingPRs != nil {
		return
	}

	renderTable(prs)
}

// Clear and position the output at the top
func enterAltScreen() {
	fmt.Print("\033[H\033[2J")
}
