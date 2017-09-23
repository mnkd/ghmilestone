package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
	ExitCodeFileError
)

var (
	version  string
	revision string
)

func usage() {
	str := `Usage:
 ghmilestone [--list] [-r repo]       : Print milestones for a repository
 ghmilestone [-r repo] [-m milestone] : Print issues for a milestone.

Examples:
 $ ghmilestone --list -r awesome-app
 $ ghmilestone -r awesome-app -m 15
`
	fmt.Fprintln(os.Stderr, str)
}

var app App

func init() {
	var configPath string
	var milestone string
	var owner, repo string
	var list, ver bool

	flag.StringVar(&configPath, "c", "", "/path/to/config.json. (default: $HOME/.config/ghmilestone/config.json)")
	flag.StringVar(&owner, "o", "", "owner (e.g. github)")
	flag.StringVar(&repo, "r", "", "repo (e.g. hub)")
	flag.StringVar(&milestone, "m", "", "milestone number")
	flag.BoolVar(&list, "list", false, "Print milestone list.")
	flag.BoolVar(&ver, "v", false, "Print version.")
	flag.Parse()

	if ver {
		fmt.Fprintln(os.Stdout, "Version:", version)
		fmt.Fprintln(os.Stdout, "Revision:", revision)
		os.Exit(ExitCodeOK)
	}

	if len(owner) == 0 || len(repo) == 0 {
		usage()
		os.Exit(ExitCodeOK)
	}

	if list == false && len(milestone) == 0 {
		usage()
		os.Exit(ExitCodeOK)
	}

	// Prepare config
	config, err := NewConfig(configPath)
	if err != nil {
		os.Exit(ExitCodeError)
	}

	// Prepare app
	app, err = NewApp(config, list, owner, repo, milestone)
	if err != nil {
		os.Exit(ExitCodeError)
	}
}

func main() {
	os.Exit(app.Run())
}
