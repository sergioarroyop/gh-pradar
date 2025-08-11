package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func loadConfig(configPath string, configFile string) error {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	jsonFile, err := os.Open(userHome + configPath + configFile)
	if err != nil {
		fmt.Println(errorText("Ups, seems like you don't have a configuration file."))
		return err
	}

	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(errorText("Ups, seems like your JSON config file is not well formated...: " + err.Error()))
		return err
	}

	json.Unmarshal(byteValue, &config)

	return err
}
