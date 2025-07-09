package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"prayer-time-cli/internal/domain"
	"slices"
	"strconv"
	"strings"

	"github.com/adrg/xdg"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/hablullah/go-prayer"
)

type PrayerTimeConfig struct {
	City               string                     `json:"city"`
	Continent          string                     `json:"continent"`
	TwilightConvention *prayer.TwilightConvention `json:"twilight_convention"`
	AsrConvention      prayer.AsrConvention       `json:"asr_convention,omitempty"`
	Timezone           string                     `json:"string"` // TODO : Add fallback based on City/Country based on tz database
	Latitude           float64                    `json:"longitude,omitempty"`
	Longitude          float64                    `json:"latitude,omitempty"`
	PreciseToSeconds   bool                       `json:"precise_to_seconds,omitempty"`
}

func (config PrayerTimeConfig) IsConfigComplete() bool {
	return config.Timezone != "" && config.TwilightConvention != nil && config.Continent != ""
}

func GetConfigPath() string {
	// TODO Get Config file here
	configHome, err := xdg.ConfigFile(filepath.Join("prayer-time-cli", "config.json"))
	if err != nil {
		panic("Unable to determine XDG config")
	}
	return configHome
}

func marshallFloat64(input float64) string {
	return strconv.FormatFloat(float64(input), 'f', 6, 64)
}

func unmarshallFloat64(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

func PromptForConfig() (PrayerTimeConfig, error) {
	var conf PrayerTimeConfig
	var tz_continents []huh.Option[string]
	var latitudeString string
	var longitudeString string
	var cityString string

	continents := make([]string, 0, len(domain.IanaTimezonesByRegion))
	for continent := range domain.IanaTimezonesByRegion {
		continents = append(continents, continent)
	}

	for _, tz_continent := range continents {
		tz_continents = append(tz_continents, huh.NewOption(tz_continent, tz_continent))
	}

	twilightConventions := []huh.Option[*prayer.TwilightConvention]{
		huh.NewOption("Astronomical Twilight", prayer.AstronomicalTwilight()),
		huh.NewOption("Muslim World League (MWL)", prayer.MWL()),
		huh.NewOption("Islamic Society of North America (ISNA)", prayer.ISNA()),
		huh.NewOption("Umm al-Qura", prayer.UmmAlQura()),
		huh.NewOption("Gulf", prayer.Gulf()),
		huh.NewOption("Algerian Ministry of Religious Affairs", prayer.Algerian()),
		huh.NewOption("University of Islamic Sciences (Karachi)", prayer.Karachi()),
		huh.NewOption("Diyanet İşleri Başkanlığı (Turkey)", prayer.Diyanet()),
		huh.NewOption("Egyptian General Authority of Survey", prayer.Egypt()),
		huh.NewOption("Egypt(BIS)", prayer.EgyptBis()),
		huh.NewOption("Kementrian Agama Republic Indonesia (Kemenag RI)", prayer.Kemenag()),
		huh.NewOption("Majlis Ugama Islam Singapura (MUIS)", prayer.MUIS()),
		huh.NewOption("Jabatan Kemajuan Islam Malaysia (JAKIM)", prayer.JAKIM()),
		huh.NewOption("Union Des Organisations Islamiques De France (UOIF)", prayer.UOIF()),
		huh.NewOption("France 15", prayer.France15()),
		huh.NewOption("France 18", prayer.France18()),
		huh.NewOption("Tunisian Ministry of Religious Affairs (Tunisia)", prayer.Tunisia()),
		huh.NewOption("Institute of Geophysics at University of Tehran.", prayer.Tehran()),
		huh.NewOption("Jafari", prayer.Jafari()),
	}

	asrConvention := []huh.Option[prayer.AsrConvention]{
		huh.NewOption("Shafii", prayer.Shafii),
		huh.NewOption("Hanafi", prayer.Hanafi),
	}

	huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().
			Title("Please Select a Continent").
			Options(tz_continents...).
			Value(&conf.Continent),

		huh.NewInput().
			Title("Please Input Your City").
			Prompt(">> ").
			Value(&conf.City),
	)).
		WithProgramOptions(tea.WithAltScreen()).
		Run()

	huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Enter Latitude").
			Prompt(">> ").
			Value(&latitudeString),

		huh.NewInput().
			Title("Enter Longitude").
			Prompt(">> ").
			Value(&longitudeString),
	)).
		WithProgramOptions(tea.WithAltScreen()).
		Run()

	form := huh.NewForm(huh.NewGroup(

		huh.NewSelect[*prayer.TwilightConvention]().
			Title("Select A Twlilight Convention (To determine Fajr & Isha Angle)").
			Options(twilightConventions...).
			Value(&conf.TwilightConvention),

		huh.NewSelect[prayer.AsrConvention]().
			Title("Select Asr Convention").
			Options(asrConvention...).
			Value(&conf.AsrConvention),

		huh.NewSelect[string]().
			Title("Select Timezone").
			OptionsFunc(func() []huh.Option[string] {
				var options []huh.Option[string]

				timezones, ok := domain.IanaTimezonesByRegion[conf.Continent]

				// cityIndex := strings.IndexFunc(options, func(option huh.Option[string]) bool {
				// 	city := strings.Split(option.Key, "/")
				// 	return city[len(city)-1] == conf.City
				// })

				if !ok {
					timezones = domain.IanaTimezonesByRegion["General / Other"]
				}

				for _, tz := range timezones {
					options = append(options, huh.NewOption(tz, tz))
				}

				cityIndex := slices.IndexFunc(options, func(option huh.Option[string]) bool {
					cityPart := strings.Split(option.Key, "/")
					return cityPart[len(cityPart)-1] == conf.City
				})

				if cityIndex != -1 {
					return []huh.Option[string]{options[cityIndex]}
				}

				return options
			}, &conf.Continent).
			Value(&conf.Timezone),
	)).WithProgramOptions(tea.WithAltScreen())

	if err := form.Run(); err != nil {
		return conf, fmt.Errorf("Form exited with error: %w", err)
	}

	var unMarshalError error
	conf.Latitude, unMarshalError = unmarshallFloat64(latitudeString)

	if unMarshalError != nil {
		fmt.Errorf("Unable to parse latitude: %w", unMarshalError)
	}

	conf.Longitude, unMarshalError = unmarshallFloat64(longitudeString)
	if unMarshalError != nil {
		fmt.Errorf("Unable to parse longitude: %w", unMarshalError)
	}

	conf.City = cityString

	return conf, nil
}

func SaveConfig(path string, conf PrayerTimeConfig) error {
	data, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		fmt.Errorf("Failed to marshal config file: %w", err)
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("Failed to create config directory '%s': '%w'", dir, err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("Failed to write config file '%s':'%w'", path, err)
	}

	return nil
}

func LoadConfig(path string) (PrayerTimeConfig, error) {
	var conf PrayerTimeConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, fmt.Errorf("Failed to read config file '%s' : '%w'", path, err)
	}

	if err := json.Unmarshal(data, &conf); err != nil {
		return conf, fmt.Errorf("Failed to unmarshal config from '%s' : '%w'", path, err)
	}
	return conf, nil
}
