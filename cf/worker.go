package cf

import (
	"github.com/mike-carey/change-all-stacks/errors"
	"github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_worker.go Worker
type Worker interface {
	Work(apps []cfclient.App, stackName string) error
}

type worker struct {
	runner Runner
	config *cfclient.Config
	pluginPath string
	orgName string
	spaceName string
}

func NewWorker(runner Runner, config *cfclient.Config, pluginPath string, orgName string, spaceName string) Worker {
	return &worker{
		runner: runner,
		config: config,
		pluginPath: pluginPath,
		orgName: orgName,
		spaceName: spaceName,
	}
}

func (w *worker) Work(apps []cfclient.App, stackName string) error {
	err := w.runner.Setup(w.config, w.pluginPath, w.orgName, w.spaceName)
	if err != nil {
		return err
	}

	errCh := make(chan error, len(apps))
	for _, app := range apps {
		go func(app cfclient.App) {
			errCh <- w.runner.Run(app.Name, stackName)
		}(app)
	}

	errPool := make([]error, 0)
	for _, _ = range apps {
		if err = <-errCh; err != nil {
			errPool = append(errPool, err)
		}
	}

	return errors.NewErrorStack("Unable to properly run for apps", errPool)
}
