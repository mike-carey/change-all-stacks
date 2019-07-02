package services

import (
	"github.com/mike-carey/change-all-stacks/data"
	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/mike-carey/change-all-stacks/query"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_query_service.go QueryService
type QueryService interface {
	GetAllAppsWithinOrgs(orgNames ...string) ([]cfclient.App, error)
	FilterAppsByStackName(apps []cfclient.App, stackName string) ([]cfclient.App, error)
	FilterAppsByBuildpackName(apps []cfclient.App, buildpackName string) ([]cfclient.App, error)
	GetAppData(foundationName string, apps []cfclient.App) (data.Data, error)
	GetBuildpackSet(apps []cfclient.App) ([]string, error)
	GroupAppsByOrgAndSpace(apps []cfclient.App) (map[string]map[string][]cfclient.App, error)
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

func (s *queryService) FilterAppsByBuildpackName(apps []cfclient.App, buildpackName string) ([]cfclient.App, error) {

	// logger.Debugf("Filtering out apps with buildpack(guid:%s, name:%s)", buildpack.Guid, buildpack.Name)
	logger.Debugf("Filtering out apps with buildpack(name:%s)", buildpackName)
	apps, err := query.AppFilterBy(apps, func (app cfclient.App) (bool, error) {
		b, e := s.getBuildpackForApp(app)
		if e != nil {
			return false, e
		}

		if b != "" && b == buildpackName {
			return true, nil
		}

		return false, nil
	})

	return apps, err
}

func (s *queryService) GetAppData(foundationName string, apps []cfclient.App) (data.Data, error) {
	d := make(data.Data, len(apps))
	for i, app := range apps {
		space, err := s.inquisitor.GetSpaceByGuid(app.SpaceGuid)
		if err != nil {
			return nil, err
		}

		org, err := s.inquisitor.GetOrgByGuid(space.OrganizationGuid)
		if err != nil {
			return nil, err
		}

		d[i] = *data.NewDataEntry(foundationName, org, space, app, "", cfclient.User{})
	}

	return d, nil
}

func (s *queryService) getBuildpackForApp(app cfclient.App) (string, error) {
	if app.Buildpack != "" {
		logger.Debugf("App(%s) explicitly specified a buildpack(%s)", app.Name, app.Buildpack)
		return app.Buildpack, nil
	}

	if app.DetectedBuildpackGuid != "" {
		logger.Debugf("App(%s)'s buildpack was detected as a guid %s", app.Name, app.DetectedBuildpackGuid)
		buildpack, err := s.inquisitor.GetBuildpackByGuid(app.DetectedBuildpackGuid)
		if err != nil {
			return "", err
		}

		return buildpack.Name, nil
	}

	if app.DetectedBuildpack != "" {
		logger.Debugf("App(%s)'s buildpack was detected as %s", app.Name, app.DetectedBuildpack)
		return app.DetectedBuildpack, nil
	}

	logger.Warningf("Buildpack for app(%s) could not be determined", app.Name)

	return "", nil
}

func (s *queryService) GetBuildpackSet(apps []cfclient.App) ([]string, error) {
	bs := make([]string, 0)
	bc := make(map[string]string, 0)
	for _, app := range apps {
		b, e := s.getBuildpackForApp(app)
		if e != nil {
			return nil, e
		}

		if b == "" {
			continue
		}

		if _, ok := bc[b]; ok {
			continue
		}

		bc[b] = b

		bs = append(bs, b)
	}

	return bs, nil
}

func (s *queryService) GroupAppsByOrgAndSpace(apps []cfclient.App) (map[string]map[string][]cfclient.App, error) {
	pool := make(map[string]map[string][]cfclient.App, 0)

	for _, app := range apps {
		space, err := s.inquisitor.GetSpaceByGuid(app.SpaceGuid)
		if err != nil {
			return nil, err
		}

		org, err := s.inquisitor.GetOrgByGuid(space.OrganizationGuid)
		if err != nil {
			return nil, err
		}

		o := org.Name
		s := space.Name

		if _, ok := pool[o]; !ok {
			pool[o] = make(map[string][]cfclient.App, 0)
		}

		if _, ok := pool[o][s]; !ok {
			pool[o][s] = make([]cfclient.App, 0)
		}

		pool[o][s] = append(pool[o][s], app)
	}

	return pool, nil
}

var _ QueryService = &queryService{}
