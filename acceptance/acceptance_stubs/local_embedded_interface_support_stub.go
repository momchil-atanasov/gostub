package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance"
)

type LocalEmbeddedInterfaceSupportStub struct {
	MethodStub        func(arg1 int) (result1 int)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 int
	}
	methodReturns struct {
		result1 int
	}
	ScheduleStub        func(arg1 string, arg2 alias1.Customer) (result1 int)
	scheduleMutex       sync.RWMutex
	scheduleArgsForCall []struct {
		arg1 string
		arg2 alias1.Customer
	}
	scheduleReturns struct {
		result1 int
	}
}

func (stub *LocalEmbeddedInterfaceSupportStub) Method(arg1 int) int {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 int
	}{arg1})
	if stub.MethodStub != nil {
		return stub.MethodStub(arg1)
	} else {
		return stub.methodReturns.result1
	}
}
func (stub *LocalEmbeddedInterfaceSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *LocalEmbeddedInterfaceSupportStub) MethodArgsForCall(index int) int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *LocalEmbeddedInterfaceSupportStub) MethodReturns(result1 int) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 int
	}{result1}
}
func (stub *LocalEmbeddedInterfaceSupportStub) Schedule(arg1 string, arg2 alias1.Customer) int {
	stub.scheduleMutex.Lock()
	defer stub.scheduleMutex.Unlock()
	stub.scheduleArgsForCall = append(stub.scheduleArgsForCall, struct {
		arg1 string
		arg2 alias1.Customer
	}{arg1, arg2})
	if stub.ScheduleStub != nil {
		return stub.ScheduleStub(arg1, arg2)
	} else {
		return stub.scheduleReturns.result1
	}
}
func (stub *LocalEmbeddedInterfaceSupportStub) ScheduleCallCount() int {
	stub.scheduleMutex.RLock()
	defer stub.scheduleMutex.RUnlock()
	return len(stub.scheduleArgsForCall)
}
func (stub *LocalEmbeddedInterfaceSupportStub) ScheduleArgsForCall(index int) (string, alias1.Customer) {
	stub.scheduleMutex.RLock()
	defer stub.scheduleMutex.RUnlock()
	return stub.scheduleArgsForCall[index].arg1, stub.scheduleArgsForCall[index].arg2
}
func (stub *LocalEmbeddedInterfaceSupportStub) ScheduleReturns(result1 int) {
	stub.scheduleMutex.Lock()
	defer stub.scheduleMutex.Unlock()
	stub.scheduleReturns = struct {
		result1 int
	}{result1}
}