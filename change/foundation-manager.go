package change

import (
	"fmt"

	"github.com/mike-carey/cfquery/query"

	"github.com/cloudfoundry-community/go-cfclient"
)

type FoundationManager interface {
	ChangeStacksInFoundation(foundationName string, config cfclient.Config, fromStack string, toStack string, dryrun bool, pluginPath string) error
}

func NewDefaultFoundationManager(logger Logger) FoundationManager {
	return NewFoundationManager(logger, NewDefaultSpaceManager(logger), NewInquisitorFactory())
}

func NewFoundationManager(logger Logger, spaceManager SpaceManager, inquisitorFactory InquisitorFactory) FoundationManager {
	return &foundationManager{
		logger: logger,
		spaceManager: spaceManager,
		inquisitorFactory: inquisitorFactory,
	}
}

type foundationManager struct {
	logger Logger
	spaceManager SpaceManager
	inquisitorFactory InquisitorFactory
}

func (m *foundationManager) ChangeStacksInFoundation(foundationName string, config cfclient.Config, fromStack string, toStack string, dryrun bool, pluginPath string) error {
	i, err := m.inquisitorFactory.CreateInquisitor(config)
	if err != nil {
		return err
	}

	// Grab every app
	m.logger.Debugf("Loading all apps in %s", foundationName)
	apps, err := i.GetAllApps()
	if err != nil {
		return err
	}

	m.logger.Debugf("Filtering %i apps by stackname '%s' in %s", len(apps), fromStack, foundationName)
	apps, err = apps.FilterByStackName(i, fromStack)
	if err != nil {
		return err
	}

	// Group by space
	m.logger.Debugf("Grouping %i apps by space in %s", len(apps), foundationName)
	mapps, err := apps.GroupBySpace(i)
	if err != nil {
		return err
	}

	m.logger.Debugf("Changing stacks from '%s' to '%s' for %i apps in %s", fromStack, toStack, len(mapps), foundationName)
	errCh := make(chan error)
	for spaceGuid, apps := range mapps {
		go func(i query.Inquisitor, config cfclient.Config, spaceGuid string, apps query.Apps) {
			m.logger.Debugf("Getting space from guid")
			space, err := i.GetSpaceByGuid(spaceGuid)
			if err != nil {
				errCh <- err
				return
			}

			m.logger.Debugf("Getting org from guid")
			org, err := i.GetOrgByGuid(space.OrganizationGuid)
			if err != nil {
				errCh <- err
				return
			}

			errCh <- m.spaceManager.ChangeStacksInSpace(config, *org, *space, apps, toStack, dryrun, pluginPath)
		}(i, config, spaceGuid, apps)
	}


	errPool := make([]error, 0)
	for _, _ = range mapps {
		if err := <-errCh; err != nil {
			errPool = append(errPool, err)
		}
	}

	if len(errPool) > 0 {
		return NewErrorStack(fmt.Sprintf("Failed to change stacks in foundation: %s", foundationName), errPool)
	}

	return nil
}
