package change_all_stacks

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"

	"github.com/mike-carey/cfquery/config"
	"github.com/mike-carey/cfquery/query"
	"github.com/mike-carey/cfquery/util"
)

const (
	Stack = "cflinuxfs3"
)

type Options struct {
	Config  string `short:"c" long:"config" description:"The configuration file to load" default:"cf.json"`
	DryRun  bool   `short:"d" long:"dry-run" description:"Does not actually do the stack change, but instead prints what it would do"`
	Verbose bool   `short:"v" long:"verbose" description:"Prints more output"`
	Version bool   `long:"version" description:"Prints the version of the cli"`

	FromStack string
	ToStack string
}

func Go(opts *Options) error {
	return NewRunner(opts.Config, opts.Verbose, opts.DryRun, opts.FromStack, opts.ToStack).Run()
}

type Runner struct {
	Config string
	Logger *Logger
	DryRun bool
	FromStack string
	ToStack string
}

func NewRunner(config string, verbose bool, dryrun bool, fromStack string, toStack string) *Runner {
	return &Runner{
		Config: config,
		Logger: &Logger{
			Verbose: verbose,
		},
		DryRun: dryrun,
		FromStack: fromStack,
		ToStack: toStack,
	}
}

func (r *Runner) Run() error {
	r.Logger.Debugf("Loading config: %s", r.Config)

	foundations, err := config.LoadConfig(r.Config)
	if err != nil {
		return err
	}

	err = r.ChangeAllStacks(foundations)
	if err != nil {
		return err
	}

	r.Logger.Debug("All done!")

	return nil
}

func (r *Runner) ChangeAllStacks(foundations config.Foundations) error {
	errCh := make(chan error, 0)

	for name, conf := range foundations {
		go func(name string, conf *cfclient.Config) {
			errCh <- r.ChangeAllStacksInFoundation(name, conf)
		}(name, conf)
	}

	errPool := make([]error, 0)
	for _, _ = range foundations {
		if err := <-errCh; err != nil {
			errPool = append(errPool, err)
		}
	}

	if len(errPool) > 0 {
		return util.StackErrors(errPool)
	}

	return nil
}

func (r *Runner) ChangeAllStacksInFoundation(foundationName string, conf *cfclient.Config) error {
	r.Logger.Debugf("Changing stacks in: %s", foundationName)

	cli, err := cfclient.NewClient(conf)
	if err != nil {
		return err
	}

	i := query.NewInquisitor(cli)

	// Grab every app
	apps, err := i.GetAllApps()
	if err != nil {
		return err
	}

	apps, err = apps.FilterByStackName(i, r.FromStack)
	if err != nil {
		return err
	}

	// Group by space
	mapps, err := apps.GroupBySpace(i)
	if err != nil {
		return err
	}

	errCh := make(chan error)
	for spaceGuid, apps := range mapps {
		go func(i query.Inquisitor, spaceGuid string, apps query.Apps) {
			errCh <- r.ChangeStacksInSpace(i, spaceGuid, apps)
		}(i, spaceGuid, apps)
	}

	errPool := make([]error, 0)
	for _, _ = range mapps {
		if err := <-errCh; err != nil {
			errPool = append(errPool, err)
		}
	}

	if len(errPool) > 0 {
		return util.StackErrors(errPool)
	}

	return nil
}

func (r *Runner) ChangeStacksInSpace(i query.Inquisitor, spaceGuid string, apps query.Apps) error {
	// Every space needs a new cliconnection
	ch, err := GetChanger(i, spaceGuid)
	if err != nil {
		return err
	}

	errCh := make(chan error)

	for _, app := range apps {
		go func(ch Changer, app cfclient.App) {
			errCh <- r.ChangeStackInApp(ch, app)
		}(ch, app)
	}

	errPool := make([]error, 0)
	for _, _ = range apps {
		if err = <-errCh; err != nil {
			errPool = append(errPool, err)
		}
	}

	if len(errPool) > 0 {
		return util.StackErrors(errPool)
	}

	return nil
}

func (r *Runner) ChangeStackInApp(ch Changer, app cfclient.App) error {
	stack := r.ToStack

	r.Logger.Debugf("Changing %s's stack to '%s'", app.Name, stack)
	r.Logger.Infof("%s", app.Name)

	var err error
	if r.DryRun {
		r.Logger.Debug("Dry Run enabled")
	} else {
		str, err := ch.ChangeStack(app.Name, stack)
		if err == nil {
			r.Logger.Debug(str)
		}
	}

	return err
}
