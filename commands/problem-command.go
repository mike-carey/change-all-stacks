package commands

import (
	"os"
	"bytes"
	"github.com/mike-carey/change-all-stacks/data"
)

type ProblemCommand struct {
	Format string `short:"F" long:"format" description:"The format to output" choice:"csv" choice:"" default:""`
	Stacks struct {
		From string
		To string
	} `positional-args:"yes" required:"yes"`

}

func (c *ProblemCommand) GetFormatter() data.Formatter {
	switch c.Format {
	case "csv":
		return data.NewCsvFormatter()
	default:
		return data.NewDefaultFormatter()
	}
}

func (c *ProblemCommand) Execute([]string) error {
	pss, err := manager.ProblemServices()
	if err != nil {
		return err
	}

	qss, err := manager.QueryServices()
	if err != nil {
		return err
	}

	mOpts := manager.GetOptions()
	formatter := c.GetFormatter()
	buffer := bytes.NewBuffer(nil)

	for foundationName, qs := range qss {
		apps, err := qs.GetAllAppsWithinOrgs(mOpts.Orgs...)
		if err != nil {
			return err
		}

		apps, err = qs.FilterAppsByStackName(apps, c.Stacks.From)
		if err != nil {
			return err
		}

		ps := pss[foundationName]
		problems, err := ps.FindProblems(foundationName, apps, c.Stacks.From, c.Stacks.To)
		if err != nil {
			return err
		}

		s, err := formatter.FormatProblemSet(problems)
		if err != nil {
			return err
		}

		buffer.WriteString(s)
	}

	buffer.WriteTo(os.Stdout)

	return nil
}

var _ Command = &ProblemCommand{}
