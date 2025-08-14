package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
)

// api.config types
type appConfig struct {
	PAT            string   `json:"personal_access_token"`
	Sound          string   `json:"sound"`
	RepositoryList []string `json:"repository_list"`
}

// loader.spinner types
type doneMsg struct {
	err error
}

type openURLErrorMsg struct {
	err error
}

type modelSpinner struct {
	spinner         spinner.Model
	functionSpinner func() error
	text            string
	quitting        bool
	done            bool

	err error
}

// util.table types
type (
	tickMsg    struct{}
	refreshMsg struct {
		rows []table.Row
		err  error
	}
	tableModel struct {
		table table.Model
	}
)
