package cf

import (
	//
)

// Wraps the creation of runners so generated runners can be faked

//go:generate counterfeiter -o fakes/fake_runner_factory.go RunnerFactory
type RunnerFactory interface {
	CreateRunner(command CFCommand, opts *RunnerOptions) Runner
	CreateRunnerWithDefaultCommand(opts *RunnerOptions) Runner
}

type runnerFactory struct {}

func NewRunnerFactory() RunnerFactory {
	return &runnerFactory{}
}

func (r *runnerFactory) CreateRunnerWithDefaultCommand(opts *RunnerOptions) Runner {
	return NewRunnerWithDefaultCommand(opts)
}

func (r *runnerFactory) CreateRunner(command CFCommand, opts *RunnerOptions) Runner {
	return NewRunner(command, opts)
}
