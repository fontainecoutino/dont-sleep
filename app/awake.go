package app

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
		m.Exiting = true
		m.SetWindow(QuitWindow)
		return m, tea.Quit
	case tickMsg:
		if m.TimeAwake+1 >= math.MaxInt32 {
			return m, nil
		}
		m.TimeAwake++
		return m, tick()
	case startCaffeinateMsg:
		m.Awake = true
		cmdCtx, cancelFunc := context.WithCancel(m.Ctx)
		m.CancelCmd = cancelFunc
		return m, tea.Batch(caffeinate(cmdCtx, msg), tick())
	}
	return m, nil
}

// ── view ────────────────────────────────────────────────────────────

var (
	awakeSteamStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	awakeLiquidStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	awakeCupStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	awakeCupDetailStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
)

func awakeView(m Model) string {
	steam1 := awakeSteamStyle.Render("      (      ") + "\n"
	steam1 += awakeSteamStyle.Render("       )   ) ") + "\n"
	steam1 += awakeSteamStyle.Render("      (   (  ") + "\n"
	steam1 += awakeLiquidStyle.Render("    _______") +
		awakeSteamStyle.Render(")") +
		awakeLiquidStyle.Render("_")

	steam2 := awakeSteamStyle.Render("      )      ") + "\n"
	steam2 += awakeSteamStyle.Render("     (     ( ") + "\n"
	steam2 += awakeSteamStyle.Render("      )     )") + "\n"
	steam2 += awakeLiquidStyle.Render("    _______") +
		awakeSteamStyle.Render("(") +
		awakeLiquidStyle.Render("_")

	steam := steam1
	switch m.TimeAwake % 4 {
	case 0:
		steam = steam1
	case 1:
		steam = steam1
	case 2:
		steam = steam2
	case 3:
		steam = steam2
	}

	coffee := awakeCupStyle.Render(" .-'---------|") + "\n"
	coffee += awakeCupStyle.Render("( C|") +
		awakeCupDetailStyle.Render(". . . . .") +
		awakeCupStyle.MarginLeft(0).Render("|") + "\n"
	coffee += awakeCupStyle.Render(" '-.         |") + "\n"
	coffee += awakeCupStyle.Render("   '_________'") + "\n"
	coffee += awakeCupStyle.Render("    '-------' ") + "\n"

	tpl := steam + "\n"
	tpl += coffee + "\n"

	return "\n" + tpl

}

// ── cmd ─────────────────────────────────────────────────────────────

type caffeinateFinishedMsg struct{ err error }

func caffeinate(ctx context.Context, msg startCaffeinateMsg) tea.Cmd {
	return func() tea.Msg {
		var opt string
		if msg.isProcess {
			opt = fmt.Sprintf("-w %d", msg.process)
		} else if msg.isTime {
			opt = fmt.Sprintf("-t %d", msg.seconds)
		} else {
			opt = "-t 1"
		}

		awakeCmd := exec.CommandContext(ctx, "caffeinate", "-d", opt)
		awakeCmd.Stdout = nil

		err := awakeCmd.Start()
		if err != nil {
			return caffeinateFinishedMsg{err: errors.New("error while starting caffeinate")}
		}
		awakeCmd.Cancel = awakeCmd.Process.Kill

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

type tickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}
