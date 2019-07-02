package services

import (
	"github.com/mike-carey/change-all-stacks/cf"
)

//go:generate counterfeiter -o fakes/fake_runner_service.go RunnerService
type RunnerService interface {
	GetRunner(executor cf.Executor) cf.Runner
}

type runnerService struct {}

func NewRunnerService() RunnerService {
	return &runnerService{}
}

func (s *runnerService) GetRunner(executor cf.Executor) cf.Runner {
	return cf.NewRunner(executor)
}
