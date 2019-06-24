package change

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/change-all-stacks/data"
)

type FoundationManager interface {
	ReadAllStacksInFoundation(foundationName string, config *cfclient.Config, opts *Options) (data.Data, error)
	ChangeStacksInFoundation(foundationName string, config *cfclient.Config, opts *Options) error
}

func NewDefaultFoundationManager(logger Logger) FoundationManager {
	return NewFoundationManager(logger, NewDefaultSpaceManager(logger), NewInquisitorManager())
}

func NewFoundationManager(logger Logger, spaceManager SpaceManager, inquisitorManager InquisitorManager) FoundationManager {
	return &foundationManager{
		logger: logger,
		spaceManager: spaceManager,
		inquisitorManager: inquisitorManager,
	}
}

type foundationManager struct {
	logger Logger
	spaceManager SpaceManager
	inquisitorManager InquisitorManager
}

func (m *foundationManager) getApps(foundationName string, config *cfclient.Config, opts *Options) (map[string]map[string][]cfclient.App, error) {
	fromStack := opts.Stacks.From
	// toStack := opts.Stacks.To
	// threads := opts.Threads
	orgs := opts.Orgs

	// Build the inquisitorhelper
	m.logger.Debugf("Getting inquisitor helper")
	helper, err := m.inquisitorManager.GetHelper(config)
	if err != nil {
		return nil, err
	}

	// Grab every app
	m.logger.Debugf("Loading all apps in %s filtered on orgs: %v", foundationName, orgs)
	apps, err := helper.GetAllApps(orgs...)
	if err != nil {
		return nil, err
	}

	m.logger.Debugf("Filtering %i apps by stackname '%s' in %s", len(apps), fromStack, foundationName)
	apps, err = helper.FilterAppsByStackName(apps, fromStack)
	if err != nil {
		return nil, err
	}

	// Group by space
	m.logger.Debugf("Grouping %i apps by org and space in %s", len(apps), foundationName)
	return helper.GroupAppsByOrgAndSpace(apps)
}

func (m *foundationManager) ReadAllStacksInFoundation(foundationName string, config *cfclient.Config, opts *Options) (data.Data, error) {
	m.logger.Debugf("Getting all apps for %s", foundationName)
	mapps, err := m.getApps(foundationName, config, opts)
	if err != nil {
		return nil, err
	}

	m.logger.Debugf("Getting inquisitor helper for %s", foundationName)
	helper, err := m.inquisitorManager.GetHelper(config)
	if err != nil {
		return nil, fmt.Errorf("Error getting helper: %s", foundationName)
		return nil, err
	}

	m.logger.Debugf("Building a data pool for %d orgs", len(mapps))
	dataPool := make(data.Data, 0)
	for _, gapps := range mapps {
		for spaceGuid, apps := range gapps {
			m.logger.Debugf("Retrieving org and space for space-guid: %s", spaceGuid)
			org, space, err := helper.GetOrgAndSpaceForSpaceGuid(spaceGuid)
			if err != nil {
				return nil, fmt.Errorf("Error getting org and space for space-guid: %s", spaceGuid)
			}

			m.logger.Debugf("Adding %d data entries to the data pool", len(apps))
			for _, app := range apps {
				de := data.NewDataEntry(foundationName, org, space, app, "", cfclient.User{})
				dataPool = append(dataPool, *de)
			}
		}
	}

	m.logger.Debugf("Returning a data pool of size: %d", len(dataPool))
	return dataPool, nil
}

func (m *foundationManager) ChangeStacksInFoundation(foundationName string, config *cfclient.Config, opts *Options) error {
	fromStack := opts.Stacks.From
	toStack := opts.Stacks.To
	dryrun := opts.DryRun
	pluginPath := opts.Plugin
	threads := opts.Threads

	mapps, err := m.getApps(foundationName, config, opts)
	if err != nil {
		return err
	}

	helper, err := m.inquisitorManager.GetHelper(config)
	if err != nil {
		return err
	}

	m.logger.Debugf("Changing stacks from '%s' to '%s' for %i apps in %s", fromStack, toStack, len(mapps), foundationName)
	errCh := make(chan error)

	var sem = make(chan int, threads)
	for _, gapps := range mapps {
		for spaceGuid, apps := range gapps {
			go func(helper InquisitorHelper, config cfclient.Config, spaceGuid string, apps []cfclient.App) {
				sem <- 1
				m.logger.Debugf("Acquired lock")

				defer func() {
					<-sem
					m.logger.Debugf("Released lock")
				}()

				m.logger.Debugf("Getting org and space")
				org, space, err := helper.GetOrgAndSpaceForSpaceGuid(spaceGuid)
				if err != nil {
					errCh <- err
					return
				}

				errCh <- m.spaceManager.ChangeStacksInSpace(config, org, space, apps, toStack, dryrun, pluginPath)
			}(helper, *config, spaceGuid, apps)
		}
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
