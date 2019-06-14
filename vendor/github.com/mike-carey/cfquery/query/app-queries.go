package query

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

func (g Apps) FilterByStackName(i Inquisitor, stackName string) (Apps, error) {
	stack, err := i.GetStackByName(stackName)
	if err != nil {
		return nil, err
	}

	stackGuid := stack.Guid

	return AppFilterBy(g, func(app cfclient.App) (bool, error) {
		return app.StackGuid == stackGuid, nil
	})
}

func (g Apps) GroupBySpace(_ Inquisitor) (AppGroup, error) {
	return AppGroupBy(g, func(app cfclient.App) (string, error) {
		return app.SpaceGuid, nil
	})
}

func (g Apps) GroupByOrg(i Inquisitor) (AppGroup, error) {
	return AppGroupBy(g, func(app cfclient.App) (string, error) {
		s, e := i.GetSpaceByGuid(app.SpaceGuid)
		if e != nil {
			return "", e
		}

		return s.OrganizationGuid, nil
	})
}

// GroupBySpaceAndOrg ...
func (g Apps) GroupBySpaceAndOrg(i Inquisitor) (MappedAppGroup, error) {
	ag, err := g.GroupBySpace(i)
	if err != nil {
		return nil, err
	}

	return ag.GroupByOrg(i)
}

func (g AppGroup) GroupByOrg(i Inquisitor) (MappedAppGroup, error) {
	return AppGroupMappedSliceBy(g, func(_ string, apps Apps) (string, error) {
		s, e := i.GetSpaceByGuid(apps[0].SpaceGuid)
		if e != nil {
			return "", e
		}

		return s.OrganizationGuid, nil
	})
}
