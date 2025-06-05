package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	p := tea.NewProgram(model{
		currentTime: time.Now(),
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error while running the program %v\n", err)
	}
}
