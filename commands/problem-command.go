package commands

import (
	"fmt"
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

	for foundationName, qs := range qss {
		apps, err := qs.GetAllAppsWithinOrgs(mOpts.Orgs...)
		if err != nil {
			return err
		}

		fmt.Printf("Foundation: %s", foundationName)

		ps := pss[foundationName]
		problems, err := ps.FindProblems(apps)
		if err != nil {
			return err
		}

		for _, p := range problems {
			fmt.Printf("- %s", p.GetReason())
		}
	}

	return nil
}

var _ Command = &ProblemCommand{}
