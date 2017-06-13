package main

import (
	"fmt"
	"os"
)

type App struct {
	Config    Config
	Repo      string
	Milestone string
	GitHubAPI GitHubAPI
}

func (app App) Run() int {
	var issues []Issue
	issues, err := app.GitHubAPI.GetMilestoneIssues(app.Repo, app.Milestone)
	if err != nil {
		return ExitCodeError
	}

	for _, issue := range issues {
		fmt.Fprintf(os.Stdout, "* [%v - %v](%v)\n", issue.Number, issue.Title, issue.HTMLURL)
	}

	return ExitCodeOK
}

func NewApp(config Config, repo string, milestone string) (App, error) {
	var app = App{}
	var err error
	app.Config = config
	app.Repo = repo
	app.Milestone = milestone
	app.GitHubAPI = NewGitHubAPI(config)
	return app, err
}
