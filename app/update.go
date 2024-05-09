package app

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if exitKey(msg) {
		m.Quitting = true
		return m, tea.Quit
	}

	switch m.GetWindow() {
	case HomeWindow:
		return homeUpdate(msg, m)
	case OnWindow:
		return onUpdate(msg, m)
	case ErrorWindow:
		return errorUpdate(msg, m)
	default:
		return homeUpdate(msg, m)
	}
}

func exitKey(msg tea.Msg) bool {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return true
		}
	}
	return false
}

func homeUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	return m, nil
}

func onUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	return m, nil
}

func errorUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	return m, nil
}
