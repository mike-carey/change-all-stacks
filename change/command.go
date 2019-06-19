package change

import (
	"io"
	"os"
	"fmt"
	"os/exec"
	"strings"
	"io/ioutil"
)

//go:generate counterfeiter -o fakes/fake_cf_command.go CFCommand
type CFCommand interface {
	Execute(args ...string) error
	String(args ...string) (string, error)
}

//go:generate counterfeiter -o fakes/fake_cf_home.go CFHome
type CFHome interface {
	Directory() (string, error)
}

type cfHome struct {
	prefix string
	directory string
}

func (c *cfHome) Directory() (string, error) {
	if c.directory == "" {
		dir, err := ioutil.TempDir(os.TempDir(), c.prefix)
		if err != nil {
			return "", err
		}

		c.directory = dir
	}

	return c.directory, nil
}

func NewTempCFHome(prefix string) CFHome {
	return &cfHome{
		prefix: prefix,
	}
}

type cfCommand struct {
	home CFHome
	path string
	stdout io.Writer
	stderr io.Writer
}

func NewDefaultCFCommand(home CFHome) CFCommand {
	return NewCFCommand(home, "cf", os.Stdout, os.Stderr)
}

func NewCFCommand(home CFHome, path string, stdout io.Writer, stderr io.Writer) CFCommand {
	return &cfCommand{
		home: home,
		path: path,
		stdout: stdout,
		stderr: stderr,
	}
}

func (cf *cfCommand) command(args ...string) (*exec.Cmd, error) {
	home, err := cf.home.Directory()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(cf.path, args...)

	env := os.Environ()
	env = append(env, fmt.Sprintf("CF_HOME=%s", home))

	// Set the stage
	cmd.Stdout = cf.stdout
	cmd.Stderr = cf.stderr
	cmd.Env = env

	return cmd, nil
}

func (cf *cfCommand) String(args ...string) (string, error) {
	cmd, err := cf.command(args...)
	if err != nil {
		return "", err
	}

	allArgs := append([]string{cmd.Env[len(cmd.Env)-1], cmd.Path}, cmd.Args[1:]...)

	return strings.Join(allArgs, " "), nil
}

func (cf *cfCommand) Execute(args ...string) error {
	cmd, err := cf.command(args...)
	if err != nil {
		return err
	}

	return cmd.Run()
}

func (cf *cfCommand) ExecuteAsync(args ...string) error {
	cmd, err := cf.command(args...)
	if err != nil {
		return err
	}

	return cmd.Start()
}
