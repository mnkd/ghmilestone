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
	Version  string
	Revision string
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
	var repo string
	var list, version bool

	flag.StringVar(&configPath, "c", "", "/path/to/config.json. (default: $HOME/.config/ghmilestone/config.json)")
	flag.StringVar(&milestone, "m", "", "milestone number")
	flag.StringVar(&repo, "r", "", "repo")
	flag.BoolVar(&version, "v", false, "Print version.")
	flag.BoolVar(&list, "list", false, "Print milestone list.")
	flag.Parse()

	if version {
		fmt.Fprintln(os.Stdout, "Version:", Version)
		fmt.Fprintln(os.Stdout, "Revision:", Revision)
		os.Exit(ExitCodeOK)
	}

	if len(repo) == 0 {
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
	app, err = NewApp(config, list, repo, milestone)
	if err != nil {
		os.Exit(ExitCodeError)
	}
}

func main() {
	os.Exit(app.Run())
}
