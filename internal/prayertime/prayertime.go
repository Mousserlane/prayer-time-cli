package prayertime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type PrayerTime struct {
	Name string
	Time string
}

type AladhanResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		Timings struct {
			Fajr    string `json:"Fajr"`
			Dhuhr   string `json:"Dhuhr"`
			Asr     string `json:"Asr"`
			Maghrib string `json:"Maghrib"`
			Isha    string `json:"Isha"`
			// Add other timings if needed, e.g., Sunrise, Midnight
		} `json:"timings"`
		// Other fields like date, Gregorian, Hijri, etc.
	} `json:"data"`
}

func GetTodayPrayerTime(today string, city string, country string, method int) ([]PrayerTime, error) {
	apiURL := fmt.Sprintf("https://api.aladhan.com/v1/timingsByCity/%s?city=%s&country=%s&method=%d",
		today,
		strings.ReplaceAll(city, " ", "%20"),
		strings.ReplaceAll(country, " ", "%20"),
		method,
	)

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request to %s: %w", apiURL, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("API return non 200 status %s: %s", response.Status, string(bodyBytes))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response body: %w", err)
	}

	var aladhanResponse AladhanResponse
	err = json.Unmarshal(body, &aladhanResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal API response: %w (body: %s)", err, string(body))
	}

	// Check API's internal status code
	if aladhanResponse.Code != 200 {
		return nil, fmt.Errorf("Aladhan API error code %d: %s", aladhanResponse.Code, aladhanResponse.Status)
	}

	dailyPrayerTimes := []PrayerTime{
		{Name: "Fajr", Time: aladhanResponse.Data.Timings.Fajr},
		{Name: "Dhuhr", Time: aladhanResponse.Data.Timings.Dhuhr},
		{Name: "Asr", Time: aladhanResponse.Data.Timings.Asr},
		{Name: "Maghrib", Time: aladhanResponse.Data.Timings.Maghrib},
		{Name: "Isha", Time: aladhanResponse.Data.Timings.Isha},
	}

	return dailyPrayerTimes, nil
}
