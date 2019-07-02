package cf

import (
	"fmt"
	"bytes"
	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_runner.go Runner
type Runner interface {
	Setup(c *cfclient.Config, pluginPath string, org string, space string) error
	Run(appName string, stackName string) (*bytes.Buffer, error)
}

type runner struct {
	executor Executor
	isSetup bool
}

func NewRunner(executor Executor) Runner {
	return &runner{
		executor: executor,
		isSetup: false,
	}
}

func (r *runner) Setup(c *cfclient.Config, pluginPath string, org string, space string) error {
	var err error

	withOrWithout := "with"
	if c.SkipSslValidation {
		withOrWithout = "without"
	}

	logger.Debugf("Setting api endpoint to '%s' %s ssl validation", c.ApiAddress, withOrWithout)
	err = r.executor.Api(c.ApiAddress, c.SkipSslValidation)
	if err != nil {
		return err
	}

	logger.Debugf("Authorizing via %s:<REDACTED>", c.Username)
	err = r.executor.Auth(c.Username, c.Password)
	if err != nil {
		return err
	}

	logger.Debugf("Targeting org: %s and space: %s", org, space)
	err = r.executor.Target(org, space)
	if err != nil {
		return err
	}

	logger.Debugf("Installing plugin: %s", pluginPath)
	err = r.executor.InstallPlugin(pluginPath)
	if err != nil {
		return err
	}

	r.isSetup = true

	return nil

}

func (r *runner) Run(appName string, stackName string) (*bytes.Buffer, error) {
	if !r.isSetup {
		return nil, fmt.Errorf("Runner is not setup!")
	}

	logger.Debugf("Changing %s's stack to %s", appName, stackName)
	err := r.executor.ChangeStack(appName, stackName)
	if err != nil {
		logger.Debugf("Failed to change %s's stack to %s.  Reason: %v", err)
		return nil, err
	}

	logger.Debugf("Successfully changed %s's stack to %s", appName, stackName)
	return r.executor.Buffer(), nil
}
