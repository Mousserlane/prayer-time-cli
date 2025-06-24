package prayertime

import (
	"log"
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
	// todayIdx := slices.IndexFunc(schedules, func(s prayer.Schedule) bool {
	// 	return s.Date == todayISO
	// })

	todayPrayerTime := searchDate(todayISO, schedules)
	// Return empty if there's no matching date
	if todayPrayerTime == nil {
		return []PrayerTime{}
	}

	format := "15:04:05"
	prayerTimeToday := []PrayerTime{
		{Name: "Sunrise", Time: todayPrayerTime.Sunrise.Format(format), IsNearest: false},
		{Name: "Fajr", Time: todayPrayerTime.Fajr.Format(format), IsNearest: false},
		{Name: "Dhuhr", Time: todayPrayerTime.Zuhr.Format(format), IsNearest: false},
		{Name: "Asr", Time: todayPrayerTime.Asr.Format(format), IsNearest: false},
		{Name: "Maghrib", Time: todayPrayerTime.Maghrib.Format(format), IsNearest: false},
		{Name: "Isha", Time: todayPrayerTime.Isha.Format(format), IsNearest: false},
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
	isLeapYear := IsLeapYear(year)

	// TODO : Implement leap year
	startIndex := MonthStartIndex[month]
	if isLeapYear {
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

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func daysInMonth(month, year int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}
