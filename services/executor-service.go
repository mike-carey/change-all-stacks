package services

import (
	"github.com/mike-carey/change-all-stacks/cf"
)

//go:generate counterfeiter -o fakes/fake_executor_service.go ExecutorService
type ExecutorService interface {
	CreateExecutor(command cf.CFCommand, dryrun bool) cf.Executor
	CreateExecutorWithDefaultCommand(dryrun bool) cf.Executor
}

type executorService struct {}

func NewExecutorService() ExecutorService {
	return &executorService{}
}

func (r *executorService) CreateExecutorWithDefaultCommand(dryrun bool) cf.Executor {
	return cf.NewExecutorWithDefaultCommand(dryrun)
}

func (r *executorService) CreateExecutor(command cf.CFCommand, dryrun bool) cf.Executor {
	return cf.NewExecutor(command, dryrun)
}
