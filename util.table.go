package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/google/go-github/v74/github"
)

const (
	columnKeyURL        = "url"
	columnKeyRepository = "repository"
	columnKeyTitle      = "title"
	columnKeyCreatedBy  = "created_by"
	columnKeyStatus     = "status"
	columnKeyDraft      = "draft"
	columnKeyCreatedAt  = "created_at"
)

const (
	minTableWidth = 60
	widthSafety   = 1
	footerText    = "📡 GH PRadar | From Madrid with 🧡"
)

func openURL(targetURL string) tea.Cmd {
	return func() tea.Msg {
		parsed, err := url.Parse(targetURL)
		if err != nil || parsed.Scheme != "https" {
			return openURLErrorMsg{err: fmt.Errorf("invalid URL: %s", targetURL)}
		}
		cmd := exec.Command("xdg-open", targetURL)
		err = cmd.Start()
		return openURLErrorMsg{err}
	}
}

func newTableModel(rows []prTableRow) tableModel {
	initialWidth := 140
	t := table.New(buildColumns()).
		WithRows(toBubbleRows(rows)).
		Focused(true).
		WithBaseStyle(lipgloss.NewStyle().Align(lipgloss.Left)).
		HeaderStyle(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("117"))).
		HighlightStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("230")).Background(lipgloss.Color("62"))).
		WithStaticFooter(footerForWidth(initialWidth)).
		WithFooterVisibility(true).
		WithTargetWidth(initialWidth)

	return tableModel{table: t, rows: rows}
}

func (m tableModel) Init() tea.Cmd {
	return tea.Batch(
		fetchPRs(),
		scheduleTick(),
	)
}

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case refreshMsg:
		if config.Sound != "" && shouldPlayAudio(m.rows, msg.rows) {
			playAudio()
		}

		m.rows = msg.rows
		m.table = m.table.WithRows(toBubbleRows(msg.rows))
		return m, tea.Batch(cmds...)

	case tea.WindowSizeMsg:
		tableWidth := msg.Width - widthSafety
		if tableWidth < minTableWidth {
			tableWidth = minTableWidth
		}
		m.table = m.table.
			WithTargetWidth(tableWidth).
			WithStaticFooter(footerForWidth(tableWidth))
		return m, tea.Batch(cmds...)

	case tickMsg:
		cmds = append(cmds, fetchPRs(), scheduleTick())
		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			cmds = append(cmds, tea.Quit)
			return m, tea.Batch(cmds...)
		case "enter":
			highlighted := m.table.HighlightedRow()
			if highlighted.Data != nil {
				if url, ok := highlighted.Data[columnKeyURL].(string); ok && url != "" {
					cmds = append(cmds, openURL(url))
				}
			}
			return m, tea.Batch(cmds...)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m tableModel) View() string {
	return m.table.View()
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

func generateRows(prs []*github.PullRequest) []prTableRow {
	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		loc = time.UTC
	}
	rows := []prTableRow{}

	for _, v := range prs {
		row := prTableRow{
			URL:        v.GetHTMLURL(),
			Repository: v.GetHead().GetRepo().GetName(),
			Title:      v.GetTitle(),
			CreatedBy:  v.GetUser().GetLogin(),
			Status:     v.GetState(),
			Draft:      v.GetDraft(),
			CreatedAt:  v.GetCreatedAt().In(loc).Format("02 Jan 06 15:04"),
		}
		rows = append(rows, row)
	}

	prs = nil
	return rows
}

func shouldPlayAudio(currentRows, incomingRows []prTableRow) bool {
	currentRowsByURL := mapRowsByURL(currentRows)
	incomingRowsByURL := mapRowsByURL(incomingRows)

	for url, incoming := range incomingRowsByURL {
		if current, ok := currentRowsByURL[url]; !ok || current != incoming {
			if !incoming.Draft {
				return true
			}
		}
	}

	for url, current := range currentRowsByURL {
		if _, ok := incomingRowsByURL[url]; !ok {
			if !current.Draft {
				return true
			}
		}
	}

	return false
}

func mapRowsByURL(rows []prTableRow) map[string]prTableRow {
	rowsByURL := make(map[string]prTableRow, len(rows))
	for _, row := range rows {
		rowsByURL[row.URL] = row
	}
	return rowsByURL
}

func toBubbleRows(rows []prTableRow) []table.Row {
	bubbleRows := make([]table.Row, 0, len(rows))
	for _, row := range rows {
		statusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("78"))
		if row.Status != "open" {
			statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
		}

		draftText := fmt.Sprintf("%t", row.Draft)
		draftStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("78"))
		if row.Draft {
			draftStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
		}

		bubbleRows = append(bubbleRows, table.NewRow(table.RowData{
			columnKeyURL:        row.URL,
			columnKeyRepository: row.Repository,
			columnKeyTitle:      row.Title,
			columnKeyCreatedBy:  row.CreatedBy,
			columnKeyStatus:     table.NewStyledCell(row.Status, statusStyle),
			columnKeyDraft:      table.NewStyledCell(draftText, draftStyle),
			columnKeyCreatedAt:  row.CreatedAt,
		}))
	}
	return bubbleRows
}

func buildColumns() []table.Column {
	return []table.Column{
		table.NewColumn(columnKeyRepository, "Repository", 24),
		table.NewFlexColumn(columnKeyTitle, "Title", 4),
		table.NewFlexColumn(columnKeyCreatedBy, "Created by", 2),
		table.NewColumn(columnKeyStatus, "Status", 10),
		table.NewColumn(columnKeyDraft, "Draft", 7),
		table.NewColumn(columnKeyCreatedAt, "Created At", 16),
	}
}

func footerForWidth(tableWidth int) string {
	footerWidth := tableWidth - 2
	if footerWidth < 1 {
		footerWidth = 1
	}
	return lipgloss.NewStyle().Align(lipgloss.Right).Width(footerWidth).Render(footerText)
}

func renderTable(prs []*github.PullRequest) {
	m := newTableModel(generateRows(prs))
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
