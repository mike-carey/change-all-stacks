package commands

import (
	"io"
	"fmt"

	"github.com/jessevdk/go-flags"

	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/mike-carey/change-all-stacks/services"
)

type Command interface {
	Execute([]string) error
}

type Commander struct {
	managerService services.ManagerService
}

func NewDefaultCommander() *Commander {
	return NewCommander(services.NewManagerService())
}

func NewCommander(managerService services.ManagerService) *Commander {
	return &Commander{
		managerService: managerService,
	}
}

var cOpts *Options

func (c *Commander) Go(args []string, stdout io.Writer, stderr io.Writer) error {
	vOpts := InitialOptions{}
	args, err := flags.NewParser(&vOpts, flags.PassDoubleDash | flags.IgnoreUnknown).ParseArgs(args)
	if vOpts.Version {
		fmt.Fprintln(stdout, Version)
		return nil
	}

	v := len(vOpts.Verbose) > 0
	d := len(vOpts.Verbose) > 1

	// Initialize the logger
	logger.Init(&v, &d)
	logger.Debug("Intialized logger")

	mOpts := getDefaultManagerOptions()

	args, err = flags.NewParser(mOpts, flags.PassDoubleDash | flags.IgnoreUnknown).ParseArgs(args)
	if err != nil {
		return err
	}

	// Translate options
	opts := TranslateManagerOptions(mOpts)

	logger.Debugf("Creating manager: %v", opts)
	manager, err = c.managerService.CreateManager(opts)
	if err != nil {
		return err
	}

	cOpts = &Options{}

	// Parse the leftover args
	_, err = flags.NewParser(cOpts, flags.Default).ParseArgs(args)
	if err != nil || flags.WroteHelp(err) {
		return nil
	}

	return err
}
