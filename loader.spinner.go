package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m modelSpinner) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg {
			err := m.functionSpinner()
			if err != nil {
				return doneMsg{err: err}
			} else {
				return doneMsg{err: nil}
			}
		},
	)
}

func (m modelSpinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case doneMsg:
		m.done = true
		m.err = msg.err
		return m, tea.Quit

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m modelSpinner) View() string {
	str := fmt.Sprintf("\n\n   %s %s\n\n", m.spinner.View(), spinnerTextStyle((m.text)))

	if m.err != nil {
		str = fmt.Sprintf("\n\n   %s %s %s\n\n", m.spinner.View(), spinnerTextStyle((m.text)), errorText("KO!"))
		return str
	}

	if m.done {
		str = fmt.Sprintf("\n\n   %s %s %s\n\n", m.spinner.View(), spinnerTextStyle((m.text)), successText("OK!"))
		return str
	}

	if m.quitting {
		return str + "\n"
	}
	return str
}

func startSpinner(spinnerText string, wrappedFunction func() error) error {
	loadingSpinner := spinner.New()
	loadingSpinner.Spinner = spinner.Globe

	modelLoadingSpinner := modelSpinner{
		spinner:         loadingSpinner,
		text:            spinnerText,
		functionSpinner: wrappedFunction,
	}

	p := tea.NewProgram(modelLoadingSpinner)

	finalModel, err := p.Run()
	if err != nil {
		return err
	}

	if finalModel, ok := finalModel.(modelSpinner); ok && finalModel.err != nil {
		return finalModel.err
	}

	return nil
}
