package app

func (m Model) View() string {
	switch m.GetWindow() {
	case HomeWindow:
		return homeView(m)
	case OnWindow:
		return onView(m)
	case ErrorWindow:
		return errorView(m)
	case QuitWindow:
		return quitView(m)
	default:
		return homeView(m)
	}
}

func homeView(m Model) string {
	return "... home ..."
}

func onView(m Model) string {
	return "... on ..."
}

func errorView(m Model) string {
	return "... error ..."
}

func quitView(m Model) string {
	return "... qitting .."
}
