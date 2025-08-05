package main

import (
	"encoding/json"
	"io"
	"os"
)

func loadConfig(configPath string, configFile string) (appConfig, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		print(errorText("Error reading user home dir: " + err.Error()))
	}

	jsonFile, err := os.Open(userHome + configPath + configFile)
	if err != nil {
		print(errorText("Error opening JSON config file: " + err.Error()))
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	if err != nil {
		print(errorText("Error deconding JSON config file: " + err.Error()))
	}

	json.Unmarshal(byteValue, &config)

	return config, err
}
