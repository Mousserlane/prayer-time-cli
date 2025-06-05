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
	digitColor                    = lipgloss.NewStyle().Foreground(lipgloss.Color("#BDFE58"))
	dateStyle                     = lipgloss.NewStyle().Foreground(lipgloss.Color("#BDFE58")).Bold(true).Padding(0, 2)
	mainContainerStyle            = lipgloss.NewStyle().Width(80)
	dateContainerStyle            = lipgloss.NewStyle().Width(80).Align(lipgloss.Center)
	clockContainerStyle           = lipgloss.NewStyle().Width(80).Align(lipgloss.Center).Margin(2, 0)
	todayPrayerTimeContainerStyle = lipgloss.NewStyle().Width(80).Align(lipgloss.Center)
	prayerTimeBoxStyle            = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#444444")).
					Background(lipgloss.Color("#FFFFFF")).
					Padding(0, 1).
					Border(lipgloss.RoundedBorder()).
					BorderForeground(lipgloss.Color("#BDFE58")).
					Align(lipgloss.Center)
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

	todayString := fmt.Sprintf("%d-%d-%d", m.currentTime.Day, m.currentTime.Month, m.currentTime.Year)
	prayerTimeResp, err := prayertime.GetTodayPrayerTime(todayString, "jakarta", "ID", 10)
	if err != nil {
		fmt.Errorf("Unable to get response from API: %v", err)
	}

	renderedClock := strings.Join(outputLines, "\n")
	renderedPrayerTimes := renderTodayPrayerTimes(prayerTimeResp)
	hijriDate := prayertime.DateNowToHijri(m.currentTime)

	const assumeWidth = 120

	var sections []string

	sections = append(sections, clockContainerStyle.Render(renderedClock))
	sections = append(sections, dateContainerStyle.Render(dateStyle.Render(hijriDate)))
	sections = append(sections, todayPrayerTimeContainerStyle.Render())
	sections = append(sections, todayPrayerTimeContainerStyle.Render(renderedPrayerTimes))
	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func renderTodayPrayerTimes(times []prayertime.PrayerTime) string {
	var boxes []string
	for _, p := range times {
		content := fmt.Sprintf("%s\n%s", p.Name, p.Time)
		boxes = append(boxes, prayerTimeBoxStyle.Render(content))
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, boxes...)
}
