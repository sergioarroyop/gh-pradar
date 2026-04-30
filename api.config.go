package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func loadConfig() error {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(userHome, ".config", "gh-pradar")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	configFilePath := filepath.Join(dir, "config.json")
	jsonFile, err := os.Open(configFilePath)
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

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println(errorText("Ups, seems like your JSON config file is not well formated...: " + err.Error()))
		return err
	}

	if err := os.Chmod(configFilePath, 0600); err != nil {
		// Log a warning but don't fail — the file might already have correct perms
		fmt.Println(warnText("Warning: could not set config file permissions: " + err.Error()))
	}

	return err
}
