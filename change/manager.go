package change

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"

	"github.com/mike-carey/cfquery/config"
)

type Manager struct {
	Options *Options
	logger Logger
	foundationManager FoundationManager
}

func NewDefaultManager(opts *Options) *Manager {
	logger := NewLogger(opts.Verbose)

	return NewManager(opts, logger, NewDefaultFoundationManager(logger))
}

func NewManager(opts *Options, logger Logger, foundationManager FoundationManager) *Manager {
	return &Manager{
		Options: opts,
		logger: logger,
		foundationManager: foundationManager,
	}
}

func (m *Manager) Go() error {
	m.logger.Debugf("Loading config: %s", m.Options.Config)

	foundations, err := config.LoadConfig(m.Options.Config)
	if err != nil {
		return err
	}

	err = m.ChangeAllStacks(foundations)
	if err != nil {
		return err
	}

	m.logger.Debug("All done!")

	return nil
}

func (m *Manager) ChangeAllStacks(foundations config.Foundations) error {
	errCh := make(chan error, 0)

	for name, conf := range foundations {
		go func(name string, conf cfclient.Config) {
			errCh <- m.foundationManager.ChangeStacksInFoundation(name, conf, m.Options.Stacks.From, m.Options.Stacks.To, m.Options.DryRun, m.Options.Plugin)
		}(name, *conf)
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
