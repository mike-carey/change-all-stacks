package cf

import (
	"github.com/mike-carey/change-all-stacks/logger"
)

// Wraps the creation of runners so generated runners can be faked

//go:generate counterfeiter -o fakes/fake_runner_factory.go RunnerFactory
type RunnerFactory interface {
	CreateRunner(command CFCommand, logger logger.Logger, dryrun bool) Runner
	CreateRunnerWithDefaultCommand(logger logger.Logger, dryrun bool) Runner
}

type runnerFactory struct {}

func NewRunnerFactory() RunnerFactory {
	return &runnerFactory{}
}

func (r *runnerFactory) CreateRunnerWithDefaultCommand(logger logger.Logger, dryrun bool) Runner {
	return NewRunnerWithDefaultCommand(logger, dryrun)
}

func (r *runnerFactory) CreateRunner(command CFCommand, logger logger.Logger, dryrun bool) Runner {
	return NewRunner(command, logger, dryrun)
}
