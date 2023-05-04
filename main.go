package main

import (
	"os"

	"tlsg10x-cli/cmd"
)

var (
	version = "dev"
	// https://goreleaser.com/cookbooks/using-main.version
	// commit  = "none"
	// date    = "unknown"
	// builtBy = "unknown"
)

func main() {
	root := cmd.RootCmd(version)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
