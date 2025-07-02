package prayertime

import (
	"log"
	"prayer-time-cli/internal/config"
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

func LoadSchedules(year int, config config.PrayerTimeConfig) []prayer.Schedule {
	timezone, _ := time.LoadLocation(config.Timezone)

	schedulesYearly, _ := prayer.Calculate(prayer.Config{
		Latitude:           config.Latitude,
		Longitude:          config.Longitude,
		Timezone:           timezone,
		TwilightConvention: config.TwilightConvention,
		AsrConvention:      config.AsrConvention,
		PreciseToSeconds:   true,
	}, year)

	return schedulesYearly
}

func formatToHHMM(timeStr string) string {
	parsedTime, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		log.Printf("Error while parsing time string: %v", err)
	}
	return parsedTime.Format("15:04")
}

func GetTodaySchedule(todayISO string, schedules []prayer.Schedule) []PrayerTime {
	todayPrayerTime := searchDate(todayISO, schedules)
	// Return empty if there's no matching date
	if todayPrayerTime == nil {
		return []PrayerTime{}
	}

	format := "15:04:05"
	prayerTimeToday := []PrayerTime{
		{Name: "Sunrise", Time: formatToHHMM(todayPrayerTime.Sunrise.Format(format)), IsNearest: false},
		{Name: "Fajr", Time: formatToHHMM(todayPrayerTime.Fajr.Format(format)), IsNearest: false},
		{Name: "Dhuhr", Time: formatToHHMM(todayPrayerTime.Zuhr.Format(format)), IsNearest: false},
		{Name: "Asr", Time: formatToHHMM(todayPrayerTime.Asr.Format(format)), IsNearest: false},
		{Name: "Maghrib", Time: formatToHHMM(todayPrayerTime.Maghrib.Format(format)), IsNearest: false},
		{Name: "Isha", Time: formatToHHMM(todayPrayerTime.Isha.Format(format)), IsNearest: false},
	}

	return prayerTimeToday
}

func searchDate(targetDate string, schedules []prayer.Schedule) *prayer.Schedule {
	parsedTime, error := time.Parse("2006-01-02", targetDate)
	if error != nil {
		log.Printf("Error while parsing date %v", error)
	}
	month := parsedTime.Month()
	year := parsedTime.Year()

	// TODO : Implement leap year
	startIndex := MonthStartIndex[month]
	if isLeapYear(year) {
		startIndex += 1
	}

	endIndex := startIndex + daysInMonth(int(month), year)

	for startIndex <= endIndex {
		pivot := (startIndex + endIndex) / 2

		pivotDate := schedules[pivot].Date
		if pivotDate == targetDate {
			return &schedules[pivot]
		} else if pivotDate < targetDate {
			startIndex = pivot + 1
		} else {
			endIndex = pivot - 1
		}
	}

	return nil
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func daysInMonth(month, year int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}
