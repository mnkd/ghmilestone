package main

type GitHubAPI struct {
	AccessToken string
	Owner       string
}

func NewGitHubAPI(config Config) GitHubAPI {
	var gh = GitHubAPI{}
	gh.AccessToken = config.GitHub.AccessToken
	gh.Owner = config.GitHub.Owner
	return gh
}
