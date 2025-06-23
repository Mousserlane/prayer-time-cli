package main

import (
	"log"
	"time"

	prayertime "prayer-time-cli/internal/prayertime"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hablullah/go-prayer"
)

type MonthlyPrayerData struct {
	Day   string
	Times []string
}

type fetchDailyPrayerResp struct {
	Prayers []prayertime.PrayerTime
}

type schedulesLoadedResp struct {
	Schedules []prayer.Schedule
}

type fetchDailyPrayerErr struct {
	Err error
}

type model struct {
	currentTime       time.Time
	isQuitting        bool
	dailyPrayerTimes  []prayertime.PrayerTime
	monthlyPrayerData []prayertime.MonthlyPrayerTime
	yearlySchedules   []prayer.Schedule
	isLoadingPrayer   bool
	Error             error
	width             int
	height            int
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	m.isLoadingPrayer = true
	return tea.Batch(
		tea.EnterAltScreen,
		tickCmd(),
		m.loadSchedules(),
		// m.fetchDailyPrayerTimes(),
	)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tickMsg:
		m.currentTime = time.Now()
		highlightChanged := m.updateUpcomingPrayerTime()
		if highlightChanged {
			return m, tickCmd()
		}

		return m, tickCmd()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case schedulesLoadedResp:
		m.yearlySchedules = msg.Schedules
		return m, m.fetchDailyPrayerTimes()

	case fetchDailyPrayerResp:
		m.isLoadingPrayer = false
		m.dailyPrayerTimes = msg.Prayers
		m.updateUpcomingPrayerTime()
		return m, nil

	case fetchDailyPrayerErr:
		m.isLoadingPrayer = false
		m.Error = msg.Err
		return m, nil
	}

	return m, nil
}

func (m *model) loadSchedules() tea.Cmd {
	return func() tea.Msg {
		schedules := prayertime.LoadSchedules(m.currentTime.Year())
		return schedulesLoadedResp{schedules}
	}
}

// TODO : Add city, country, and method arguments
func (m model) fetchDailyPrayerTimes() tea.Cmd {
	return func() tea.Msg {
		todayString := m.currentTime.Format("2006-01-02")

		prayerTimeResp := prayertime.GetTodaySchedule(todayString, m.yearlySchedules)

		return fetchDailyPrayerResp{Prayers: prayerTimeResp}
	}
}

func (m *model) updateUpcomingPrayerTime() bool {
	if len(m.dailyPrayerTimes) == 0 {
		return false
	}

	now := m.currentTime
	nearestIndex := -1
	minDuration := time.Hour * 24 * 365

	lastPassedIndex := -1
	maxPassedDuration := -time.Hour * 24 * 365

	oldNearestIndex := -1
	for i := range m.dailyPrayerTimes {
		m.dailyPrayerTimes[i].IsNearest = false
	}

	for i, pt := range m.dailyPrayerTimes {
		parsedTime, err := time.Parse("15:04:05", pt.Time)
		if err != nil {
			log.Printf("Error parsing prayer time %s: %v ", pt.Name, err)
			continue
		}

		prayerTimeToday := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())

		diff := prayerTimeToday.Sub(now)
		if diff > 0 {
			if diff < minDuration {
				minDuration = diff
				nearestIndex = i
			}
		} else {
			if diff > maxPassedDuration {
				maxPassedDuration = diff
				lastPassedIndex = i
			}
		}
	}

	newNearestIndex := -1
	if nearestIndex != -1 {
		newNearestIndex = nearestIndex
	} else if lastPassedIndex != -1 {
		// TODO : If all prayer time has passed, should it highlight the last prayer time or
		// the first?? for now, it's highlighting the last prayer time
		newNearestIndex = lastPassedIndex
	}

	if newNearestIndex != -1 {
		m.dailyPrayerTimes[newNearestIndex].IsNearest = true
	}

	return newNearestIndex != oldNearestIndex
}
