package domain

import (
	"slices"
	"sort"
	"strings"

	"github.com/zlasd/tzloc"
)

var IanaTimezonesByRegion map[string][]string

var Continents = []string{
	"Africa",
	"America",
	"Asia",
	"Atlantic",
	"Australia",
	"Europe",
	"Pacific",
}

func findIn(source []string, key string) bool {
	if len(source) == 0 {
		return false
	}

	return slices.Contains(source, key)
}

func init() {
	Timezones := tzloc.GetLocationList()

	IanaTimezonesByRegion = make(map[string][]string)

	for _, tz := range Timezones {
		region, _, found := strings.Cut(tz, "/")

		if found && findIn(Continents, region) {
			IanaTimezonesByRegion[region] = append(IanaTimezonesByRegion[region], tz)
		}
	}

	for region, sortedTzs := range IanaTimezonesByRegion {
		sort.Strings(sortedTzs)
		IanaTimezonesByRegion[region] = sortedTzs
	}
}
