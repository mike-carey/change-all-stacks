package commands

import (
	"os"
	"fmt"
	"sync"
	"bytes"
	"github.com/mike-carey/change-all-stacks/errors"
	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/mike-carey/change-all-stacks/services"
	"github.com/cloudfoundry-community/go-cfclient"
)

type ChangeCommand struct {
	DryRun bool `short:"d" long:"dry-run" description:"Simply prints the commands instead of running them"`
	Interactive bool `short:"i" long:"interactive" description:"Print the dry run before apply"`
	Threads int `short:"t" long:"threads" description:"The number of threads to run" default:"10"`

	Stacks struct {
		From string
		To string
	} `positional-args:"yes" required:"yes"`
}

func (c *ChangeCommand) Execute([]string) error {
	qss, err := manager.QueryServices()
	if err != nil {
		return err
	}

	workerService, err := manager.WorkerService()
	if err != nil {
		return err
	}

	mOpts := manager.GetOptions()
	confs := manager.GetConfigs()

	buffCh := make(chan bytes.Buffer, 0)
	errCh := make(chan error, 0)

	sem := make(chan int, c.Threads)


	for foundationName, qs := range qss {
		conf, err := confs.Get(foundationName)
		if err != nil {
			return err
		}

		apps, err := qs.GetAllAppsWithinOrgs(mOpts.Orgs...)
		if err != nil {
			return err
		}

		// Let's group our apps by org and space
		groupedApps, err := qs.GroupAppsByOrgAndSpace(apps)
		if err != nil {
			return err
		}

		go func(sem chan int, foundationName string, conf cfclient.Config, groupedApps map[string]map[string][]cfclient.App) {
			buff, err := c.run(sem, foundationName, conf, workerService, groupedApps, c.Stacks.To, c.DryRun)
			if err != nil {
				logger.Debugf("%s had an error: %v!", foundationName, err)
				errCh <- err
				return
			}

			logger.Debugf("%s did not have an error!", foundationName)
			buffCh <- buff
		}(sem, foundationName, *conf, groupedApps)
	}

	logger.Debugf("...")
	errPool := make([]error, 0)
	buffer := bytes.NewBuffer(nil)
	for _, _ = range qss {
		select {
		case e := <-errCh:
			if e != nil {
				errPool = append(errPool, e)
			}
		case b := <-buffCh:
			_b := &b
			_b.WriteTo(buffer)
		}
	}

	if len(errPool) > 0 {
		return errors.NewErrorStack("Error!", errPool)
	}

	logger.Debugf("Writing buffer to stdout")
	buffer.WriteTo(os.Stdout)

	return nil
}

func (c *ChangeCommand) run(sem chan int, foundationName string, conf cfclient.Config, workerService services.WorkerService, groupedApps map[string]map[string][]cfclient.App, stack string, dryrun bool) (bytes.Buffer, error) {
	buff := NewAsyncBuffer([]byte(fmt.Sprintf("Foundation: %s\n", foundationName)))
	errs := NewAsyncErrorPool()

	count := 0
	wg := sync.WaitGroup{}
	for orgName, gApps := range groupedApps {
		for spaceName, apps := range gApps {
			threadName := fmt.Sprintf("Thread-%s-%d", foundationName, count)

			count += 1
			sem <- 1
			wg.Add(1)

			logger.Debugf("%s is starting", threadName)
			go func(name string, conf cfclient.Config, orgName string, spaceName string, apps []cfclient.App, dryrun bool) {
				defer func() {
					logger.Debugf("%s is done", name)
					<-sem
					wg.Done()
				}()

				worker, err := workerService.GetWorker(&conf, orgName, spaceName, dryrun)
				if err != nil {
					logger.Debugf("%s had an error getting the worker", name)
					errs.Add(err)
					return
				}

				logger.Debugf("%s: Working to change %d apps to %s", threadName, len(apps), stack)
				b, err := worker.Work(threadName, apps, stack)
				if err != nil {
					logger.Debugf("%s had an error doing the work", name)
					errs.Add(err)
					return
				}

				buff.CopyFrom(b)
			}(threadName, conf, orgName, spaceName, apps, dryrun)
		}
	}

	wg.Wait()

	logger.Debugf("%s is done", foundationName)

	if errs.Len() > 0 {
		return bytes.Buffer{}, errors.NewErrorStack("Failed to run change commands", errs.Pool())
	}

	return buff.Buffer(), nil
}
