package services

import (
	"fmt"

	"github.com/mike-carey/change-all-stacks/config"
	"github.com/mike-carey/change-all-stacks/logger"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type ManagerOptions struct {
	Config  string
	DryRun  bool
	Threads int
	Foundations []string
	Orgs []string
	PluginPath string
}

type Manager interface {
	QueryServices() (map[string]QueryService, error)
	ProblemServices() (map[string]ProblemService, error)
	WorkerService() (WorkerService, error)
	GetOptions() ManagerOptions
	GetConfigs() config.Configs
}

type manager struct {
	options *ManagerOptions
	configs config.Configs
	inqServ InquisitorService
	excServ ExecutorService
	runServ RunnerService
}

func NewManager(opts *ManagerOptions, conf config.Configs, inqServ InquisitorService, excServ ExecutorService, runServ RunnerService) (Manager, error) {
	return &manager{
		options: opts,
		configs: conf,
		inqServ: inqServ,
		excServ: excServ,
		runServ: runServ,
	}, nil

}

func NewDefaultManager(opts *ManagerOptions) (Manager, error) {
	// Grab our configurations
	configs, err := config.LoadConfigFromFile(opts.Config)
	if err != nil {
		return nil, err
	}

	// Create singletons
	logger.Debug("Creating singletons")
	inqServ := NewInquisitorService()
	excServ := NewExecutorService()
	runServ := NewRunnerService()

	return NewManager(opts, configs, inqServ, excServ, runServ)
}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func (m *manager) getFoundations() (map[string]*cfclient.Config, error) {
	if len(m.options.Foundations) < 1 {
		return m.configs, nil
	}

	configs := make(map[string]*cfclient.Config, len(m.configs))
	for _, foundationName := range m.options.Foundations {
		if _, ok := m.configs[foundationName]; !ok {
			return nil, fmt.Errorf("Could not find config with name '%s'", foundationName)
		}

		configs[foundationName] = m.configs[foundationName]
	}

	return configs, nil
}

func (m *manager) GetOptions() ManagerOptions {
	opts := m.options
	return *opts
}

func (m *manager) GetConfigs() config.Configs {
	return m.configs
}

func (m *manager) QueryServices() (map[string]QueryService, error) {
	confs, err := m.getFoundations()
	if err != nil {
		return nil, err
	}

	pool := make(map[string]QueryService, len(confs))
	for foundationName, conf := range confs {
		inq, err := m.inqServ.GetInquisitor(conf)
		if err != nil {
			return nil, err
		}

		pool[foundationName] = NewQueryService(inq)
	}

	return pool, nil
}

func (m *manager) ProblemServices() (map[string]ProblemService, error) {
	confs, err := m.getFoundations()
	if err != nil {
		return nil, err
	}

	pool := make(map[string]ProblemService, len(confs))
	for foundationName, conf := range confs {
		inq, err := m.inqServ.GetInquisitor(conf)
		if err != nil {
			return nil, err
		}

		pool[foundationName] = NewProblemService(inq)
	}

	return pool, nil
}

func (m *manager) WorkerService() (WorkerService, error) {
	return NewWorkerService(m.excServ, m.runServ, m.options.PluginPath), nil
}
