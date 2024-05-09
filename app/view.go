package app

type Window int

const (
	Home Window = iota
	On
)

func (m Model) View() string {
	return ""
}

func homeView(m Model) string {
	return ""
}

func onView(m Model) string {
	return ""
}
