package main

type appConfig struct {
	PAT          string   `json:"personal_access_token"`
	Owner        string   `json:"owner"`
	Sound        bool     `json:"sound"`
	ReposityList []string `json:repository_list`
}
