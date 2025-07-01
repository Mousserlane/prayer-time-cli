package commandline

var digits = map[int][]string{
	0: {
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	},
	1: {
		"  █",
		"  █",
		"  █",
		"  █",
		"  █",
	},
	2: {
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	},
	3: {
		"███",
		"  █",
		"███",
		"  █",
		"███",
	},
	4: {
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	},
	5: {
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	},
	6: {
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	},
	7: {
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	},
	8: {
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	},
	9: {
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	},
}

var separator = []string{
	"   ",
	" █ ",
	"   ",
	" █ ",
	"   ",
}

func getDigit(digit int) []string {
	if pattern, ok := digits[digit]; ok {
		return pattern
	}
	return []string{"???", "???", "???", "???", "???"}
}

func getSeparator() []string {
	return separator
}
