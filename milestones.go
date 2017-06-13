package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Milestone struct {
	ID      int64  `json:"id"`
	Number  int64  `json:"number"`
	State   string `json:"state"`
	Title   string `json:"title"`
	HTMLURL string `json:"html_url"`
}

func (gh GitHubAPI) GetMilestones(repo string) ([]Milestone, error) {
	var milestones []Milestone

	// Prepare HTTP Request
	url := "https://api.github.com/repos/" + gh.Owner + "/" + repo + "/milestones" + "?access_token=" + gh.AccessToken + "&state=open"

	req, err := http.NewRequest("GET", url, nil)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Fprintln(os.Stderr, "GitHubAPI - Milestones: <error> parse http request form:", parseFormErr)
		return milestones, parseFormErr
	}

	// Fetch Request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "GitHubAPI - Milestones: <error> fetch milestones:", err)
		return milestones, err
	}

	// Read Response Body
	resBody, _ := ioutil.ReadAll(res.Body)

	// Display Results
	// fmt.Println("response Status : ", res.Status)
	// fmt.Println("response Headers : ", res.Header)
	// fmt.Println("response Body : ", string(resBody))

	// Decode JSON
	if err := json.Unmarshal(resBody, &milestones); err != nil {
		fmt.Fprintln(os.Stderr, "GitHubAPI - Milestones: <error> json unmarshal:", err)
		return milestones, err
	}

	return milestones, nil
}
