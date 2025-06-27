package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		// Run config prompt
	}
	p := tea.NewProgram(&model{
		currentTime: time.Now(),
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error while running the program %v\n", err)
	}
}
