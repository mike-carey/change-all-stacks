package main

import (
	"os"
	"fmt"
	"strconv"

	"github.com/jessevdk/go-flags"

	. "github.com/mike-carey/change-all-stacks/change"
)

const (
	EnvConfig = "CHANGE_ALL_STACKS_CONFIG"
	EnvDryRun = "CHANGE_ALL_STACKS_DRY_RUN"
	EnvVerbose = "CHANGE_ALL_STACKS_VERBOSE"
)

func onError(err error) {
	fmt.Println("Writting error")
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func options() (*Options, error) {
	opts := &Options{}

	if c := os.Getenv(EnvConfig); c != "" {
		opts.Config = c
	}

	if v := os.Getenv(EnvVerbose); v != "" {
		b, e := strconv.ParseBool(v)
		if e != nil {
			return nil, e
		}
		opts.Verbose = b
	}

	if d := os.Getenv(EnvDryRun); d != "" {
		b, e := strconv.ParseBool(d)
		if e != nil {
			return nil, e
		}
		opts.DryRun = b
	}

	return opts, nil
}

type VersionOption struct {
	Version bool   `long:"version" description:"Prints the version of the cli"`
}

func main() {
	vOpt := VersionOption{}
	args, err := flags.NewParser(&vOpt, flags.PassDoubleDash).Parse()
	if vOpt.Version {
		fmt.Println(Version)
		return
	}

	opts, err := options()
	if err != nil {
		onError(err)
		return
	}

	parser := flags.NewParser(opts, flags.Default)

	_, err = parser.ParseArgs(args)
	if err != nil || flags.WroteHelp(err) {
		return
	}

	err = Go(opts)
	if err != nil {
		onError(err)
		return
	}
}
