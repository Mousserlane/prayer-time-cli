package domain

import (
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

	for _, val := range source {
		if key == val {
			return true
		}
	}

	return false
}

func init() {
	Timezones := tzloc.GetLocationList()

	IanaTimezonesByRegion = make(map[string][]string)

	for _, tz := range Timezones {
		region, _, found := strings.Cut(tz, "/")

		if found && findIn(Continents, region) {
			IanaTimezonesByRegion[region] = append(IanaTimezonesByRegion[region], tz)
		}
		// if !found && findIn(SpecialRegions, region) {
		// 	panic("Invalid Timezone format. It should follow the IANA TZ convention of Region/City")
		// }
	}

	for region, sortedTzs := range IanaTimezonesByRegion {
		sort.Strings(sortedTzs)
		IanaTimezonesByRegion[region] = sortedTzs
	}
}
