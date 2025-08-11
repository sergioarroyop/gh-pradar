package main

import "github.com/charmbracelet/bubbles/spinner"

type appConfig struct {
	PAT            string   `json:"personal_access_token"`
	Sound          bool     `json:"sound"`
	RepositoryList []string `json:"repository_list"`
}

type doneMsg struct {
	err error
}

type modelSpinner struct {
	spinner         spinner.Model
	text            string
	quitting        bool
	done            bool
	err             error
	functionSpinner func() error
}
