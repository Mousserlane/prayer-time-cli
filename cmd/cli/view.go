package main

import (
	"fmt"
	"strings"

	prayertime "prayer-time-cli/internal/prayertime"

	"github.com/charmbracelet/lipgloss"
)

// func getLine(lines []string, index int) string {
// 	if index >= 0 && index < len(lines) {
// 		return lines[index]
// 	}
// 	if len(lines) > 0 {
// 		return strings.Repeat(" ", len(lines[0]))
// 	}
// 	return ""
// }

var (
	digitColor          = lipgloss.NewStyle().Foreground(lipgloss.Color("#BDFE58"))
	dateStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#BDFE58")).Bold(true).Padding(0, 2)
	mainContainerStyle  = lipgloss.NewStyle().Width(80)
	dateContainerStyle  = lipgloss.NewStyle().Width(80).Align(lipgloss.Center)
	clockContainerStyle = lipgloss.NewStyle().Width(80).Align(lipgloss.Center).MarginTop(2).MarginBottom(2)
)

func (m model) View() string {
	if m.isQuitting {
		return "See you around!"
	}

	hour := m.currentTime.Hour()
	minute := m.currentTime.Minute()
	second := m.currentTime.Second()

	hDigit1 := hour / 10
	hDigit2 := hour % 10
	mDigit1 := minute / 10
	mDigit2 := minute % 10
	sDigit1 := second / 10
	sDigit2 := second % 10

	hDigit1Lines := getDigit(hDigit1)
	hDigit2Lines := getDigit(hDigit2)
	mDigit1Lines := getDigit(mDigit1)
	mDigit2Lines := getDigit(mDigit2)
	sDigit1Lines := getDigit(sDigit1)
	sDigit2Lines := getDigit(sDigit2)

	separatorLines := getSeparator()

	digitHeight := len(hDigit1Lines)

	outputLines := make([]string, digitHeight)

	for i := 0; i < digitHeight; i++ {
		outputLines[i] = fmt.Sprintf("%s %s %s %s %s %s %s %s",
			digitColor.Render(hDigit1Lines[i]), digitColor.Render(hDigit2Lines[i]),
			digitColor.Render(separatorLines[i]),
			digitColor.Render(mDigit1Lines[i]), digitColor.Render(mDigit2Lines[i]),
			digitColor.Render(separatorLines[i]),
			digitColor.Render(sDigit1Lines[i]), digitColor.Render(sDigit2Lines[i]),
		)
	}

	hijriDate := prayertime.DateNowToHijri()
	renderedClock := strings.Join(outputLines, "\n")

	const assumeWidth = 120

	var sections []string

	sections = append(sections, clockContainerStyle.Render(renderedClock))
	sections = append(sections, dateContainerStyle.Render(dateStyle.Render(hijriDate)))
	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}
