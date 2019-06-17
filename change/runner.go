package change

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"

	"github.com/mike-carey/cfquery/config"
	"github.com/mike-carey/cfquery/query"
	"github.com/mike-carey/cfquery/util"

	"code.cloudfoundry.org/cli/plugin"
)

type Runner struct {
	Config string
	Logger *Logger
	DryRun bool
	FromStack string
	ToStack string
	CliConnection plugin.CliConnection
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
		CliConnection: GetCliConnection(),
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
	r.Logger.Debugf("Checking stacks in: %s", foundationName)

	cli, err := cfclient.NewClient(conf)
	if err != nil {
		return err
	}

	r.Logger.Debugf("Creating inquisitor for: %s", foundationName)
	i := query.NewInquisitor(cli)

	// Grab every app
	r.Logger.Debugf("Loading all apps in %s", foundationName)
	apps, err := i.GetAllApps()
	if err != nil {
		return err
	}

	r.Logger.Debugf("Filtering %i apps by stackname '%s' in %s", len(apps), r.FromStack, foundationName)
	apps, err = apps.FilterByStackName(i, r.FromStack)
	if err != nil {
		return err
	}

	// Group by space
	r.Logger.Debugf("Grouping %i apps by space in %s", len(apps), foundationName)
	mapps, err := apps.GroupBySpace(i)
	if err != nil {
		return err
	}

	r.Logger.Debugf("Changing stacks from '%s' to '%s' for %i apps in %s", r.FromStack, r.ToStack, len(mapps), foundationName)
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
	r.Logger.Debugf("Getting space from guid")
	space, err := i.GetSpaceByGuid(spaceGuid)
	if err != nil {
		return err
	}

	r.Logger.Debugf("Getting org from guid")
	org, err := i.GetOrgByGuid(space.OrganizationGuid)
	if err != nil {
		return err
	}

	r.Logger.Debugf("Building changer")
	ch, err := NewChanger(r.CliConnection, space)
	if err != nil {
		return err
	}

	r.Logger.Debugf("Building Handler")
	h := NewHandlerWithStdout(ch)

	errCh := make(chan error)
	for _, app := range apps {
		go func(org *cfclient.Org, space *cfclient.Space, app *cfclient.App, stack string) {
			var err error
			if r.DryRun {
				err = h.HandleDryRun(org.Name, space.Name, app.Name, r.ToStack)
			} else {
				err = h.Handle(org.Name, space.Name, app.Name, r.ToStack)
			}
			errCh <- err
		}(org, space, &app, r.ToStack)
	}

	errPool := make([]error, 0)
	for _, _ = range apps {
		if err := <-errCh; err != nil {
			errPool = append(errPool, err)
		}
	}

	if len(errPool) > 0 {
		return util.StackErrors(errPool)
	}

	return nil
}
