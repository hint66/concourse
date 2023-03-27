// Code generated by counterfeiter. DO NOT EDIT.
package pipelineserverfakes

import (
	"net/http"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeLogger struct {
	DebugStub        func(string, ...lager.Data)
	debugMutex       sync.RWMutex
	debugArgsForCall []struct {
		arg1 string
		arg2 []lager.Data
	}
	ErrorStub        func(string, error, ...lager.Data)
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
		arg1 string
		arg2 error
		arg3 []lager.Data
	}
	FatalStub        func(string, error, ...lager.Data)
	fatalMutex       sync.RWMutex
	fatalArgsForCall []struct {
		arg1 string
		arg2 error
		arg3 []lager.Data
	}
	InfoStub        func(string, ...lager.Data)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		arg1 string
		arg2 []lager.Data
	}
	RegisterSinkStub        func(lager.Sink)
	registerSinkMutex       sync.RWMutex
	registerSinkArgsForCall []struct {
		arg1 lager.Sink
	}
	SessionStub        func(string, ...lager.Data) lager.Logger
	sessionMutex       sync.RWMutex
	sessionArgsForCall []struct {
		arg1 string
		arg2 []lager.Data
	}
	sessionReturns struct {
		result1 lager.Logger
	}
	sessionReturnsOnCall map[int]struct {
		result1 lager.Logger
	}
	SessionNameStub        func() string
	sessionNameMutex       sync.RWMutex
	sessionNameArgsForCall []struct {
	}
	sessionNameReturns struct {
		result1 string
	}
	sessionNameReturnsOnCall map[int]struct {
		result1 string
	}
	WithDataStub        func(lager.Data) lager.Logger
	withDataMutex       sync.RWMutex
	withDataArgsForCall []struct {
		arg1 lager.Data
	}
	withDataReturns struct {
		result1 lager.Logger
	}
	withDataReturnsOnCall map[int]struct {
		result1 lager.Logger
	}
	WithTraceInfoStub        func(*http.Request) lager.Logger
	withTraceInfoMutex       sync.RWMutex
	withTraceInfoArgsForCall []struct {
		arg1 *http.Request
	}
	withTraceInfoReturns struct {
		result1 lager.Logger
	}
	withTraceInfoReturnsOnCall map[int]struct {
		result1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogger) Debug(arg1 string, arg2 ...lager.Data) {
	fake.debugMutex.Lock()
	fake.debugArgsForCall = append(fake.debugArgsForCall, struct {
		arg1 string
		arg2 []lager.Data
	}{arg1, arg2})
	stub := fake.DebugStub
	fake.recordInvocation("Debug", []interface{}{arg1, arg2})
	fake.debugMutex.Unlock()
	if stub != nil {
		fake.DebugStub(arg1, arg2...)
	}
}

func (fake *FakeLogger) DebugCallCount() int {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return len(fake.debugArgsForCall)
}

func (fake *FakeLogger) DebugCalls(stub func(string, ...lager.Data)) {
	fake.debugMutex.Lock()
	defer fake.debugMutex.Unlock()
	fake.DebugStub = stub
}

func (fake *FakeLogger) DebugArgsForCall(i int) (string, []lager.Data) {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	argsForCall := fake.debugArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) Error(arg1 string, arg2 error, arg3 ...lager.Data) {
	fake.errorMutex.Lock()
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
		arg1 string
		arg2 error
		arg3 []lager.Data
	}{arg1, arg2, arg3})
	stub := fake.ErrorStub
	fake.recordInvocation("Error", []interface{}{arg1, arg2, arg3})
	fake.errorMutex.Unlock()
	if stub != nil {
		fake.ErrorStub(arg1, arg2, arg3...)
	}
}

func (fake *FakeLogger) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeLogger) ErrorCalls(stub func(string, error, ...lager.Data)) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = stub
}

func (fake *FakeLogger) ErrorArgsForCall(i int) (string, error, []lager.Data) {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	argsForCall := fake.errorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLogger) Fatal(arg1 string, arg2 error, arg3 ...lager.Data) {
	fake.fatalMutex.Lock()
	fake.fatalArgsForCall = append(fake.fatalArgsForCall, struct {
		arg1 string
		arg2 error
		arg3 []lager.Data
	}{arg1, arg2, arg3})
	stub := fake.FatalStub
	fake.recordInvocation("Fatal", []interface{}{arg1, arg2, arg3})
	fake.fatalMutex.Unlock()
	if stub != nil {
		fake.FatalStub(arg1, arg2, arg3...)
	}
}

func (fake *FakeLogger) FatalCallCount() int {
	fake.fatalMutex.RLock()
	defer fake.fatalMutex.RUnlock()
	return len(fake.fatalArgsForCall)
}

func (fake *FakeLogger) FatalCalls(stub func(string, error, ...lager.Data)) {
	fake.fatalMutex.Lock()
	defer fake.fatalMutex.Unlock()
	fake.FatalStub = stub
}

