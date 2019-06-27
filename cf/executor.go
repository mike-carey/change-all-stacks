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

type Redact string

//go:generate counterfeiter -o fakes/fake_executor.go Executor
type Executor interface {
	Api(apiAddress string, skipSSlValidation bool) error
	Auth(username string, password string) error
	Target(org string, space string) error
	InstallPlugin(plugin string) error
	ChangeStack(app string, stack string) error
}

func NewExecutorWithDefaultCommand(dryrun bool) Executor {
	home := NewTempCFHome("")

	return NewExecutor(NewDefaultCFCommand(home), dryrun)
}

func NewExecutor(command CFCommand, dryrun bool) Executor {
	return &executor{
		command: command,
		dryrun: dryrun,
	}
}

type executor struct {
	Name string
	home CFHome
	command CFCommand
	dryrun bool
}

func (r *executor) run(args ...interface{}) error {
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

func (r *executor) Api(apiAddress string, skipSslValidation bool) error {
	ssv := ""
	if skipSslValidation {
		ssv = "--skip-ssl-validation"
	}

	return r.run(ApiCmd, apiAddress, ssv)
}

func (r *executor) Auth(username string, password string) error {
	return r.run(AuthCmd, username, Redact(password))
}

func (r *executor) Target(org string, space string) error {
	return r.run(TargetCmd, "-o", org, "-s", space)
}

func (r *executor) InstallPlugin(plugin string) error {
	return r.run(InstallPluginCmd, plugin, "-f")
}

func (r *executor) ChangeStack(app string, stack string) error {
	return r.run(ChangeStackCmd, app, stack)
}
