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

var app App

func init() {
	var configPath string
	var milestone string
	var repo string
	var version bool

	flag.StringVar(&configPath, "c", "", "/path/to/config.json. (default: $HOME/.config/prnotify/config.json)")
	flag.StringVar(&milestone, "m", "", "milestone number")
	flag.StringVar(&repo, "r", "", "repo")
	flag.BoolVar(&version, "v", false, "Print version.")
	flag.Parse()

	if version {
		fmt.Fprintln(os.Stdout, "Version:", Version)
		fmt.Fprintln(os.Stdout, "Revision:", Revision)
		os.Exit(ExitCodeOK)
	}

	if len(milestone) == 0 || len(repo) == 0 {
		os.Exit(ExitCodeOK)
	}

	// Prepare config
	config, err := NewConfig(configPath)
	if err != nil {
		os.Exit(ExitCodeError)
	}

	// Prepare app
	app, err = NewApp(config, repo, milestone)
	if err != nil {
		os.Exit(ExitCodeError)
	}
}

func main() {
	os.Exit(app.Run())
}