func (fake *FakeLogger) FatalArgsForCall(i int) (string, error, []lager.Data) {
	fake.fatalMutex.RLock()
	defer fake.fatalMutex.RUnlock()
	argsForCall := fake.fatalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLogger) Info(arg1 string, arg2 ...lager.Data) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		arg1 string
		arg2 []lager.Data
	}{arg1, arg2})
	stub := fake.InfoStub
	fake.recordInvocation("Info", []interface{}{arg1, arg2})
	fake.infoMutex.Unlock()
	if stub != nil {
		fake.InfoStub(arg1, arg2...)
	}
}

func (fake *FakeLogger) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeLogger) InfoCalls(stub func(string, ...lager.Data)) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = stub
}

func (fake *FakeLogger) InfoArgsForCall(i int) (string, []lager.Data) {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	argsForCall := fake.infoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) RegisterSink(arg1 lager.Sink) {
	fake.registerSinkMutex.Lock()
	fake.registerSinkArgsForCall = append(fake.registerSinkArgsForCall, struct {
		arg1 lager.Sink
	}{arg1})
	stub := fake.RegisterSinkStub
	fake.recordInvocation("RegisterSink", []interface{}{arg1})
	fake.registerSinkMutex.Unlock()
	if stub != nil {
		fake.RegisterSinkStub(arg1)
	}
}

func (fake *FakeLogger) RegisterSinkCallCount() int {
	fake.registerSinkMutex.RLock()
	defer fake.registerSinkMutex.RUnlock()
	return len(fake.registerSinkArgsForCall)
}

func (fake *FakeLogger) RegisterSinkCalls(stub func(lager.Sink)) {
	fake.registerSinkMutex.Lock()
	defer fake.registerSinkMutex.Unlock()
	fake.RegisterSinkStub = stub
}

func (fake *FakeLogger) RegisterSinkArgsForCall(i int) lager.Sink {
	fake.registerSinkMutex.RLock()
	defer fake.registerSinkMutex.RUnlock()
	argsForCall := fake.registerSinkArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLogger) Session(arg1 string, arg2 ...lager.Data) lager.Logger {
	fake.sessionMutex.Lock()
	ret, specificReturn := fake.sessionReturnsOnCall[len(fake.sessionArgsForCall)]
	fake.sessionArgsForCall = append(fake.sessionArgsForCall, struct {
		arg1 string
		arg2 []lager.Data
	}{arg1, arg2})
	stub := fake.SessionStub
	fakeReturns := fake.sessionReturns
	fake.recordInvocation("Session", []interface{}{arg1, arg2})
	fake.sessionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLogger) SessionCallCount() int {
	fake.sessionMutex.RLock()
	defer fake.sessionMutex.RUnlock()
	return len(fake.sessionArgsForCall)
}

func (fake *FakeLogger) SessionCalls(stub func(string, ...lager.Data) lager.Logger) {
	fake.sessionMutex.Lock()
	defer fake.sessionMutex.Unlock()
	fake.SessionStub = stub
}

func (fake *FakeLogger) SessionArgsForCall(i int) (string, []lager.Data) {
	fake.sessionMutex.RLock()
	defer fake.sessionMutex.RUnlock()
	argsForCall := fake.sessionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) SessionReturns(result1 lager.Logger) {
	fake.sessionMutex.Lock()
	defer fake.sessionMutex.Unlock()
	fake.SessionStub = nil
	fake.sessionReturns = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeLogger) SessionReturnsOnCall(i int, result1 lager.Logger) {
	fake.sessionMutex.Lock()
	defer fake.sessionMutex.Unlock()
	fake.SessionStub = nil
	if fake.sessionReturnsOnCall == nil {
		fake.sessionReturnsOnCall = make(map[int]struct {
			result1 lager.Logger
		})
	}
	fake.sessionReturnsOnCall[i] = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeLogger) SessionName() string {
	fake.sessionNameMutex.Lock()
	ret, specificReturn := fake.sessionNameReturnsOnCall[len(fake.sessionNameArgsForCall)]
	fake.sessionNameArgsForCall = append(fake.sessionNameArgsForCall, struct {
	}{})
	stub := fake.SessionNameStub
	fakeReturns := fake.sessionNameReturns
	fake.recordInvocation("SessionName", []interface{}{})
	fake.sessionNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLogger) SessionNameCallCount() int {
	fake.sessionNameMutex.RLock()
	defer fake.sessionNameMutex.RUnlock()
	return len(fake.sessionNameArgsForCall)
}

func (fake *FakeLogger) SessionNameCalls(stub func() string) {
	fake.sessionNameMutex.Lock()
	defer fake.sessionNameMutex.Unlock()
	fake.SessionNameStub = stub
}

