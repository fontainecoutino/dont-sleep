package app

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// ── update ──────────────────────────────────────────────────────────

func awakeUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case caffeinateFinishedMsg:
		if msg.err != nil {
			m.Err = msg.err
			m.SetWindow(ErrorWindow)
			return m, tea.Quit
		}
		m.SetWindow(QuitWindow)
		return m, tea.Quit
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case startCaffeinateMsg:
		m.Awake = true
		return m, caffeinate()
	}
	return m, nil
}

// ── view ────────────────────────────────────────────────────────────

func awakeView(m Model) string {
	return fmt.Sprintf("\n  chose: %s\n  %s\n\n", m.Choice, m.spinner.View())
}

// ── cmd ─────────────────────────────────────────────────────────────

type caffeinateFinishedMsg struct{ err error }

func caffeinate() tea.Cmd {
	return func() tea.Msg {
		// build command
		awakeCmd := exec.Command("caffeinate", "-dt 5")
		awakeCmd.Stdout = nil

		err := awakeCmd.Start()
		if err != nil {
			return caffeinateFinishedMsg{err: errors.New("error while starting caffeinate")}
		}

		err = awakeCmd.Wait()
		if err != nil {
			_, ok := err.(*exec.ExitError)
			if !ok {
				return caffeinateFinishedMsg{err: errors.New("caffeinate error")}
			}
			return caffeinateFinishedMsg{err: errors.New("error while executing caffeinate")}
		}
		return caffeinateFinishedMsg{}
	}
}
