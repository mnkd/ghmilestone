package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Issue struct {
	Assignee struct {
		ID    int64  `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"assignee"`
	Body string `json:"body"`
	// ClosedAt    string        `json:"closed_at"`
	// Comments    int64         `json:"comments"`
	// CreatedAt   string        `json:"created_at"`
	ID int64 `json:"id"`
	// Labels      []interface{} `json:"labels"`
	// Locked      bool          `json:"locked"`
	// Milestone   struct {
	// 	ClosedAt     interface{} `json:"closed_at"`
	// 	ClosedIssues int64       `json:"closed_issues"`
	// 	CreatedAt    string      `json:"created_at"`
	// 	Creator      struct {
	// 		ID                int64  `json:"id"`
	// 		Login             string `json:"login"`
	// 		URL               string `json:"url"`
	// 	} `json:"creator"`
	// 	Description string `json:"description"`
	// 	DueOn       string `json:"due_on"`
	// 	ID          int64  `json:"id"`
	// 	Number      int64  `json:"number"`
	// 	OpenIssues  int64  `json:"open_issues"`
	// 	State       string `json:"state"`
	// 	Title       string `json:"title"`
	// 	UpdatedAt   string `json:"updated_at"`
	// 	URL         string `json:"url"`
	// } `json:"milestone"`
	Number int64 `json:"number"`
	// RepositoryURL string `json:"repository_url"`
	State string `json:"state"`
	Title string `json:"title"`
	// UpdatedAt     string `json:"updated_at"`
	HTMLURL string `json:"html_url"`
	User    struct {
		ID    int64  `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
}

func (gh GitHubAPI) GetMilestoneIssues(repo string, milestone string) ([]Issue, error) {
	var issues []Issue

	// Prepare HTTP Request
	url := "https://api.github.com/repos/" + gh.Owner + "/" + repo + "/issues" + "?access_token=" + gh.AccessToken + "&milestone=" + milestone + "&state=all"

	req, err := http.NewRequest("GET", url, nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Fprintln(os.Stderr, "GitHubAPI - Issue Milestone: <error> parse http request form:", parseFormErr)
		return issues, parseFormErr
	}

	// Fetch Request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "GitHubAPI - Issue Milestone: <error> fetch issues:", err)
		return issues, err
	}

	// Read Response Body
	resBody, _ := ioutil.ReadAll(res.Body)

	// Display Results
	// fmt.Println("response Status : ", res.Status)
	// fmt.Println("response Headers : ", res.Header)
	// fmt.Println("response Body : ", string(resBody))

	// Decode JSON
	if err := json.Unmarshal(resBody, &issues); err != nil {
		fmt.Fprintln(os.Stderr, "GitHubAPI - Issue Milestone: <error> json unmarshal:", err)
		return issues, err
	}

	return issues, nil
}
