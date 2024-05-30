package app

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	//  view
	window Window

	// model state
	Awake     bool
	TimeAwake int32
	Ctx       context.Context
	CancelCmd context.CancelFunc

	// choice state
	List         list.Model
	Choice       string
	TxtInput     textinput.Model
	TxtInputHelp string

	Err error
}

func NewModel() Model {
	m := Model{
		window:   HomeWindow,
		Awake:    false,
		Ctx:      context.Background(),
		List:     getHomeViewList(),
		Choice:   "",
		TxtInput: getChoiceInput(),
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if exitKey(msg) {
		m.SetWindow(QuitWindow)
		m.CancelCmd()
		return m, tea.Quit
	}

	switch m.GetWindow() {
	case HomeWindow:
		return homeUpdate(msg, m)
	case AwakeWindow:
		return awakeUpdate(msg, m)
	default:
		return homeUpdate(msg, m)
	}
}

func (m Model) View() string {
	switch m.GetWindow() {
	case HomeWindow:
		return homeView(m)
	case AwakeWindow:
		return awakeView(m)
	case ErrorWindow:
		return errorView(m)
	case QuitWindow:
		return quitView(m)
	default:
		return homeView(m)
	}
}

// ── style ───────────────────────────────────────────────────────────

// ── window ──────────────────────────────────────────────────────────

type Window int

const (
	QuitWindow Window = iota
	HomeWindow
	AwakeWindow
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

// ── utils ───────────────────────────────────────────────────────────

func exitKey(msg tea.Msg) bool {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "ctrl+c" {
			return true
		}
	}
	return false
}
