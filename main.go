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
	str := `
Usage:

 ghmilestone -o owner -r repo --list [--state open] : Print milestones for a repository
 ghmilestone -o owner -r repo -m milestone          : Print issues for a milestone.

Examples:

 Print owner/awesome-app/milesstones
 $ ghmilestone -o owner -r awesome-app --list

 Print owner/awesome-app/milesstones?state=open
 $ ghmilestone -o owner -r awesome-app --list --state open

 Print owner/awesome-app/milesstones?state=closed
 $ ghmilestone -o owner -r awesome-app --list --state closed

 Print owner/awesome-app/milesstone/15
 $ ghmilestone -o owner -r awesome-app -m 15
`
	fmt.Fprintln(os.Stderr, str)
}

var app App

func init() {
	var configPath string
	var milestone string
	var owner, repo, state string
	var list, ver bool

	flag.StringVar(&configPath, "c", "", "/path/to/config.json. (default: $HOME/.config/ghmilestone/config.json)")
	flag.StringVar(&owner, "o", "", "owner (e.g. github)")
	flag.StringVar(&repo, "r", "", "repo (e.g. hub)")
	flag.StringVar(&milestone, "m", "", "milestone number")
	flag.BoolVar(&list, "list", false, "Print milestone list.")
	flag.StringVar(&state, "state", "all", "(optional) milestone state: 'open' or 'closed'.")
	flag.BoolVar(&ver, "v", false, "Print version.")
	flag.Parse()

	if ver {
		fmt.Fprintln(os.Stdout, "Version:", Version)
		fmt.Fprintln(os.Stdout, "Revision:", Revision)
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

	if list == true && len(state) > 0 {
		if state != "open" && state != "closed" && state != "all" {
			usage()
			os.Exit(ExitCodeOK)
		}
	}

	// Prepare config
	config, err := NewConfig(configPath)
	if err != nil {
		os.Exit(ExitCodeError)
	}

	// Prepare app
	app, err = NewApp(config, list, owner, repo, milestone, state)
	if err != nil {
		os.Exit(ExitCodeError)
	}
}

func main() {
	os.Exit(app.Run())
}
