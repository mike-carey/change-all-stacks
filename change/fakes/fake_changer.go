// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/mike-carey/change-all-stacks/change"
)

type FakeChanger struct {
	ChangeStackStub        func(string, string) (string, error)
	changeStackMutex       sync.RWMutex
	changeStackArgsForCall []struct {
		arg1 string
		arg2 string
	}
	changeStackReturns struct {
		result1 string
		result2 error
	}
	changeStackReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeChanger) ChangeStack(arg1 string, arg2 string) (string, error) {
	fake.changeStackMutex.Lock()
	ret, specificReturn := fake.changeStackReturnsOnCall[len(fake.changeStackArgsForCall)]
	fake.changeStackArgsForCall = append(fake.changeStackArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ChangeStack", []interface{}{arg1, arg2})
	fake.changeStackMutex.Unlock()
	if fake.ChangeStackStub != nil {
		return fake.ChangeStackStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.changeStackReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeChanger) ChangeStackCallCount() int {
	fake.changeStackMutex.RLock()
	defer fake.changeStackMutex.RUnlock()
	return len(fake.changeStackArgsForCall)
}

func (fake *FakeChanger) ChangeStackCalls(stub func(string, string) (string, error)) {
	fake.changeStackMutex.Lock()
	defer fake.changeStackMutex.Unlock()
	fake.ChangeStackStub = stub
}

func (fake *FakeChanger) ChangeStackArgsForCall(i int) (string, string) {
	fake.changeStackMutex.RLock()
	defer fake.changeStackMutex.RUnlock()
	argsForCall := fake.changeStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeChanger) ChangeStackReturns(result1 string, result2 error) {
	fake.changeStackMutex.Lock()
	defer fake.changeStackMutex.Unlock()
	fake.ChangeStackStub = nil
	fake.changeStackReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeChanger) ChangeStackReturnsOnCall(i int, result1 string, result2 error) {
	fake.changeStackMutex.Lock()
	defer fake.changeStackMutex.Unlock()
	fake.ChangeStackStub = nil
	if fake.changeStackReturnsOnCall == nil {
		fake.changeStackReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.changeStackReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeChanger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.changeStackMutex.RLock()
	defer fake.changeStackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeChanger) recordInvocation(key string, args []interface{}) {
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

var _ change.Changer = new(FakeChanger)
