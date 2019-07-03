// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bytes"
	"sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/change-all-stacks/cf"
)

type FakeWorker struct {
	WorkStub        func(string, []cfclient.App, string) (*bytes.Buffer, error)
	workMutex       sync.RWMutex
	workArgsForCall []struct {
		arg1 string
		arg2 []cfclient.App
		arg3 string
	}
	workReturns struct {
		result1 *bytes.Buffer
		result2 error
	}
	workReturnsOnCall map[int]struct {
		result1 *bytes.Buffer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorker) Work(arg1 string, arg2 []cfclient.App, arg3 string) (*bytes.Buffer, error) {
	var arg2Copy []cfclient.App
	if arg2 != nil {
		arg2Copy = make([]cfclient.App, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.workMutex.Lock()
	ret, specificReturn := fake.workReturnsOnCall[len(fake.workArgsForCall)]
	fake.workArgsForCall = append(fake.workArgsForCall, struct {
		arg1 string
		arg2 []cfclient.App
		arg3 string
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("Work", []interface{}{arg1, arg2Copy, arg3})
	fake.workMutex.Unlock()
	if fake.WorkStub != nil {
		return fake.WorkStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.workReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorker) WorkCallCount() int {
	fake.workMutex.RLock()
	defer fake.workMutex.RUnlock()
	return len(fake.workArgsForCall)
}

func (fake *FakeWorker) WorkCalls(stub func(string, []cfclient.App, string) (*bytes.Buffer, error)) {
	fake.workMutex.Lock()
	defer fake.workMutex.Unlock()
	fake.WorkStub = stub
}

func (fake *FakeWorker) WorkArgsForCall(i int) (string, []cfclient.App, string) {
	fake.workMutex.RLock()
	defer fake.workMutex.RUnlock()
	argsForCall := fake.workArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeWorker) WorkReturns(result1 *bytes.Buffer, result2 error) {
	fake.workMutex.Lock()
	defer fake.workMutex.Unlock()
	fake.WorkStub = nil
	fake.workReturns = struct {
		result1 *bytes.Buffer
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) WorkReturnsOnCall(i int, result1 *bytes.Buffer, result2 error) {
	fake.workMutex.Lock()
	defer fake.workMutex.Unlock()
	fake.WorkStub = nil
	if fake.workReturnsOnCall == nil {
		fake.workReturnsOnCall = make(map[int]struct {
			result1 *bytes.Buffer
			result2 error
		})
	}
	fake.workReturnsOnCall[i] = struct {
		result1 *bytes.Buffer
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.workMutex.RLock()
	defer fake.workMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorker) recordInvocation(key string, args []interface{}) {
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

var _ cf.Worker = new(FakeWorker)
