package main

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v74/github"
)

const (
	configPath = "/.config/gh-pradar/"
	configFile = "config.json"
)

var (
	config   appConfig
	err      error
	quitting bool = false
)

func main() {
	enterAltScreen()

	err = loadConfig(configPath, configFile)

	startSpinner("Loading user data...", func() error {
		_, err := checkPAT(config)
		if err != nil {
			fmt.Println(errorText("Error loading user details: " + err.Error()))
		}
		return err
	})
	if err != nil {
		return
	}

	var prs []*github.PullRequest
	startSpinner("Retrieving PR information...", func() error {
		prs, err = getPRs(config.ReposityList)
		if err != nil {
			fmt.Println(errorText("Error loading user details: " + err.Error()))
		}

		return err
	})

	for _, v := range prs {
		fmt.Println(hyperlinkText(*v.HTMLURL, *v.Title))
	}

	renderTable(prs)

	if quitting {
		print(warnText("Exiting module loading..."))
		return
	}
}

// Function to check for the `q` key in any input and exit the program
func checkForQuit(input string) bool {
	if strings.ToLower(input) == "q" {
		quitting = true
		print(warnText("Exiting program..."))
		return true
	}
	return false
}

// Clear and position the output at the top
func enterAltScreen() {
	fmt.Print("\033[H\033[2J")
}
