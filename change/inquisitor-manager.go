package change

import (
	"sync"

	"github.com/mike-carey/cfquery/query"

	"github.com/cloudfoundry-community/go-cfclient"
)

type InquisitorManager interface {
	GetInquisitor(config *cfclient.Config) (query.Inquisitor, error)
}

func NewInquisitorManager() InquisitorManager {
	return &inquisitorManager{
		instances: make(map[*cfclient.Config]query.Inquisitor, 0),
		mutex: sync.RWMutex{},
	}
}

type inquisitorManager struct {
	instances map[*cfclient.Config]query.Inquisitor
	mutex sync.RWMutex
}

func (i *inquisitorManager) GetInquisitor(config *cfclient.Config) (query.Inquisitor, error) {
	i.mutex.RLock()
	if _, ok := i.instances[config]; !ok {
		i.mutex.Lock()
		cli, err := cfclient.NewClient(config)
		if err != nil {
			return nil, err
		}

		inquisitor, err := query.NewInquisitor(cli), nil
		if err != nil {
			return inquisitor, err
		}

		i.instances[config] = inquisitor
		i.mutex.Unlock()
	}

	instance := i.instances[config]

	i.mutex.RUnlock()

	return instance, nil
}
