package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/go-github/v74/github"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

const (
	urlCol = iota
	draftCol
)

func openURL(url string) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("xdg-open", url)
		err := cmd.Start()
		return openURLErrorMsg{err}
	}
}

func (m tableModel) Init() tea.Cmd {
	return tea.Batch(
		fetchPRs(),
		scheduleTick(),
	)
}

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case refreshMsg:
		if config.Sound != "" {
			if shouldPlayAudio(m.table.Rows(), msg.rows) {
				playAudio()
			}
		}
		m.table.SetRows(msg.rows)
		return m, nil
	case tickMsg:
		return m, tea.Batch(
			fetchPRs(),
			scheduleTick(),
		)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				openURL(m.table.SelectedRow()[0]),
			)
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func scheduleTick() tea.Cmd {
	return tea.Tick(time.Minute, func(time.Time) tea.Msg { return tickMsg{} })
}

func fetchPRs() tea.Cmd {
	return func() tea.Msg {
		prs, err := getPRs(config.RepositoryList)

		rows := generateRows(prs)

		return refreshMsg{rows, err}
	}
}

func generateRows(prs []*github.PullRequest) []table.Row {
	loc, _ := time.LoadLocation("Europe/Madrid")
	rows := []table.Row{}

	for _, v := range prs {
		row := table.Row{
			v.GetHTMLURL(),
			v.GetHead().GetRepo().GetName(),
			v.GetTitle(),
			v.GetUser().GetLogin(),
			v.GetState(),
			fmt.Sprintf("%t", v.GetDraft()),
			v.GetCreatedAt().In(loc).Format("02 Jan 06 15:04"),
		}
		rows = append(rows, row)
	}

	prs = nil

	return rows
}

func shouldPlayAudio(currentRows, incomingRows []table.Row) bool {
	currentRowsByURL := mapRowsByURL(currentRows)
	incomingRowsByURL := mapRowsByURL(incomingRows)

	for url, incoming := range incomingRowsByURL {
		if current, ok := currentRowsByURL[url]; !ok || !slices.Equal(current, incoming) {
			if !isDraftRow(incoming) {
				return true
			}
		}
	}

	for url, current := range currentRowsByURL {
		if _, ok := incomingRowsByURL[url]; !ok {
			if !isDraftRow(current) {
				return true
			}
		}
	}

	return false
}

func mapRowsByURL(rows []table.Row) map[string]table.Row {
	rowsByURL := make(map[string]table.Row, len(rows))
	for _, row := range rows {
		if len(row) <= urlCol {
			continue
		}
		rowsByURL[row[urlCol]] = row
	}
	return rowsByURL
}

func isDraftRow(row table.Row) bool {
	if len(row) <= draftCol {
		return false
	}
	return row[draftCol] == "true"
}

func renderTable(prs []*github.PullRequest) {
	columns := []table.Column{
		{Title: "URL", Width: 0},
		{Title: "Repository", Width: 20},
		{Title: "Title", Width: 50},
		{Title: "Created by", Width: 20},
		{Title: "Status", Width: 20},
		{Title: "Draft", Width: 8},
		{Title: "Created At", Width: 20},
	}

	rows := generateRows(prs)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := tableModel{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
