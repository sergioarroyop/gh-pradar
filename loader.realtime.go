package main

// A simple example that shows how to send activity to Bubble Tea in real-time
// through a channel.

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var spinnerTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render

// A message used to indicate that activity has occurred.
type responseMsg struct {
	err error
}

// Listen for activity.
func listenForActivity(sub chan tea.Msg, function func() error) tea.Cmd {
	return func() tea.Msg {
		// Execute function
		err := function()
		if err != nil {
			sub <- responseMsg{err: err}
		} else {
			sub <- responseMsg{err: nil}
		}

		return sub
	}
}

// A command that waits for the activity on a channel.
func waitForActivity(sub chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

type modelSpinner struct {
	sub         chan tea.Msg // where we'll receive activity notifications
	spinner     spinner.Model
	spinnerText string
	function    func() error
	err         error
}

func (m modelSpinner) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		listenForActivity(m.sub, m.function), // generate activity
		waitForActivity(m.sub),               // wait for activity
	)
}

func (m modelSpinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			quitting = true
			return m, tea.Quit
		}
	case responseMsg:
		if msg.err != nil {
			m.err = msg.err
		}
		return m, tea.Quit // exit
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
	return m, nil
}

func (m modelSpinner) View() string {
	s := fmt.Sprintf("\n %s%s%s\n\n", m.spinner.View(), "  ", spinnerTextStyle(m.spinnerText))

	return s
}

func startRealtimeLoader(spinnerText string, function func() error) error {
	loadingSpinner := spinner.New()
	loadingSpinner.Spinner = spinner.Globe

	spinnerModel := modelSpinner{
		sub:         make(chan tea.Msg),
		spinner:     loadingSpinner,
		spinnerText: spinnerText,
		function:    function,
	}

	enterAltScreen()
	p := tea.NewProgram(spinnerModel)

	finalModel, err := p.Run()
	if err != nil {
		return err
	}
	// Check if the final model has an error
	if finalModel, ok := finalModel.(modelSpinner); ok && finalModel.err != nil {
		return finalModel.err
	}
	return nil
}
