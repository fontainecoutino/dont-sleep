package app

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// ── view ────────────────────────────────────────────────────────────

var (
	exitMainTitleStyle   = lipgloss.NewStyle().Bold(true)
	exitSecondTitleStyle = lipgloss.NewStyle().Italic(true)
)

func errorView() string {
	return "\nThere was an error\n\n"
}

func quitView(m Model) string {
	tmp := exitMainTitleStyle.Render("\nI am finally allowed to sleep!") + "\n"
	tmp += exitSecondTitleStyle.Render(fmt.Sprintf("(PID: %d)\n\n", PID))
	return tmp
}
