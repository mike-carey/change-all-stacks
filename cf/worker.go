package cf

import (
	"bytes"
	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/mike-carey/change-all-stacks/errors"
	"github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_worker.go Worker
type Worker interface {
	Work(name string, apps []cfclient.App, stackName string) (*bytes.Buffer, error)
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

func (w *worker) Work(name string, apps []cfclient.App, stackName string) (*bytes.Buffer, error) {
	logger.Debugf("%s Worker: Setting up", name)
	err := w.runner.Setup(w.config, w.pluginPath, w.orgName, w.spaceName)
	if err != nil {
		return nil, err
	}

	logger.Debugf("%s Worker: Running for %d apps", name, len(apps))
	buffCh := make(chan *bytes.Buffer, 0)
	errCh := make(chan error, 0)
	for _, app := range apps {
		go func(name string, app cfclient.App) {
			logger.Debugf("%s Worker: Running to change %s to %s", name, app.Name, stackName)
			b, e := w.runner.Run(app.Name, stackName)
			if e != nil {
				logger.Debugf("%s Worker: Error running for %s", name, app.Name)
				errCh <- e
			} else {
				logger.Debugf("%s Worker: Success running for %s", name, app.Name)
				buffCh <- b
			}
		}(name, app)
	}

	buffPool := bytes.NewBuffer(nil)
	errPool := make([]error, 0)
	for _, _ = range apps {
		select {
		case err = <-errCh:
			if err != nil {
				errPool = append(errPool, err)
			}
		case buff := <-buffCh:
			buff.WriteTo(buffPool)
		}
	}
	logger.Debugf("%s Worker: Done, releasing buffer", name)

	if len(errPool) > 0 {
		return nil, errors.NewErrorStack("Unable to properly run for apps", errPool)
	}

	return buffPool, nil
}
