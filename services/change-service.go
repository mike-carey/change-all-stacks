package services

import (
	"github.com/mike-carey/change-all-stacks/cf"

	// cfclient "github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_change_service.go ChangeService
type ChangeService interface {
	GetWorkers(queryServices map[string]QueryService, orgNames ...string) (map[string][]cf.Runner, error)
}

func NewChangeService(executorService ExecutorService, workerService WorkerService) ChangeService {
	return &changeService{
		executorService: executorService,
		workerService: workerService,
	}
}

type changeService struct {
	executorService ExecutorService
	workerService WorkerService
}

func (s *changeService) GetWorkers(queryServices map[string]QueryService, orgNames ...string) (map[string][]cf.Runner, error) {
	runners := make(map[string][]cf.Runner, 0)
	for foundationName, _ := range queryServices {
		runners[foundationName] = make([]cf.Runner, 0)
	}
	return runners, nil
}
