package commandline

import (
	"fmt"
	"strings"

	prayertime "prayer-time-cli/internal/prayertime"

	"github.com/charmbracelet/lipgloss"
)

var (
	digitColor                    = lipgloss.NewStyle().Foreground(lipgloss.Color("#BDFE58"))
	dateStyle                     = lipgloss.NewStyle().Foreground(lipgloss.Color("#BDFE58")).Bold(true).Padding(0, 2)
	dateContainerStyle            = lipgloss.NewStyle().Align(lipgloss.Center)
	clockContainerStyle           = lipgloss.NewStyle().Align(lipgloss.Center).Margin(2, 0)
	todayPrayerTimeContainerStyle = lipgloss.NewStyle().Align(lipgloss.Center)
	prayerTimeBoxStyle            = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#BDFE58")).
					Bold(true).
					Padding(0, 1).
					Border(lipgloss.RoundedBorder()).
					BorderForeground(lipgloss.Color("#BDFE58")).
					Align(lipgloss.Center)
	highlightBoxStyle = lipgloss.NewStyle().
				Inherit(prayerTimeBoxStyle).
				Foreground(lipgloss.Color("#FFA500")).
				BorderForeground(lipgloss.Color("#FFA500")).
				Width(8)
	sunriseSectionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Margin(1, 0)
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

	renderedClock := strings.Join(outputLines, "\n")
	renderedPrayerTimes := renderTodayPrayerTimes(m.dailyPrayerTimes)
	hijriDate := prayertime.DateNowToHijri(m.currentTime)
	combinedDate := fmt.Sprintf("%s / %s", m.currentTime.UTC().Format("2 January 2006"), hijriDate)

	var sections []string

	sections = append(sections, clockContainerStyle.Width(m.width).Render(renderedClock))
	sections = append(sections, dateContainerStyle.Width(m.width).Render(dateStyle.Render(combinedDate)))
	sections = append(sections, todayPrayerTimeContainerStyle.Render(renderedPrayerTimes))

	return lipgloss.JoinVertical(lipgloss.Center, sections...)
}

func renderTodayPrayerTimes(times []prayertime.PrayerTime) string {
	var boxes []string
	var sunrise string
	var timeSection []string

	for index, prayerTime := range times {
		if index == 0 {
			// skip adding sunrise to the box
			sunrise = fmt.Sprintf("Sunrise is at: %s", prayerTime.Time)
			timeSection = append(timeSection, sunriseSectionStyle.Render(sunrise))
			continue
		}
		content := fmt.Sprintf("%s\n%s", prayerTime.Name, prayerTime.Time)
		if prayerTime.IsNearest {
			boxes = append(boxes, highlightBoxStyle.Render(content))
		} else {
			boxes = append(boxes, prayerTimeBoxStyle.Render(content))
		}
	}

	prayerTimeBoxes := lipgloss.JoinHorizontal(lipgloss.Center, boxes...)

	timeSection = append(timeSection, prayerTimeBoxes)
	return lipgloss.JoinVertical(lipgloss.Center, timeSection...)
}
