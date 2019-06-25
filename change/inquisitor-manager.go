package change

import (
	"io/ioutil"
	"sync"

	"github.com/mike-carey/change-all-stacks/query"

	"github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_inquisitor_manager.go InquisitorManager
type InquisitorManager interface {
	GetInquisitor(config *cfclient.Config) (query.Inquisitor, error)
	GetHelper(config *cfclient.Config) (InquisitorHelper, error)
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

		inquisitor, err := query.NewInquisitor(cli, ioutil.Discard), nil
		if err != nil {
			return inquisitor, err
		}

		i.instances[config] = inquisitor
	}

	instance := i.instances[config]

	return instance, nil
}

func (i *inquisitorManager) GetHelper(config *cfclient.Config) (InquisitorHelper, error) {
	inq, err := i.GetInquisitor(config)
	if err != nil {
		return nil, err
	}

	return NewInquisitorHelper(inq), nil
}
