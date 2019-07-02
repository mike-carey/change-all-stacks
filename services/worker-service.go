package services

import (
	"sync"
	"github.com/mike-carey/change-all-stacks/cf"
	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_worker_service.go WorkerService
type WorkerService interface {
	GetWorker(config *cfclient.Config, orgName string, spaceName string, dryRun bool) (cf.Worker, error)
}

func NewWorkerService(executorService ExecutorService, runnerService RunnerService, pluginPath string) WorkerService {
	return &workerService{
		executorService: executorService,
		runnerService: runnerService,
		pluginPath: pluginPath,
		instances: make(cache, 0),
		mutex: sync.Mutex{},
	}
}

type workerService struct {
	executorService ExecutorService
	runnerService RunnerService
	pluginPath string
	instances cache
	mutex sync.Mutex
}

func (i *workerService) lock() {
	logger.Debugf("Locking worker service")
	i.mutex.Lock()
	logger.Debugf("Locked worker service")
}

func (i *workerService) unlock() {
	logger.Debugf("Unlocking worker service")
	i.mutex.Unlock()
	logger.Debugf("Unlocked worker service")
}

type cacheObject struct {
	config *cfclient.Config
	orgName string
	spaceName string
}

type spaceCache map[string]cf.Worker
type orgCache map[string]spaceCache
type cache map[*cfclient.Config]orgCache

func (c spaceCache) Has(key string) bool {
	_, ok := c[key]
	return ok
}
func (c orgCache) Has(key string) bool {
	_, ok := c[key]
	return ok
}
func (c orgCache) Ensure(key string) {
	if !c.Has(key) {
		c[key] = make(spaceCache, 0)
	}
}
func (c cache) Has(key *cfclient.Config) bool {
	_, ok := c[key]
	return ok
}
func (c cache) Ensure(key *cfclient.Config) {
	if !c.Has(key) {
		c[key] = make(orgCache, 0)
	}
}
func (c cache) EnsureAt(key1 *cfclient.Config, key2 string) spaceCache {
	c.Ensure(key1)
	c[key1].Ensure(key2)
	return c[key1][key2]
}
func (c cache) SetAt(key1 *cfclient.Config, key2 string, key3 string, value cf.Worker) {
	c[key1][key2][key3] = value
}

func (i *workerService) GetWorker(config *cfclient.Config, orgName string, spaceName string, dryRun bool) (cf.Worker, error) {
	i.lock()
	defer i.unlock()

	logger.Debugf("workerService=%v", i)

	if !i.instances.EnsureAt(config, orgName).Has(spaceName) {
		// Create a runner
		executor := i.executorService.CreateExecutorWithDefaultCommand(dryRun)
		logger.Debugf("%v", executor)
		runner := i.runnerService.GetRunner(executor)
		worker := cf.NewWorker(runner, config, i.pluginPath, orgName, spaceName)

		i.instances.SetAt(config, orgName, spaceName, worker)
	}

	instance := i.instances[config][orgName][spaceName]

	return instance, nil
}
