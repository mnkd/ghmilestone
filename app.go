package main

import (
	"fmt"
	"os"
)

type App struct {
	Config    Config
	PrintList bool
	Repo      string
	Milestone string
	GitHubAPI GitHubAPI
}

func (app App) printMilestones() int {
	var milestones []Milestone
	milestones, err := app.GitHubAPI.GetMilestones(app.Repo)
	if err != nil {
		return ExitCodeError
	}

	for _, milestone := range milestones {
		fmt.Fprintf(os.Stdout, "* [%v - %v](%v)\n", milestone.Number, milestone.Title, milestone.HTMLURL)
	}

	return ExitCodeOK
}

func (app App) printIssues() int {
	var issues []Issue
	issues, err := app.GitHubAPI.GetMilestoneIssues(app.Repo, app.Milestone)
	if err != nil {
		return ExitCodeError
	}

	var issueItems []Issue
	var pullItems []Issue

	for _, issue := range issues {
		if len(issue.PullRequest.URL) == 0 {
			issueItems = append(issueItems, issue)
		} else {
			pullItems = append(pullItems, issue)
		}
	}

	fmt.Fprintf(os.Stdout, "***** ISSUE *****\n")
	for _, issue := range issueItems {
		fmt.Fprintf(os.Stdout, "* [%v - %v](%v)\n", issue.Number, issue.Title, issue.HTMLURL)
	}

	fmt.Fprintf(os.Stdout, "\n***** PULL REQUEST *****\n")
	for _, issue := range pullItems {
		fmt.Fprintf(os.Stdout, "* [%v - %v](%v)\n", issue.Number, issue.Title, issue.HTMLURL)
	}

	return ExitCodeOK
}

func (app App) Run() int {
	if app.PrintList {
		return app.printMilestones()
	}
	return app.printIssues()
}

func NewApp(config Config, printList bool, repo string, milestone string) (App, error) {
	var app = App{}
	var err error
	app.Config = config
	app.PrintList = printList
	app.Repo = repo
	app.Milestone = milestone
	app.GitHubAPI = NewGitHubAPI(config)
	return app, err
}
