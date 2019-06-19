package change

// Wraps the creation of runners so generated runners can be faked

//go:generate counterfeiter -o fakes/fake_runner_factory.go RunnerFactory
type RunnerFactory interface {
	CreateRunner(command CFCommand, logger Logger, dryrun bool) Runner
	CreateRunnerWithDefaultCommand(logger Logger, dryrun bool) Runner
}

type runnerFactory struct {}

func NewRunnerFactory() RunnerFactory {
	return &runnerFactory{}
}

func (r *runnerFactory) CreateRunnerWithDefaultCommand(logger Logger, dryrun bool) Runner {
	return NewRunnerWithDefaultCommand(logger, dryrun)
}

func (r *runnerFactory) CreateRunner(command CFCommand, logger Logger, dryrun bool) Runner {
	return NewRunner(command, logger, dryrun)
}
