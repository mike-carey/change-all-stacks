// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/mike-carey/change-all-stacks/data"
)

type FakeFormatter struct {
	FormatDataStub        func(data.Data) (string, error)
	formatDataMutex       sync.RWMutex
	formatDataArgsForCall []struct {
		arg1 data.Data
	}
	formatDataReturns struct {
		result1 string
		result2 error
	}
	formatDataReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	FormatProblemSetStub        func(data.ProblemSet) (string, error)
	formatProblemSetMutex       sync.RWMutex
	formatProblemSetArgsForCall []struct {
		arg1 data.ProblemSet
	}
	formatProblemSetReturns struct {
		result1 string
		result2 error
	}
	formatProblemSetReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFormatter) FormatData(arg1 data.Data) (string, error) {
	fake.formatDataMutex.Lock()
	ret, specificReturn := fake.formatDataReturnsOnCall[len(fake.formatDataArgsForCall)]
	fake.formatDataArgsForCall = append(fake.formatDataArgsForCall, struct {
		arg1 data.Data
	}{arg1})
	fake.recordInvocation("FormatData", []interface{}{arg1})
	fake.formatDataMutex.Unlock()
	if fake.FormatDataStub != nil {
		return fake.FormatDataStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.formatDataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFormatter) FormatDataCallCount() int {
	fake.formatDataMutex.RLock()
	defer fake.formatDataMutex.RUnlock()
	return len(fake.formatDataArgsForCall)
}

func (fake *FakeFormatter) FormatDataCalls(stub func(data.Data) (string, error)) {
	fake.formatDataMutex.Lock()
	defer fake.formatDataMutex.Unlock()
	fake.FormatDataStub = stub
}

func (fake *FakeFormatter) FormatDataArgsForCall(i int) data.Data {
	fake.formatDataMutex.RLock()
	defer fake.formatDataMutex.RUnlock()
	argsForCall := fake.formatDataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFormatter) FormatDataReturns(result1 string, result2 error) {
	fake.formatDataMutex.Lock()
	defer fake.formatDataMutex.Unlock()
	fake.FormatDataStub = nil
	fake.formatDataReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFormatter) FormatDataReturnsOnCall(i int, result1 string, result2 error) {
	fake.formatDataMutex.Lock()
	defer fake.formatDataMutex.Unlock()
	fake.FormatDataStub = nil
	if fake.formatDataReturnsOnCall == nil {
		fake.formatDataReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.formatDataReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFormatter) FormatProblemSet(arg1 data.ProblemSet) (string, error) {
	fake.formatProblemSetMutex.Lock()
	ret, specificReturn := fake.formatProblemSetReturnsOnCall[len(fake.formatProblemSetArgsForCall)]
	fake.formatProblemSetArgsForCall = append(fake.formatProblemSetArgsForCall, struct {
		arg1 data.ProblemSet
	}{arg1})
	fake.recordInvocation("FormatProblemSet", []interface{}{arg1})
	fake.formatProblemSetMutex.Unlock()
	if fake.FormatProblemSetStub != nil {
		return fake.FormatProblemSetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.formatProblemSetReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFormatter) FormatProblemSetCallCount() int {
	fake.formatProblemSetMutex.RLock()
	defer fake.formatProblemSetMutex.RUnlock()
	return len(fake.formatProblemSetArgsForCall)
}

func (fake *FakeFormatter) FormatProblemSetCalls(stub func(data.ProblemSet) (string, error)) {
	fake.formatProblemSetMutex.Lock()
	defer fake.formatProblemSetMutex.Unlock()
	fake.FormatProblemSetStub = stub
}

func (fake *FakeFormatter) FormatProblemSetArgsForCall(i int) data.ProblemSet {
	fake.formatProblemSetMutex.RLock()
	defer fake.formatProblemSetMutex.RUnlock()
	argsForCall := fake.formatProblemSetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFormatter) FormatProblemSetReturns(result1 string, result2 error) {
	fake.formatProblemSetMutex.Lock()
	defer fake.formatProblemSetMutex.Unlock()
	fake.FormatProblemSetStub = nil
	fake.formatProblemSetReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFormatter) FormatProblemSetReturnsOnCall(i int, result1 string, result2 error) {
	fake.formatProblemSetMutex.Lock()
	defer fake.formatProblemSetMutex.Unlock()
	fake.FormatProblemSetStub = nil
	if fake.formatProblemSetReturnsOnCall == nil {
		fake.formatProblemSetReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.formatProblemSetReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeFormatter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.formatDataMutex.RLock()
	defer fake.formatDataMutex.RUnlock()
	fake.formatProblemSetMutex.RLock()
	defer fake.formatProblemSetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFormatter) recordInvocation(key string, args []interface{}) {
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

var _ data.Formatter = new(FakeFormatter)
