package change

import (
	"fmt"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

	"github.com/mike-carey/cfquery/config"
	"github.com/mike-carey/cfquery/query"
)

type Manager struct {
	Options *Options
	Logger Logger
	RunnerFactory RunnerFactory
}

func NewDefaultManager(opts *Options) *Manager {
	return &Manager{
		Options: opts,
		Logger: NewLogger(opts.Verbose),
		RunnerFactory: NewRunnerFactory(),
	}
}

func NewManager(opts *Options, logger Logger, runnerFactory RunnerFactory) *Manager {
	return &Manager{
		Options: opts,
		Logger: logger,
		RunnerFactory: runnerFactory,
	}
}

func (m *Manager) Go() error {
	m.Logger.Debugf("Loading config: %s", m.Options.Config)

	foundations, err := config.LoadConfig(m.Options.Config)
	if err != nil {
		return err
	}

	err = m.ChangeAllStacks(foundations)
	if err != nil {
		return err
	}

	m.Logger.Debug("All done!")

	return nil
}

func (m *Manager) ChangeAllStacks(foundations config.Foundations) error {
	errCh := make(chan error, 0)

	for name, conf := range foundations {
		go func(name string, conf *cfclient.Config) {
			errCh <- m.ChangeAllStacksInFoundation(name, conf)
		}(name, conf)
	}

	errPool := make([]error, 0)
	for _, _ = range foundations {
		if err := <-errCh; err != nil {
			errPool = append(errPool, err)
		}
	}

	if len(errPool) > 0 {
		return NewErrorStack("Failed to change all stacks", errPool)
	}

	return nil
}

func (m *Manager) ChangeAllStacksInFoundation(foundationName string, conf *cfclient.Config) error {
	m.Logger.Debugf("Checking stacks in: %s", foundationName)

	cli, err := cfclient.NewClient(conf)
	if err != nil {
		return err
	}

	m.Logger.Debugf("Creating inquisitor for: %s", foundationName)
	i := query.NewInquisitor(cli)

	// Grab every app
	m.Logger.Debugf("Loading all apps in %s", foundationName)
	apps, err := i.GetAllApps()
	if err != nil {
		return err
	}

	m.Logger.Debugf("Filtering %i apps by stackname '%s' in %s", len(apps), m.Options.Stacks.From, foundationName)
	apps, err = apps.FilterByStackName(i, m.Options.Stacks.From)
	if err != nil {
		return err
	}

	// Group by space
	m.Logger.Debugf("Grouping %i apps by space in %s", len(apps), foundationName)
	mapps, err := apps.GroupBySpace(i)
	if err != nil {
		return err
	}

	m.Logger.Debugf("Changing stacks from '%s' to '%s' for %i apps in %s", m.Options.Stacks.From, m.Options.Stacks.To, len(mapps), foundationName)
	errCh := make(chan error)
	for spaceGuid, apps := range mapps {
		go func(conf *cfclient.Config, i query.Inquisitor, spaceGuid string, apps query.Apps) {
			errCh <- m.ChangeStacksInSpace(conf, i, spaceGuid, apps)
		}(conf, i, spaceGuid, apps)
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

func (m *Manager) ChangeStacksInSpace(config *cfclient.Config, i query.Inquisitor, spaceGuid string, apps query.Apps) error {
	m.Logger.Debugf("Getting space from guid")
	space, err := i.GetSpaceByGuid(spaceGuid)
	if err != nil {
		return err
	}

	m.Logger.Debugf("Getting org from guid")
	org, err := i.GetOrgByGuid(space.OrganizationGuid)
	if err != nil {
		return err
	}

	r := m.RunnerFactory.CreateRunnerWithDefaultCommand(m.Logger, m.Options.DryRun)
	m.Logger.Debugf("Setting up runner environment for org: %s and space: %s", org.Name, space.Name)
	err = r.Api(config.ApiAddress, config.SkipSslValidation)
	if err != nil {
		return err
	}
	err = r.Auth(config.Username, config.Password)
	if err != nil {
		return err
	}
	err = r.Target(org.Name, space.Name)
	if err != nil {
		return err
	}
	err = r.InstallPlugin(m.Options.Plugin)
	if err != nil {
		return err
	}
	m.Logger.Debugf("Finished setting up runner environment for org: %s and space: %s", org.Name, space.Name)

	errCh := make(chan error)
	for _, app := range apps {
		go func(runner Runner, app string, stack string) {
			errCh <- runner.ChangeStack(app, stack)
		}(r, app.Name, m.Options.Stacks.To)
	}

	errPool := make([]error, 0)
	for _, _ = range apps {
		if err := <-errCh; err != nil {
			errPool = append(errPool, err)
		}
	}

	if len(errPool) > 0 {
		return NewErrorStack(fmt.Sprintf("Failed to change stacks in space: %s", space.Name), errPool)
	}

	return nil
}
