package app

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ── update ──────────────────────────────────────────────────────────

func homeUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" {
			m.Choice = ""
			m.TxtInputHelp = ""
			return m, nil
		}
	}

	switch m.Choice {
	case "":
		return choosingUpdate(msg, m)
	case TimeChoice:
		return choosenUpdate(msg, m, inputToTime)
	case ProcessChoice:
		return choosenUpdate(msg, m, inputToProcess)
	}

	return m, nil
}

// ── sub-update ──────────────────────────────────────────────────────

func choosingUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			i, ok := m.List.SelectedItem().(item)
			if ok {
				m.Choice = string(i)
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func choosenUpdate(msg tea.Msg, m Model, inputToVal func(string) (int, error)) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			input, err := inputToVal(m.TxtInput.Value())
			if err != nil {
				m.TxtInputHelp = err.Error()
				return m, nil
			}

			m.SetWindow(AwakeWindow)
			return m, startCaffeinate(m, input)
		}
		m.TxtInputHelp = ""
	}

	var cmd tea.Cmd
	m.TxtInput, cmd = m.TxtInput.Update(msg)
	return m, cmd

}

// ── view ────────────────────────────────────────────────────────────

var (
	homeTitleStyle = lipgloss.NewStyle().Bold(true).Italic(true).Underline(true)
	homeHelpStyle  = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

func homeView(m Model) string {
	tpl := homeTitleStyle.Render("Don't sleep!") + "\n\n"
	tpl += choiceView(m)

	return "\n" + tpl + "\n"

}

// ── sub-views ───────────────────────────────────────────────────────

var (
	homeListTitleStyle    = lipgloss.NewStyle()
	homeItemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	homeSelectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	homeTxtInputStyle     = lipgloss.NewStyle().MarginLeft(1)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)

	fn := homeItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return homeSelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

const (
	TimeChoice    = "after some time passes"
	ProcessChoice = "when a process ends"
)

func getHomeViewList() list.Model {
	items := []list.Item{
		item(TimeChoice),
		item(ProcessChoice),
	}

	const (
		defaultWidth  = 30
		defaultHeight = 4
	)

	l := list.New(items, itemDelegate{}, defaultWidth, defaultHeight)
	l.Title = "When can I sleep again?"
	l.SetShowStatusBar(false)
	l.SetShowPagination(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = homeListTitleStyle

	return l
}

func getChoiceInput() textinput.Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 20
	return ti
}

func choiceView(m Model) string {
	if m.Choice == "" {
		return m.List.View() + "\n\n"
	}

	var mode string
	if m.Choice == TimeChoice {
		mode = "time"
		m.TxtInput.Placeholder = "2d 1h 2m 30s"
	}
	if m.Choice == ProcessChoice {
		mode = "process"
		m.TxtInput.Placeholder = "000000"
	}

	tpl := homeTxtInputStyle.Render(fmt.Sprintf("%s %s", mode, m.TxtInput.View())) + "\n"

	if m.TxtInputHelp != "" {
		tpl += "\n" + homeTxtInputStyle.Render("? "+m.TxtInputHelp) + "\n"
	}

	return tpl + "\n"
}

// ── cmd ─────────────────────────────────────────────────────────────

type startCaffeinateMsg struct {
	process   int
	isProcess bool
	seconds   int
	isTime    bool
}

func startCaffeinate(m Model, input int) tea.Cmd {
	return func() tea.Msg {
		var msg startCaffeinateMsg
		if m.Choice == ProcessChoice {
			msg.isProcess, msg.process = true, input
		} else if m.Choice == TimeChoice {
			msg.isTime, msg.seconds = true, input
		}
		return msg
	}
}

func inputToProcess(txt string) (int, error) {
	val, err := strconv.Atoi(txt)
	if err != nil {
		return val, errors.New("input must be an integer")
	}
	return val, nil
}

func inputToTime(txt string) (int, error) {
	unitConversion := map[string]int{
		"d": 24 * 60 * 60,
		"h": 60 * 60,
		"m": 60,
		"s": 1,
	}
	splits := strings.Split(txt, " ")
	if len(splits) < 1 {
		return 0, errors.New("empty input")
	}
	var time int
	for _, split := range splits {
		digit, unit := split[:len(split)-1], split[len(split)-1:]
		val, err := strconv.Atoi(digit)
		if err != nil {
			return 0, errors.New("example: '2d 3h 7m 2s'")
		}
		conv, ok := unitConversion[unit]
		if !ok {
			return 0, errors.New("example: '2d 3h 7m 2s'")
		}
		time += val * conv
	}
	return time, nil
}
