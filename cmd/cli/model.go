package main

import (
	"time"

	prayertime "prayer-time-cli/internal/prayertime"

	tea "github.com/charmbracelet/bubbletea"
)

type WeeklyPrayerData struct {
	Day   string
	Times []string
}

type model struct {
	currentTime      time.Time
	isQuitting       bool
	dailyPrayerTimes []prayertime.PrayerTime
	weeklyPrayerData []WeeklyPrayerData
	width            int
	height           int
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tickCmd(),
	)
	// return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tickMsg:
		m.currentTime = time.Now()
		return m, tickCmd()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	return m, nil
}
