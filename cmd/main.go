package main

import (
	"os"
	"fmt"

	. "github.com/mike-carey/change-all-stacks/commands"
)

const (
	EnvConfig = "CHANGE_ALL_STACKS_CONFIG"
	EnvDryRun = "CHANGE_ALL_STACKS_DRY_RUN"
	EnvVerbose = "CHANGE_ALL_STACKS_VERBOSE"
)

func main() {
	err := NewDefaultCommander().Go(os.Args[1:], os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
