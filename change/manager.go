package change

import (
	"fmt"
	"os"
	"bufio"

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
	m.logger.Debugf("Config loaded")

	if m.Options.Interactive {
		m.logger.Info("Running dry run")
		m.logger.Info("-----------------------------------")

		m.Options.DryRun = true

		err = m.ChangeAllStacks(foundations)
		if err != nil {
			return err
		}

		reader := bufio.NewReader(os.Stdin)

		m.logger.Info("-----------------------------------")

		text := ""
		for text == "" {
			fmt.Print("Proceed? [Y/n]: ")
			text, _ = reader.ReadString('\n')

			switch text[0:1] {
			case "":
				continue
			case "y":
				break
			case "n":
				m.logger.Info("Bailing out.")
				return nil
			default:
				m.logger.Info("Please specify 'y' or 'n'.")
				continue
			}
		}

		m.Options.DryRun = false

		m.logger.Info("-----------------------------------")
	}

	m.logger.Debugf("Changing all stacks in all foundations")
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
		go func(name string, conf *cfclient.Config) {
			errCh <- m.foundationManager.ChangeStacksInFoundation(name, conf, m.Options)
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
