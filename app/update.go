package app

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// quitting keys
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	switch m.Window {
	case Home:
		return homeUpdate(msg, m)
	case On:
		return onUpdate(msg, m)
	default:
		return homeUpdate(msg, m)
	}
}

func homeUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	return nil, nil
}

func onUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	return nil, nil
}
