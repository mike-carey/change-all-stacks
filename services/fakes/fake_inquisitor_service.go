// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/change-all-stacks/query"
	"github.com/mike-carey/change-all-stacks/services"
)

type FakeInquisitorService struct {
	GetInquisitorStub        func(*cfclient.Config) (query.Inquisitor, error)
	getInquisitorMutex       sync.RWMutex
	getInquisitorArgsForCall []struct {
		arg1 *cfclient.Config
	}
	getInquisitorReturns struct {
		result1 query.Inquisitor
		result2 error
	}
	getInquisitorReturnsOnCall map[int]struct {
		result1 query.Inquisitor
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInquisitorService) GetInquisitor(arg1 *cfclient.Config) (query.Inquisitor, error) {
	fake.getInquisitorMutex.Lock()
	ret, specificReturn := fake.getInquisitorReturnsOnCall[len(fake.getInquisitorArgsForCall)]
	fake.getInquisitorArgsForCall = append(fake.getInquisitorArgsForCall, struct {
		arg1 *cfclient.Config
	}{arg1})
	fake.recordInvocation("GetInquisitor", []interface{}{arg1})
	fake.getInquisitorMutex.Unlock()
	if fake.GetInquisitorStub != nil {
		return fake.GetInquisitorStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInquisitorReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInquisitorService) GetInquisitorCallCount() int {
	fake.getInquisitorMutex.RLock()
	defer fake.getInquisitorMutex.RUnlock()
	return len(fake.getInquisitorArgsForCall)
}

func (fake *FakeInquisitorService) GetInquisitorCalls(stub func(*cfclient.Config) (query.Inquisitor, error)) {
	fake.getInquisitorMutex.Lock()
	defer fake.getInquisitorMutex.Unlock()
	fake.GetInquisitorStub = stub
}

func (fake *FakeInquisitorService) GetInquisitorArgsForCall(i int) *cfclient.Config {
	fake.getInquisitorMutex.RLock()
	defer fake.getInquisitorMutex.RUnlock()
	argsForCall := fake.getInquisitorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInquisitorService) GetInquisitorReturns(result1 query.Inquisitor, result2 error) {
	fake.getInquisitorMutex.Lock()
	defer fake.getInquisitorMutex.Unlock()
	fake.GetInquisitorStub = nil
	fake.getInquisitorReturns = struct {
		result1 query.Inquisitor
		result2 error
	}{result1, result2}
}

func (fake *FakeInquisitorService) GetInquisitorReturnsOnCall(i int, result1 query.Inquisitor, result2 error) {
	fake.getInquisitorMutex.Lock()
	defer fake.getInquisitorMutex.Unlock()
	fake.GetInquisitorStub = nil
	if fake.getInquisitorReturnsOnCall == nil {
		fake.getInquisitorReturnsOnCall = make(map[int]struct {
			result1 query.Inquisitor
			result2 error
		})
	}
	fake.getInquisitorReturnsOnCall[i] = struct {
		result1 query.Inquisitor
		result2 error
	}{result1, result2}
}

func (fake *FakeInquisitorService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getInquisitorMutex.RLock()
	defer fake.getInquisitorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInquisitorService) recordInvocation(key string, args []interface{}) {
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

var _ services.InquisitorService = new(FakeInquisitorService)
