package services

import (
	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/mike-carey/change-all-stacks/query"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_query_service QueryService
type QueryService interface {
	GetAllAppsWithinOrgs(orgNames ...string) ([]cfclient.App, error)
	FilterAppsByStackName(apps []cfclient.App, stackName string) ([]cfclient.App, error)
}

func NewQueryService(inquisitor query.Inquisitor) QueryService {
	return &queryService{
		inquisitor: inquisitor,
	}
}

type queryService struct {
	inquisitor query.Inquisitor
}

func (s *queryService) GetAllAppsWithinOrgs(orgNames ...string) ([]cfclient.App, error) {
	logger.Debug("Getting all apps")
	apps, err := s.inquisitor.GetAllApps()
	if err != nil {
		return nil, err
	}

	if len(orgNames) > 0 {
		logger.Debugf("Filter apps by %d orgs", len(orgNames))
		orgs := make([]cfclient.Org, len(orgNames))
		for i, orgName := range orgNames {
			logger.Debugf("Finding org by name: %s", orgName)
			org, err := s.inquisitor.GetOrgByName(orgName)
			if err != nil {
				return nil, err
			}

			orgs[i] = org
		}

		logger.Debugf("Performing filter on %d apps", len(apps))
		apps, err = query.AppFilterBy(apps, func (app cfclient.App) (bool, error) {
			logger.Debugf("Grabbing the space for app(%s)", app.Name)
			space, err := s.inquisitor.GetSpaceByGuid(app.SpaceGuid)
			if err != nil {
				return false, err
			}

			logger.Debugf("Checking if app(%s) belongs to any of the %d orgs in filter", app.Name, len(orgs))
			for _, org := range orgs {
				if space.OrganizationGuid == org.Guid {
					logger.Debugf("App(%s) is part of the orgs in the filter", app.Name)
					return true, nil
				}
			}

			logger.Debugf("App(%s) is not in any of the orgs in the filter", app.Name)
			return false, nil
		})

		if err != nil {
			return nil, err
		}
	}

	return apps, nil
}

func (s *queryService) FilterAppsByStackName(apps []cfclient.App, stackName string) ([]cfclient.App, error) {
	logger.Debugf("Grabbing the stack by name: %s", stackName)
	stack, err := s.inquisitor.GetStackByName(stackName)
	if err != nil {
		return nil, err
	}

	logger.Debugf("Filtering out apps with stack(guid:%s, name:%s)", stack.Guid, stack.Name)
	apps, err = query.AppFilterBy(apps, func (app cfclient.App) (bool, error) {
		match := app.StackGuid == stack.Guid
		if match {
			logger.Debugf("App(%s) has stack(%s)", app.Name, stack.Name)
		} else {
			logger.Debugf("App(%s) does not have stack(%s)", app.Name, stack.Name)
		}

		return match, nil
	})

	return apps, err
}