func (fake *FakeLogger) SessionNameReturns(result1 string) {
	fake.sessionNameMutex.Lock()
	defer fake.sessionNameMutex.Unlock()
	fake.SessionNameStub = nil
	fake.sessionNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeLogger) SessionNameReturnsOnCall(i int, result1 string) {
	fake.sessionNameMutex.Lock()
	defer fake.sessionNameMutex.Unlock()
	fake.SessionNameStub = nil
	if fake.sessionNameReturnsOnCall == nil {
		fake.sessionNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sessionNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeLogger) WithData(arg1 lager.Data) lager.Logger {
	fake.withDataMutex.Lock()
	ret, specificReturn := fake.withDataReturnsOnCall[len(fake.withDataArgsForCall)]
	fake.withDataArgsForCall = append(fake.withDataArgsForCall, struct {
		arg1 lager.Data
	}{arg1})
	stub := fake.WithDataStub
	fakeReturns := fake.withDataReturns
	fake.recordInvocation("WithData", []interface{}{arg1})
	fake.withDataMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLogger) WithDataCallCount() int {
	fake.withDataMutex.RLock()
	defer fake.withDataMutex.RUnlock()
	return len(fake.withDataArgsForCall)
}

func (fake *FakeLogger) WithDataCalls(stub func(lager.Data) lager.Logger) {
	fake.withDataMutex.Lock()
	defer fake.withDataMutex.Unlock()
	fake.WithDataStub = stub
}

func (fake *FakeLogger) WithDataArgsForCall(i int) lager.Data {
	fake.withDataMutex.RLock()
	defer fake.withDataMutex.RUnlock()
	argsForCall := fake.withDataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLogger) WithDataReturns(result1 lager.Logger) {
	fake.withDataMutex.Lock()
	defer fake.withDataMutex.Unlock()
	fake.WithDataStub = nil
	fake.withDataReturns = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeLogger) WithDataReturnsOnCall(i int, result1 lager.Logger) {
	fake.withDataMutex.Lock()
	defer fake.withDataMutex.Unlock()
	fake.WithDataStub = nil
	if fake.withDataReturnsOnCall == nil {
		fake.withDataReturnsOnCall = make(map[int]struct {
			result1 lager.Logger
		})
	}
	fake.withDataReturnsOnCall[i] = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeLogger) WithTraceInfo(arg1 *http.Request) lager.Logger {
	fake.withTraceInfoMutex.Lock()
	ret, specificReturn := fake.withTraceInfoReturnsOnCall[len(fake.withTraceInfoArgsForCall)]
	fake.withTraceInfoArgsForCall = append(fake.withTraceInfoArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	stub := fake.WithTraceInfoStub
	fakeReturns := fake.withTraceInfoReturns
	fake.recordInvocation("WithTraceInfo", []interface{}{arg1})
	fake.withTraceInfoMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeLogger) WithTraceInfoCallCount() int {
	fake.withTraceInfoMutex.RLock()
	defer fake.withTraceInfoMutex.RUnlock()
	return len(fake.withTraceInfoArgsForCall)
}

func (fake *FakeLogger) WithTraceInfoCalls(stub func(*http.Request) lager.Logger) {
	fake.withTraceInfoMutex.Lock()
	defer fake.withTraceInfoMutex.Unlock()
	fake.WithTraceInfoStub = stub
}

func (fake *FakeLogger) WithTraceInfoArgsForCall(i int) *http.Request {
	fake.withTraceInfoMutex.RLock()
	defer fake.withTraceInfoMutex.RUnlock()
	argsForCall := fake.withTraceInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLogger) WithTraceInfoReturns(result1 lager.Logger) {
	fake.withTraceInfoMutex.Lock()
	defer fake.withTraceInfoMutex.Unlock()
	fake.WithTraceInfoStub = nil
	fake.withTraceInfoReturns = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeLogger) WithTraceInfoReturnsOnCall(i int, result1 lager.Logger) {
	fake.withTraceInfoMutex.Lock()
	defer fake.withTraceInfoMutex.Unlock()
	fake.WithTraceInfoStub = nil
	if fake.withTraceInfoReturnsOnCall == nil {
		fake.withTraceInfoReturnsOnCall = make(map[int]struct {
			result1 lager.Logger
		})
	}
	fake.withTraceInfoReturnsOnCall[i] = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	fake.fatalMutex.RLock()
	defer fake.fatalMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.registerSinkMutex.RLock()
	defer fake.registerSinkMutex.RUnlock()
	fake.sessionMutex.RLock()
	defer fake.sessionMutex.RUnlock()
	fake.sessionNameMutex.RLock()
	defer fake.sessionNameMutex.RUnlock()
	fake.withDataMutex.RLock()
	defer fake.withDataMutex.RUnlock()
	fake.withTraceInfoMutex.RLock()
	defer fake.withTraceInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogger) recordInvocation(key string, args []interface{}) {
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

var _ lager.Logger = new(FakeLogger)
