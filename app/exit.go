package app

import (
	"fmt"
)

// ── view ────────────────────────────────────────────────────────────

func errorView(m Model) string {
	return fmt.Sprintf("\nThere was an error:\n%s\n\n", m.Err.Error())
}

func quitView(m Model) string {
	// TODO: add time elapsed
	tmp := "\n  I am finally allowed to sleep!\n\n"

	return tmp
}
