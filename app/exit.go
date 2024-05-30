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

func errorView(m Model) string {
	return fmt.Sprintf("\nThere was an error:\n%s\n\n", m.Err.Error())
}

func quitView(m Model) string {
	tmp := exitMainTitleStyle.Render("\nI am finally allowed to sleep!") + "\n"
	tmp += exitSecondTitleStyle.Render(fmt.Sprintf("(awake for %d sec)\n\n", m.TimeAwake))
	return tmp
}
