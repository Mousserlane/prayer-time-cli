package prayertime

import (
	"fmt"
	"time"

	hijri "github.com/hablullah/go-hijri"
)

var monthName = map[int]string{
	1:  "Muḥarram",
	2:  "Ṣafar",
	3:  "Rabī Al-Awwal",
	4:  "Rabī Al-Ākhir",
	5:  "Jumādā Al-ʾŪlā",
	6:  "Jumādā Al-Ākhirah",
	7:  "Rajab",
	8:  "Shaʿbān",
	9:  "Ramaḍān",
	10: "Shawwāl",
	11: "Ḏhu Qadah",
	12: "Ḏhu al-Ḥijjah",
}

func getMonthName(monthInt int) string {
	if month, ok := monthName[monthInt]; ok {
		return month
	}

	return "Invalid Month"
}

func DateNowToHijri() string {
	currentTime := time.Now()

	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()

	currentDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	hijriDate, _ := hijri.CreateUmmAlQuraDate(currentDate)

	monthName := getMonthName(int(hijriDate.Month))

	// Seems like the date is behind. Hence, adding 1 to the Day
	formattedHijriDate := fmt.Sprintf("%d %s %d", hijriDate.Day+1, monthName, hijriDate.Year)
	return formattedHijriDate
}
