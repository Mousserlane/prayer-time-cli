package prayertime

import (
	"slices"
	"time"

	"github.com/hablullah/go-prayer"
)

type PrayerTime struct {
	Name      string
	Time      string
	IsLoading bool
	IsNearest bool
}

type MonthlyPrayerTime struct {
	Name string
	Time string
}

func LoadSchedules(year int) []prayer.Schedule {
	// TODO : Should load from config
	timezone, _ := time.LoadLocation("Asia/Jakarta")

	schedulesYearly, _ := prayer.Calculate(prayer.Config{
		// TODO : Load these from config
		Latitude:           -6.14,
		Longitude:          106.81,
		Timezone:           timezone,
		TwilightConvention: prayer.Kemenag(),
		AsrConvention:      prayer.Shafii,
		PreciseToSeconds:   true,
	}, year)

	return schedulesYearly
}

func GetTodaySchedule(todayISO string, schedules []prayer.Schedule) []PrayerTime {
	todayIdx := slices.IndexFunc(schedules, func(s prayer.Schedule) bool {
		return s.Date == todayISO
	})

	// Return empty if there's no matching date
	if todayIdx == -1 {
		return []PrayerTime{}
	}

	format := "15:04:05"
	prayerTimeToday := []PrayerTime{
		{Name: "Sunrise", Time: schedules[todayIdx].Sunrise.Format(format), IsNearest: false},
		{Name: "Fajr", Time: schedules[todayIdx].Fajr.Format(format), IsNearest: false},
		{Name: "Dhuhr", Time: schedules[todayIdx].Zuhr.Format(format), IsNearest: false},
		{Name: "Asr", Time: schedules[todayIdx].Asr.Format(format), IsNearest: false},
		{Name: "Maghrib", Time: schedules[todayIdx].Maghrib.Format(format), IsNearest: false},
		{Name: "Isha", Time: schedules[todayIdx].Isha.Format(format), IsNearest: false},
	}

	return prayerTimeToday
}
