// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bytes"
	"sync"

	"github.com/mike-carey/change-all-stacks/cf"
)

type FakeExecutor struct {
	ApiStub        func(string, bool) error
	apiMutex       sync.RWMutex
	apiArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	apiReturns struct {
		result1 error
	}
	apiReturnsOnCall map[int]struct {
		result1 error
	}
	AuthStub        func(string, string) error
	authMutex       sync.RWMutex
	authArgsForCall []struct {
		arg1 string
		arg2 string
	}
	authReturns struct {
		result1 error
	}
	authReturnsOnCall map[int]struct {
		result1 error
	}
	BufferStub        func() *bytes.Buffer
	bufferMutex       sync.RWMutex
	bufferArgsForCall []struct {
	}
	bufferReturns struct {
		result1 *bytes.Buffer
	}
	bufferReturnsOnCall map[int]struct {
		result1 *bytes.Buffer
	}
	ChangeStackStub        func(string, string) error
	changeStackMutex       sync.RWMutex
	changeStackArgsForCall []struct {
		arg1 string
		arg2 string
	}
	changeStackReturns struct {
		result1 error
	}
	changeStackReturnsOnCall map[int]struct {
		result1 error
	}
	InstallPluginStub        func(string) error
	installPluginMutex       sync.RWMutex
	installPluginArgsForCall []struct {
		arg1 string
	}
	installPluginReturns struct {
		result1 error
	}
	installPluginReturnsOnCall map[int]struct {
		result1 error
	}
	TargetStub        func(string, string) error
	targetMutex       sync.RWMutex
	targetArgsForCall []struct {
		arg1 string
		arg2 string
	}
	targetReturns struct {
		result1 error
	}
	targetReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExecutor) Api(arg1 string, arg2 bool) error {
	fake.apiMutex.Lock()
	ret, specificReturn := fake.apiReturnsOnCall[len(fake.apiArgsForCall)]
	fake.apiArgsForCall = append(fake.apiArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("Api", []interface{}{arg1, arg2})
	fake.apiMutex.Unlock()
	if fake.ApiStub != nil {
		return fake.ApiStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.apiReturns
	return fakeReturns.result1
}

func (fake *FakeExecutor) ApiCallCount() int {
	fake.apiMutex.RLock()
	defer fake.apiMutex.RUnlock()
	return len(fake.apiArgsForCall)
}

func (fake *FakeExecutor) ApiCalls(stub func(string, bool) error) {
	fake.apiMutex.Lock()
	defer fake.apiMutex.Unlock()
	fake.ApiStub = stub
}

func (fake *FakeExecutor) ApiArgsForCall(i int) (string, bool) {
	fake.apiMutex.RLock()
	defer fake.apiMutex.RUnlock()
	argsForCall := fake.apiArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExecutor) ApiReturns(result1 error) {
	fake.apiMutex.Lock()
	defer fake.apiMutex.Unlock()
	fake.ApiStub = nil
	fake.apiReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) ApiReturnsOnCall(i int, result1 error) {
	fake.apiMutex.Lock()
	defer fake.apiMutex.Unlock()
	fake.ApiStub = nil
	if fake.apiReturnsOnCall == nil {
		fake.apiReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.apiReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) Auth(arg1 string, arg2 string) error {
	fake.authMutex.Lock()
	ret, specificReturn := fake.authReturnsOnCall[len(fake.authArgsForCall)]
	fake.authArgsForCall = append(fake.authArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Auth", []interface{}{arg1, arg2})
	fake.authMutex.Unlock()
	if fake.AuthStub != nil {
		return fake.AuthStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.authReturns
	return fakeReturns.result1
}

func (fake *FakeExecutor) AuthCallCount() int {
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	return len(fake.authArgsForCall)
}

func (fake *FakeExecutor) AuthCalls(stub func(string, string) error) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = stub
}

func (fake *FakeExecutor) AuthArgsForCall(i int) (string, string) {
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	argsForCall := fake.authArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExecutor) AuthReturns(result1 error) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = nil
	fake.authReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) AuthReturnsOnCall(i int, result1 error) {
	fake.authMutex.Lock()
	defer fake.authMutex.Unlock()
	fake.AuthStub = nil
	if fake.authReturnsOnCall == nil {
		fake.authReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.authReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) Buffer() *bytes.Buffer {
	fake.bufferMutex.Lock()
	ret, specificReturn := fake.bufferReturnsOnCall[len(fake.bufferArgsForCall)]
	fake.bufferArgsForCall = append(fake.bufferArgsForCall, struct {
	}{})
	fake.recordInvocation("Buffer", []interface{}{})
	fake.bufferMutex.Unlock()
	if fake.BufferStub != nil {
		return fake.BufferStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bufferReturns
	return fakeReturns.result1
}

func (fake *FakeExecutor) BufferCallCount() int {
	fake.bufferMutex.RLock()
	defer fake.bufferMutex.RUnlock()
	return len(fake.bufferArgsForCall)
}

func (fake *FakeExecutor) BufferCalls(stub func() *bytes.Buffer) {
	fake.bufferMutex.Lock()
	defer fake.bufferMutex.Unlock()
	fake.BufferStub = stub
}

func (fake *FakeExecutor) BufferReturns(result1 *bytes.Buffer) {
	fake.bufferMutex.Lock()
	defer fake.bufferMutex.Unlock()
	fake.BufferStub = nil
	fake.bufferReturns = struct {
		result1 *bytes.Buffer
	}{result1}
}

func (fake *FakeExecutor) BufferReturnsOnCall(i int, result1 *bytes.Buffer) {
	fake.bufferMutex.Lock()
	defer fake.bufferMutex.Unlock()
	fake.BufferStub = nil
	if fake.bufferReturnsOnCall == nil {
		fake.bufferReturnsOnCall = make(map[int]struct {
			result1 *bytes.Buffer
		})
	}
	fake.bufferReturnsOnCall[i] = struct {
		result1 *bytes.Buffer
	}{result1}
}

func (fake *FakeExecutor) ChangeStack(arg1 string, arg2 string) error {
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
		return ret.result1
	}
	fakeReturns := fake.changeStackReturns
	return fakeReturns.result1
}

func (fake *FakeExecutor) ChangeStackCallCount() int {
	fake.changeStackMutex.RLock()
	defer fake.changeStackMutex.RUnlock()
	return len(fake.changeStackArgsForCall)
}

func (fake *FakeExecutor) ChangeStackCalls(stub func(string, string) error) {
	fake.changeStackMutex.Lock()
	defer fake.changeStackMutex.Unlock()
	fake.ChangeStackStub = stub
}

func (fake *FakeExecutor) ChangeStackArgsForCall(i int) (string, string) {
	fake.changeStackMutex.RLock()
	defer fake.changeStackMutex.RUnlock()
	argsForCall := fake.changeStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExecutor) ChangeStackReturns(result1 error) {
	fake.changeStackMutex.Lock()
	defer fake.changeStackMutex.Unlock()
	fake.ChangeStackStub = nil
	fake.changeStackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) ChangeStackReturnsOnCall(i int, result1 error) {
	fake.changeStackMutex.Lock()
	defer fake.changeStackMutex.Unlock()
	fake.ChangeStackStub = nil
	if fake.changeStackReturnsOnCall == nil {
		fake.changeStackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.changeStackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) InstallPlugin(arg1 string) error {
	fake.installPluginMutex.Lock()
	ret, specificReturn := fake.installPluginReturnsOnCall[len(fake.installPluginArgsForCall)]
	fake.installPluginArgsForCall = append(fake.installPluginArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("InstallPlugin", []interface{}{arg1})
	fake.installPluginMutex.Unlock()
	if fake.InstallPluginStub != nil {
		return fake.InstallPluginStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.installPluginReturns
	return fakeReturns.result1
}

func (fake *FakeExecutor) InstallPluginCallCount() int {
	fake.installPluginMutex.RLock()
	defer fake.installPluginMutex.RUnlock()
	return len(fake.installPluginArgsForCall)
}

func (fake *FakeExecutor) InstallPluginCalls(stub func(string) error) {
	fake.installPluginMutex.Lock()
	defer fake.installPluginMutex.Unlock()
	fake.InstallPluginStub = stub
}

func (fake *FakeExecutor) InstallPluginArgsForCall(i int) string {
	fake.installPluginMutex.RLock()
	defer fake.installPluginMutex.RUnlock()
	argsForCall := fake.installPluginArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeExecutor) InstallPluginReturns(result1 error) {
	fake.installPluginMutex.Lock()
	defer fake.installPluginMutex.Unlock()
	fake.InstallPluginStub = nil
	fake.installPluginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) InstallPluginReturnsOnCall(i int, result1 error) {
	fake.installPluginMutex.Lock()
	defer fake.installPluginMutex.Unlock()
	fake.InstallPluginStub = nil
	if fake.installPluginReturnsOnCall == nil {
		fake.installPluginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.installPluginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) Target(arg1 string, arg2 string) error {
	fake.targetMutex.Lock()
	ret, specificReturn := fake.targetReturnsOnCall[len(fake.targetArgsForCall)]
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Target", []interface{}{arg1, arg2})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetReturns
	return fakeReturns.result1
}

func (fake *FakeExecutor) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakeExecutor) TargetCalls(stub func(string, string) error) {
	fake.targetMutex.Lock()
	defer fake.targetMutex.Unlock()
	fake.TargetStub = stub
}

func (fake *FakeExecutor) TargetArgsForCall(i int) (string, string) {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	argsForCall := fake.targetArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExecutor) TargetReturns(result1 error) {
	fake.targetMutex.Lock()
	defer fake.targetMutex.Unlock()
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) TargetReturnsOnCall(i int, result1 error) {
	fake.targetMutex.Lock()
	defer fake.targetMutex.Unlock()
	fake.TargetStub = nil
	if fake.targetReturnsOnCall == nil {
		fake.targetReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.targetReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.apiMutex.RLock()
	defer fake.apiMutex.RUnlock()
	fake.authMutex.RLock()
	defer fake.authMutex.RUnlock()
	fake.bufferMutex.RLock()
	defer fake.bufferMutex.RUnlock()
	fake.changeStackMutex.RLock()
	defer fake.changeStackMutex.RUnlock()
	fake.installPluginMutex.RLock()
	defer fake.installPluginMutex.RUnlock()
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExecutor) recordInvocation(key string, args []interface{}) {
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

var _ cf.Executor = new(FakeExecutor)
