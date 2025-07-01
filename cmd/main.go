package main

import (
	"os"
	commandline "prayer-time-cli/cmd/cli"
	"prayer-time-cli/internal/config"
)

func main() {
	var appConfig config.PrayerTimeConfig

	if len(os.Args) > 1 && os.Args[1] == "init" {
		newConfig, _ := config.PromptForConfig()
		// Run config prompt
		// newConfig, err := config.PromptForConfig()
		appConfig = newConfig
	}

	print()
	commandline.RunApp(appConfig)
}
