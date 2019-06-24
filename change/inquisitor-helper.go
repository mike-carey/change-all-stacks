package change

import (
	"github.com/mike-carey/change-all-stacks/query"

	"github.com/cloudfoundry-community/go-cfclient"
)

type InquisitorHelper interface {
	GetAllApps(orgGuids ...string) ([]cfclient.App, error)
	GroupAppsByOrgAndSpace(apps []cfclient.App) (map[string]map[string][]cfclient.App, error)
	FilterAppsByStackName(apps []cfclient.App, stackName string) ([]cfclient.App, error)
	GetOrgAndSpaceForSpaceGuid(spaceGuid string) (cfclient.Org, cfclient.Space, error)
}

type inquisitorHelper struct {
	inquisitor query.Inquisitor
}

func NewInquisitorHelper(inquisitor query.Inquisitor) InquisitorHelper {
	return &inquisitorHelper{
		inquisitor: inquisitor,
	}
}

func (h *inquisitorHelper) GetAllApps(orgGuids ...string) ([]cfclient.App, error) {
	apps, err := h.inquisitor.GetAllApps()
	if err != nil {
		return nil, err
	}

	if len(orgGuids) > 0 {
		apps, err = query.AppFilterBy(apps, func (app cfclient.App) (bool, error) {
			space, err := h.inquisitor.GetSpaceByGuid(app.SpaceGuid)
			if err != nil {
				return false, err
			}

			for _, orgGuid := range orgGuids {
				if orgGuid == space.OrganizationGuid {
					return true, nil
				}
			}

			return false, nil
		})
	}

	return apps, err
}

func (h *inquisitorHelper) GroupAppsByOrgAndSpace(apps []cfclient.App) (map[string]map[string][]cfclient.App, error) {
	mapps, err := query.AppGroupBy(apps, func(app cfclient.App) (string, error) {
		return app.SpaceGuid, nil
	})

	if err != nil {
		return nil, err
	}

	pool := make(map[string]map[string][]cfclient.App)
	for spaceGuid, apps := range mapps {
		space, err := h.inquisitor.GetSpaceByGuid(spaceGuid)
		if err != nil {
			return nil, err
		}

		org, err := h.inquisitor.GetOrgByGuid(space.OrganizationGuid)
		if err != nil {
			return nil, err
		}

		if _, ok := pool[org.Guid]; !ok {
			pool[org.Guid] = make(map[string][]cfclient.App, 0)
		}

		pool[org.Guid][spaceGuid] = apps
	}

	return pool, err
}

func (h *inquisitorHelper) FilterAppsByStackName(apps []cfclient.App, stackName string) ([]cfclient.App, error) {
	stack, err := h.inquisitor.GetStackByName(stackName)
	if err != nil {
		return nil, err
	}

	return query.AppFilterBy(apps, func(app cfclient.App) (bool, error) {
		return app.StackGuid == stack.Guid, nil
	})
}

func (h *inquisitorHelper) GetOrgAndSpaceForSpaceGuid(spaceGuid string) (cfclient.Org, cfclient.Space, error) {
	space, err := h.inquisitor.GetSpaceByGuid(spaceGuid)
	if err != nil {
		return cfclient.Org{}, cfclient.Space{}, err
	}

	org, err := h.inquisitor.GetOrgByGuid(space.OrganizationGuid)
	if err != nil {
		return cfclient.Org{}, cfclient.Space{}, err
	}

	return org, space, nil
}
