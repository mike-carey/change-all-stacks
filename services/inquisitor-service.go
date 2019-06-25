package services

import (
	"sync"
	"io/ioutil"

	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/mike-carey/change-all-stacks/query"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_inquisitor_service.go InquisitorService
type InquisitorService interface{
	GetInquisitor(config *cfclient.Config) (query.Inquisitor, error)
}

func NewInquisitorService() InquisitorService {
	return &inquisitorService{
		instances: make(map[*cfclient.Config]query.Inquisitor, 0),
		mutex: sync.Mutex{},
	}
}

type inquisitorService struct {
	instances map[*cfclient.Config]query.Inquisitor
	mutex sync.Mutex
}

func (i *inquisitorService) lock() {
	logger.Debugf("Locking inquisitor service")
	i.mutex.Lock()
	logger.Debugf("Locked inquisitor service")
}

func (i *inquisitorService) unlock() {
	logger.Debugf("Unlocking inquisitor service")
	i.mutex.Unlock()
	logger.Debugf("Unlocked inquisitor service")
}

func (i *inquisitorService) GetInquisitor(config *cfclient.Config) (query.Inquisitor, error) {
	i.lock()
	defer i.unlock()

	if _, ok := i.instances[config]; !ok {
		cli, err := cfclient.NewClient(config)
		if err != nil {
			return nil, err
		}

		// TODO: Replace ioutil.Discard with logger
		inquisitor, err := query.NewInquisitor(cli, ioutil.Discard), nil
		if err != nil {
			return inquisitor, err
		}

		i.instances[config] = inquisitor
	}

	instance := i.instances[config]

	return instance, nil
}
