package cf

import (
	"fmt"
	"strings"

	"github.com/mike-carey/change-all-stacks/logger"
)

const (
	ApiCmd = "api"
	AuthCmd = "auth"
	TargetCmd = "target"
	ChangeStackCmd = "change-stack"
	InstallPluginCmd = "install-plugin"

	RedactedString = "<REDACTED>"

	DefaultDryRun = false
)

type RunnerOptions struct {
	DryRun bool
}

type Redact string

//go:generate counterfeiter -o fakes/fake_runner.go Runner
type Runner interface {
	Api(apiAddress string, skipSSlValidation bool) error
	Auth(username string, password string) error
	Target(org string, space string) error
	InstallPlugin(plugin string) error
	ChangeStack(app string, stack string) error
}

func NewRunnerWithDefaultCommand(opts *RunnerOptions) Runner {
	home := NewTempCFHome("")

	return NewRunner(NewDefaultCFCommand(home), opts)
}

func NewRunner(command CFCommand, opts *RunnerOptions) Runner {
	return &runner{
		command: command,
		dryrun: opts.DryRun,
	}
}

type runner struct {
	Name string
	home CFHome
	command CFCommand
	dryrun bool
}

func (r *runner) run(args ...interface{}) error {
	printArgs := make([]string, len(args))
	runArgs := make([]string, len(args))

	for i, arg := range args {
		if str, ok := arg.(string); ok {
			runArgs[i] = str
			printArgs[i] = str
		} else if red, ok := arg.(Redact); ok {
			runArgs[i] = string(red)
			printArgs[i] = RedactedString
		} else {
			return fmt.Errorf("Unknown type provided to run command")
		}
	}

	var err error
	if r.dryrun {
		str, err := r.command.String(printArgs...)
		if err == nil {
			fmt.Println(str)
		}
	} else {
		logger.Debugf("Running cf command: %s", strings.Join(printArgs, " "))
		err = r.command.Execute(runArgs...)
	}

	return err
}

func (r *runner) Api(apiAddress string, skipSslValidation bool) error {
	ssv := ""
	if skipSslValidation {
		ssv = "--skip-ssl-validation"
	}

	return r.run(ApiCmd, apiAddress, ssv)
}

func (r *runner) Auth(username string, password string) error {
	return r.run(AuthCmd, username, Redact(password))
}

func (r *runner) Target(org string, space string) error {
	return r.run(TargetCmd, "-o", org, "-s", space)
}

func (r *runner) InstallPlugin(plugin string) error {
	return r.run(InstallPluginCmd, plugin, "-f")
}

func (r *runner) ChangeStack(app string, stack string) error {
	return r.run(ChangeStackCmd, app, stack)
}
