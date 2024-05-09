package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// current view
	window Window

	Awake bool
	Err   error
}

func NewModel() Model {
	m := Model{window: HomeWindow}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

type Window int

const (
	QuitWindow Window = iota
	HomeWindow
	OnWindow
	ErrorWindow
	_endWindowIota
)

func (m *Model) SetWindow(newWindow Window) {
	if newWindow >= _endWindowIota { // is window value valid
		newWindow = ErrorWindow
	}
	m.window = newWindow
}

func (m *Model) GetWindow() Window {
	return m.window
}
