package main

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

func loadToken() (*oauth2.Token, error) {
	var token oauth2.Token
	var dirname, err = os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(dirname + tokenPath + tokenFile); os.IsNotExist(err) {
		return nil, nil
	}

	tokenData, err := os.ReadFile(dirname + tokenPath + tokenFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(tokenData, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func saveToken(token *oauth2.Token) error {
	dirname, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	tokenData, err := json.Marshal(token)
	if err != nil {
		return err
	}

	if _, err := os.Stat(dirname + tokenPath + tokenFile); os.IsNotExist(err) {
		err_mkd := os.MkdirAll(dirname+tokenPath, os.ModePerm)
		if err_mkd != nil {
			return err_mkd
		}
	}

	err = os.WriteFile(dirname+tokenPath+tokenFile, tokenData, 0600)
	if err != nil {
		return err
	}

	return nil
}
