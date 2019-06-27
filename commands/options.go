package commands

import (
	"github.com/mike-carey/change-all-stacks/services"
)

type Options struct {
	*services.ManagerOptions
	QueryCommand `command:"query" description:"Queries the cloud controller for information"`
	ChangeCommand `command:"change" description:"Change the stacks on apps from one stack to another"`
}

type InitialOptions struct {
	Verbose []bool   `short:"v" long:"verbose" description:"Prints more output"`
	Version bool   `long:"version" description:"Prints the version of the cli"`
}

// func ParseArgs(options interface{}, args []string) ([]string, error) {
// 	return flags.NewParser(options, flags.HelpFlag | flags.PrintErrors | flags.PassDoubleDash | flags.IgnoreUnknown).ParseArgs(args)
// }
