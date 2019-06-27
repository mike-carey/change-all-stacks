package commands

import (
	"os"
	"fmt"
	"bytes"

	"github.com/mike-carey/change-all-stacks/data"
	"github.com/mike-carey/change-all-stacks/logger"
)

type queryCommand struct {
	Format string `short:"F" long:"format" description:"The format to output" choice:"csv" choice:"" default:""`
}

func (c *queryCommand) GetFormatter() data.Formatter {
	switch c.Format {
	case "csv":
		return data.NewCsvFormatter()
	default:
		return data.NewDefaultFormatter()
	}
}

type AppsWithBuildpackCommand struct {
	queryCommand
	Args struct {
	 	BuildpackName string
	} `positional-args:"true" required:"true"`
}

type AppsWithStackCommand struct {
	queryCommand
	Args struct {
		StackName string
	} `positional-args:"true" required:"true"`
}

type AppBuildpacksCommand struct {
	queryCommand
}

type QueryCommand struct {
	AppsWithBuildpackCommand `command:"apps-with-buildpack"`
	AppsWithStackCommand `command:"apps-with-stack"`
	AppBuildpacksCommand `command:"app-buildpacks"`
}

func (c *AppsWithBuildpackCommand) Execute(args []string) error {
	opts := c
	logger.Debug("Getting QueryServices")

	qs, err := manager.QueryServices()
	if err != nil {
		return err
	}

	mOpts := manager.GetOptions()
	formatter := c.GetFormatter()
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
	formatter := c.GetFormatter()
	buff := bytes.NewBuffer(nil)

	for foundationName, q := range qs {
		apps, err := q.GetAllAppsWithinOrgs(mOpts.Orgs...)
		if err != nil {
			return err
		}

		logger.Debug("Got all apps from queryService")

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

func (c *AppBuildpacksCommand) Execute(args []string) error {
	logger.Debug("Getting QueryServices")

	qs, err := manager.QueryServices()
	if err != nil {
		return err
	}

	buff := bytes.NewBuffer(nil)

	for foundationName, q := range qs {
		apps, err := q.GetAllAppsWithinOrgs()
		if err != nil {
			return err
		}

		logger.Debug("Got all apps from queryService")

		buff.WriteString(fmt.Sprintf("Foundation: %s\n", foundationName))
		bs, err := q.GetBuildpackSet(apps)
		if err != nil {
			return err
		}

		for _, b := range bs {
			buff.WriteString(fmt.Sprintf("  - %s\n", b))
		}
		buff.WriteString(fmt.Sprintf("\n"))
	}

	buff.WriteTo(os.Stdout)

	return nil
}

var _ Command = &AppsWithStackCommand{}
var _ Command = &AppsWithBuildpackCommand{}
var _ Command = &AppBuildpacksCommand{}
