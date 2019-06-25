package change

import (
	"fmt"

	"github.com/mike-carey/change-all-stacks/cf"
	"github.com/mike-carey/change-all-stacks/logger"

	"github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_space_manager.go SpaceManager
type SpaceManager interface {
	ChangeStacksInSpace(config cfclient.Config, org cfclient.Org, space cfclient.Space, apps []cfclient.App, stack string, dryrun bool, pluginPath string) error
}

func NewDefaultSpaceManager(logger logger.Logger) SpaceManager {
	return NewSpaceManager(logger, cf.NewRunnerFactory())
}

func NewSpaceManager(logger logger.Logger, runnerFactory cf.RunnerFactory) SpaceManager {
	return &spaceManager{
		logger: logger,
		runnerFactory: runnerFactory,
	}
}

type spaceManager struct {
	logger logger.Logger
	runnerFactory cf.RunnerFactory
}

func (m *spaceManager) ChangeStacksInSpace(config cfclient.Config, org cfclient.Org, space cfclient.Space, apps []cfclient.App, stack string, dryrun bool, pluginPath string) error {
	r := m.runnerFactory.CreateRunnerWithDefaultCommand(m.logger, dryrun)

	var err error

	m.logger.Debugf("Setting up runner environment for org: %s and space: %s", org.Name, space.Name)
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
	err = r.InstallPlugin(pluginPath)
	if err != nil {
		return err
	}
	m.logger.Debugf("Finished setting up runner environment for org: %s and space: %s", org.Name, space.Name)

	errCh := make(chan error)
	for _, app := range apps {
		go func(runner cf.Runner, app string, stack string) {
			errCh <- runner.ChangeStack(app, stack)
		}(r, app.Name, stack)
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
