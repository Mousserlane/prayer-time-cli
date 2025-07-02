package main

import (
	"fmt"
	"os"
	commandline "prayer-time-cli/cmd/cli"
	"prayer-time-cli/internal/config"
)

func runPrompt(configPath string) {
	conf, configErr := config.PromptForConfig()
	if configErr != nil {
		fmt.Errorf("Failed to get config from user: %v", configErr)
	}

	if saveError := config.SaveConfig(configPath, conf); saveError != nil {
		fmt.Errorf("Failed to save config from user: %v", saveError)
	} else {
		fmt.Printf("Configuration successfully saved to %s\n", configPath)
	}
}

func main() {
	var appConfig config.PrayerTimeConfig

	configPath := config.GetConfigPath()

	if len(os.Args) > 1 && os.Args[1] == "init" {
		runPrompt(configPath)
		os.Exit(0)
	}

	loadedConfig, loadError := config.LoadConfig(configPath)

	if loadError != nil || !loadedConfig.IsConfigComplete() {
		fmt.Printf("Config file not found or incomplete, starting prompt...")
		runPrompt(configPath)
	} else {
		appConfig = loadedConfig
	}

	commandline.RunApp(appConfig)
}
