package app

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Choice int

	Window Window

	Awake    bool
	Quitting bool
}

func NewModel() Model {
	m := Model{
		Choice: 0,
		Window: Home,
		Awake:  false,
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
