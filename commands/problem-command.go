package commands

import (
	"os"
	"fmt"
	"bytes"
)

type ProblemCommand struct {}

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
	buff := bytes.NewBuffer(nil)

	for foundationName, qs := range qss {
		apps, err := qs.GetAllAppsWithinOrgs(mOpts.Orgs...)
		if err != nil {
			return err
		}

		buff.WriteString(fmt.Sprintf("Foundation: %s\n", foundationName))

		ps := pss[foundationName]
		problems, err := ps.FindProblems(apps)
		if err != nil {
			return err
		}

		for _, p := range problems {
			buff.WriteString(fmt.Sprintf("- %s\n", p.GetReason()))
		}
	}

	buff.WriteTo(os.Stdout)

	return nil
}

var _ Command = &ProblemCommand{}
