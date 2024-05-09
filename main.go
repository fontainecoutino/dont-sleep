package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fcoutino/dont-sleep/app"
)

func main() {
	p := tea.NewProgram(app.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program", err)
	}

}
