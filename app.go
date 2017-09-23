package main

import (
	"fmt"
	"os"
)

type App struct {
	Config    Config
	PrintList bool
	Milestone string
	State     string // open, closed, all
	GitHubAPI GitHubAPI
}

func (app App) printMilestones() int {
	var milestones []Milestone
	milestones, err := app.GitHubAPI.GetMilestones(app.State)
	if err != nil {
		return ExitCodeError
	}

	for _, milestone := range milestones {
		fmt.Fprintf(os.Stdout, "* [%v - %v](%v)\n", milestone.Number, milestone.Title, milestone.HTMLURL)
	}

	return ExitCodeOK
}

func filterIssues(issues []Issue, state string) []Issue {
	var filterd []Issue
	for _, issue := range issues {
		if issue.State == state {
			filterd = append(filterd, issue)
		}
	}
	return filterd
}

func (app App) printIssues(issues []Issue, title string) {
	fmt.Fprintf(os.Stdout, "# %v\n", title)
	openIssues := filterIssues(issues, "open")
	closedIssues := filterIssues(issues, "closed")

	fmt.Fprintf(os.Stdout, "\n## OPEN (%v)\n", len(openIssues))
	for _, issue := range openIssues {
		fmt.Fprintf(os.Stdout, "* [%v - %v](%v) (%v)\n", issue.Number, issue.Title, issue.HTMLURL, issue.Assignee.Login)
	}

	fmt.Fprintf(os.Stdout, "\n## CLOSED (%v)\n", len(closedIssues))
	for _, issue := range closedIssues {
		fmt.Fprintf(os.Stdout, "* [%v - %v](%v) (%v)\n", issue.Number, issue.Title, issue.HTMLURL, issue.Assignee.Login)
	}
}

func (app App) printMilestoneIssues() int {
	// Get milestone issues from GitHub
	var issues []Issue
	issues, err := app.GitHubAPI.GetMilestoneIssues(app.Milestone)
	if err != nil {
		return ExitCodeError
	}

	// Divide issues
	var issueItems []Issue
	var pullItems []Issue

	for _, issue := range issues {
		if len(issue.PullRequest.URL) > 0 {
			pullItems = append(pullItems, issue)
		} else {
			issueItems = append(issueItems, issue)
		}
	}

	// Print issues and pull requests
	app.printIssues(issueItems, "ISSUE")
	fmt.Fprintln(os.Stdout, "")
	app.printIssues(pullItems, "PULL REQUEST")

	return ExitCodeOK
}

func (app App) Run() int {
	if app.PrintList {
		return app.printMilestones()
	}
	return app.printMilestoneIssues()
}

func NewApp(config Config, printList bool, owner string, repo string, milestone string, state string) (App, error) {
	var app = App{}
	var err error
	app.Config = config
	app.PrintList = printList
	app.Milestone = milestone
	app.State = state
	app.GitHubAPI = NewGitHubAPI(config, owner, repo)
	return app, err
}
