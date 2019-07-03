// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/mike-carey/change-all-stacks/cf"
	"github.com/mike-carey/change-all-stacks/services"
)

type FakeExecutorService struct {
	CreateExecutorStub        func(cf.CFCommand, bool) cf.Executor
	createExecutorMutex       sync.RWMutex
	createExecutorArgsForCall []struct {
		arg1 cf.CFCommand
		arg2 bool
	}
	createExecutorReturns struct {
		result1 cf.Executor
	}
	createExecutorReturnsOnCall map[int]struct {
		result1 cf.Executor
	}
	CreateExecutorWithDefaultCommandStub        func(bool) cf.Executor
	createExecutorWithDefaultCommandMutex       sync.RWMutex
	createExecutorWithDefaultCommandArgsForCall []struct {
		arg1 bool
	}
	createExecutorWithDefaultCommandReturns struct {
		result1 cf.Executor
	}
	createExecutorWithDefaultCommandReturnsOnCall map[int]struct {
		result1 cf.Executor
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExecutorService) CreateExecutor(arg1 cf.CFCommand, arg2 bool) cf.Executor {
	fake.createExecutorMutex.Lock()
	ret, specificReturn := fake.createExecutorReturnsOnCall[len(fake.createExecutorArgsForCall)]
	fake.createExecutorArgsForCall = append(fake.createExecutorArgsForCall, struct {
		arg1 cf.CFCommand
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("CreateExecutor", []interface{}{arg1, arg2})
	fake.createExecutorMutex.Unlock()
	if fake.CreateExecutorStub != nil {
		return fake.CreateExecutorStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createExecutorReturns
	return fakeReturns.result1
}

func (fake *FakeExecutorService) CreateExecutorCallCount() int {
	fake.createExecutorMutex.RLock()
	defer fake.createExecutorMutex.RUnlock()
	return len(fake.createExecutorArgsForCall)
}

func (fake *FakeExecutorService) CreateExecutorCalls(stub func(cf.CFCommand, bool) cf.Executor) {
	fake.createExecutorMutex.Lock()
	defer fake.createExecutorMutex.Unlock()
	fake.CreateExecutorStub = stub
}

func (fake *FakeExecutorService) CreateExecutorArgsForCall(i int) (cf.CFCommand, bool) {
	fake.createExecutorMutex.RLock()
	defer fake.createExecutorMutex.RUnlock()
	argsForCall := fake.createExecutorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExecutorService) CreateExecutorReturns(result1 cf.Executor) {
	fake.createExecutorMutex.Lock()
	defer fake.createExecutorMutex.Unlock()
	fake.CreateExecutorStub = nil
	fake.createExecutorReturns = struct {
		result1 cf.Executor
	}{result1}
}

func (fake *FakeExecutorService) CreateExecutorReturnsOnCall(i int, result1 cf.Executor) {
	fake.createExecutorMutex.Lock()
	defer fake.createExecutorMutex.Unlock()
	fake.CreateExecutorStub = nil
	if fake.createExecutorReturnsOnCall == nil {
		fake.createExecutorReturnsOnCall = make(map[int]struct {
			result1 cf.Executor
		})
	}
	fake.createExecutorReturnsOnCall[i] = struct {
		result1 cf.Executor
	}{result1}
}

func (fake *FakeExecutorService) CreateExecutorWithDefaultCommand(arg1 bool) cf.Executor {
	fake.createExecutorWithDefaultCommandMutex.Lock()
	ret, specificReturn := fake.createExecutorWithDefaultCommandReturnsOnCall[len(fake.createExecutorWithDefaultCommandArgsForCall)]
	fake.createExecutorWithDefaultCommandArgsForCall = append(fake.createExecutorWithDefaultCommandArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("CreateExecutorWithDefaultCommand", []interface{}{arg1})
	fake.createExecutorWithDefaultCommandMutex.Unlock()
	if fake.CreateExecutorWithDefaultCommandStub != nil {
		return fake.CreateExecutorWithDefaultCommandStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createExecutorWithDefaultCommandReturns
	return fakeReturns.result1
}

func (fake *FakeExecutorService) CreateExecutorWithDefaultCommandCallCount() int {
	fake.createExecutorWithDefaultCommandMutex.RLock()
	defer fake.createExecutorWithDefaultCommandMutex.RUnlock()
	return len(fake.createExecutorWithDefaultCommandArgsForCall)
}

func (fake *FakeExecutorService) CreateExecutorWithDefaultCommandCalls(stub func(bool) cf.Executor) {
	fake.createExecutorWithDefaultCommandMutex.Lock()
	defer fake.createExecutorWithDefaultCommandMutex.Unlock()
	fake.CreateExecutorWithDefaultCommandStub = stub
}

func (fake *FakeExecutorService) CreateExecutorWithDefaultCommandArgsForCall(i int) bool {
	fake.createExecutorWithDefaultCommandMutex.RLock()
	defer fake.createExecutorWithDefaultCommandMutex.RUnlock()
	argsForCall := fake.createExecutorWithDefaultCommandArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeExecutorService) CreateExecutorWithDefaultCommandReturns(result1 cf.Executor) {
	fake.createExecutorWithDefaultCommandMutex.Lock()
	defer fake.createExecutorWithDefaultCommandMutex.Unlock()
	fake.CreateExecutorWithDefaultCommandStub = nil
	fake.createExecutorWithDefaultCommandReturns = struct {
		result1 cf.Executor
	}{result1}
}

func (fake *FakeExecutorService) CreateExecutorWithDefaultCommandReturnsOnCall(i int, result1 cf.Executor) {
	fake.createExecutorWithDefaultCommandMutex.Lock()
	defer fake.createExecutorWithDefaultCommandMutex.Unlock()
	fake.CreateExecutorWithDefaultCommandStub = nil
	if fake.createExecutorWithDefaultCommandReturnsOnCall == nil {
		fake.createExecutorWithDefaultCommandReturnsOnCall = make(map[int]struct {
			result1 cf.Executor
		})
	}
	fake.createExecutorWithDefaultCommandReturnsOnCall[i] = struct {
		result1 cf.Executor
	}{result1}
}

func (fake *FakeExecutorService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createExecutorMutex.RLock()
	defer fake.createExecutorMutex.RUnlock()
	fake.createExecutorWithDefaultCommandMutex.RLock()
	defer fake.createExecutorWithDefaultCommandMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExecutorService) recordInvocation(key string, args []interface{}) {
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

var _ services.ExecutorService = new(FakeExecutorService)
