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
		mutex: sync.Mutex{},
	}
}

type inquisitorManager struct {
	instances map[*cfclient.Config]query.Inquisitor
	mutex sync.Mutex
}

func (i *inquisitorManager) GetInquisitor(config *cfclient.Config) (query.Inquisitor, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if _, ok := i.instances[config]; !ok {
		cli, err := cfclient.NewClient(config)
		if err != nil {
			return nil, err
		}

		inquisitor, err := query.NewInquisitor(cli), nil
		if err != nil {
			return inquisitor, err
		}

		i.instances[config] = inquisitor
	}

	instance := i.instances[config]

	return instance, nil
}
