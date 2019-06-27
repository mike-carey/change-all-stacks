package commands

import (
	"bytes"
	"os"
	"github.com/mike-carey/change-all-stacks/logger"
)

type AppsWithBuildpackCommand struct {
	Args struct {
	 	BuildpackName string
	} `positional-args:"true" required:"true"`
}

type AppsWithStackCommand struct {
	Args struct {
		StackName string
	} `positional-args:"true" required:"true"`
}

type QueryCommand struct {
	AppsWithBuildpackCommand `command:"apps-with-buildpack"`
	AppsWithStackCommand `command:"apps-with-stack"`
}

func (c *AppsWithBuildpackCommand) Execute(args []string) error {
	opts := c
	logger.Debug("Getting QueryServices")

	qs, err := manager.QueryServices()
	if err != nil {
		return err
	}

	mOpts := manager.GetOptions()
	formatter := manager.GetFormatter()
	buff := bytes.NewBuffer(nil)

	for foundationName, q := range qs {
		apps, err := q.GetAllAppsWithinOrgs(mOpts.Orgs...)
		if err != nil {
			return err
		}

		logger.Debug("Got all apps from queryService")

		logger.Debugf("Filtering apps by buildpack: %s", opts.Args.BuildpackName)
		apps, err = q.FilterAppsByBuildpackName(apps, opts.Args.BuildpackName)
		if err != nil {
			return err
		}

		data, err := q.GetAppData(foundationName, apps)
		if err != nil {
			return err
		}

		fs, err := formatter.Format(data)
		if err != nil {
			return err
		}

		buff.WriteString(fs)
	}

	buff.WriteTo(os.Stdout)

	return nil
}

func (c *AppsWithStackCommand) Execute(args []string) error {
	opts := c
	logger.Debug("Getting QueryServices")

	qs, err := manager.QueryServices()
	if err != nil {
		return err
	}

	mOpts := manager.GetOptions()
	formatter := manager.GetFormatter()
	buff := bytes.NewBuffer(nil)

	for foundationName, q := range qs {
		apps, err := q.GetAllAppsWithinOrgs(mOpts.Orgs...)
		if err != nil {
			return err
		}

		logger.Debug("Got all apps from queryService")

		logger.Debugf("%v", c)

		logger.Debugf("Filtering apps by stack: %s", opts.Args.StackName)
		apps, err = q.FilterAppsByStackName(apps, opts.Args.StackName)
		if err != nil {
			return err
		}

		data, err := q.GetAppData(foundationName, apps)
		if err != nil {
			return err
		}

		fs, err := formatter.Format(data)
		if err != nil {
			return err
		}

		buff.WriteString(fs)
	}

	buff.WriteTo(os.Stdout)

	return nil
}

var _ Command = &AppsWithStackCommand{}
var _ Command = &AppsWithBuildpackCommand{}
