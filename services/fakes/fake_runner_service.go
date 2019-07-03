// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/mike-carey/change-all-stacks/cf"
	"github.com/mike-carey/change-all-stacks/services"
)

type FakeRunnerService struct {
	GetRunnerStub        func(cf.Executor) cf.Runner
	getRunnerMutex       sync.RWMutex
	getRunnerArgsForCall []struct {
		arg1 cf.Executor
	}
	getRunnerReturns struct {
		result1 cf.Runner
	}
	getRunnerReturnsOnCall map[int]struct {
		result1 cf.Runner
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRunnerService) GetRunner(arg1 cf.Executor) cf.Runner {
	fake.getRunnerMutex.Lock()
	ret, specificReturn := fake.getRunnerReturnsOnCall[len(fake.getRunnerArgsForCall)]
	fake.getRunnerArgsForCall = append(fake.getRunnerArgsForCall, struct {
		arg1 cf.Executor
	}{arg1})
	fake.recordInvocation("GetRunner", []interface{}{arg1})
	fake.getRunnerMutex.Unlock()
	if fake.GetRunnerStub != nil {
		return fake.GetRunnerStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getRunnerReturns
	return fakeReturns.result1
}

func (fake *FakeRunnerService) GetRunnerCallCount() int {
	fake.getRunnerMutex.RLock()
	defer fake.getRunnerMutex.RUnlock()
	return len(fake.getRunnerArgsForCall)
}

func (fake *FakeRunnerService) GetRunnerCalls(stub func(cf.Executor) cf.Runner) {
	fake.getRunnerMutex.Lock()
	defer fake.getRunnerMutex.Unlock()
	fake.GetRunnerStub = stub
}

func (fake *FakeRunnerService) GetRunnerArgsForCall(i int) cf.Executor {
	fake.getRunnerMutex.RLock()
	defer fake.getRunnerMutex.RUnlock()
	argsForCall := fake.getRunnerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRunnerService) GetRunnerReturns(result1 cf.Runner) {
	fake.getRunnerMutex.Lock()
	defer fake.getRunnerMutex.Unlock()
	fake.GetRunnerStub = nil
	fake.getRunnerReturns = struct {
		result1 cf.Runner
	}{result1}
}

func (fake *FakeRunnerService) GetRunnerReturnsOnCall(i int, result1 cf.Runner) {
	fake.getRunnerMutex.Lock()
	defer fake.getRunnerMutex.Unlock()
	fake.GetRunnerStub = nil
	if fake.getRunnerReturnsOnCall == nil {
		fake.getRunnerReturnsOnCall = make(map[int]struct {
			result1 cf.Runner
		})
	}
	fake.getRunnerReturnsOnCall[i] = struct {
		result1 cf.Runner
	}{result1}
}

func (fake *FakeRunnerService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getRunnerMutex.RLock()
	defer fake.getRunnerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRunnerService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ services.RunnerService = new(FakeRunnerService)