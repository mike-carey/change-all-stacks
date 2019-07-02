package cf

import (
	"fmt"
	"bytes"
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
	Buffer() *bytes.Buffer
}

func NewExecutorWithDefaultCommand(dryrun bool) Executor {
	home := NewTempCFHome("")
	cmd := NewDefaultCFCommand(home)

	return NewExecutor(cmd, dryrun)
}

func NewExecutor(command CFCommand, dryrun bool) Executor {
	return &executor{
		command: command,
		dryrun: dryrun,
		buffer: bytes.NewBuffer(nil),
	}
}

type executor struct {
	command CFCommand
	dryrun bool
	buffer *bytes.Buffer
}

func (r *executor) Buffer() *bytes.Buffer {
	return r.buffer
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
			r.buffer.WriteString(str + "\n")
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
