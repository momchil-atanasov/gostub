package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/external/external_dup"
)

type ExternalEmbeddedInterfaceSupportStub struct {
	RunStub        func(arg1 alias1.Address) (result1 error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 alias1.Address
	}
	runReturns struct {
		result1 error
	}
	MethodStub        func(arg1 int) (result1 int)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 int
	}
	methodReturns struct {
		result1 int
	}
}

func (stub *ExternalEmbeddedInterfaceSupportStub) Run(arg1 alias1.Address) error {
	stub.runMutex.Lock()
	defer stub.runMutex.Unlock()
	stub.runArgsForCall = append(stub.runArgsForCall, struct {
		arg1 alias1.Address
	}{arg1})
	if stub.RunStub != nil {
		return stub.RunStub(arg1)
	} else {
		return stub.runReturns.result1
	}
}
func (stub *ExternalEmbeddedInterfaceSupportStub) RunCallCount() int {
	stub.runMutex.RLock()
	defer stub.runMutex.RUnlock()
	return len(stub.runArgsForCall)
}
func (stub *ExternalEmbeddedInterfaceSupportStub) RunArgsForCall(index int) alias1.Address {
	stub.runMutex.RLock()
	defer stub.runMutex.RUnlock()
	return stub.runArgsForCall[index].arg1
}
func (stub *ExternalEmbeddedInterfaceSupportStub) RunReturns(result1 error) {
	stub.runMutex.Lock()
	defer stub.runMutex.Unlock()
	stub.runReturns = struct {
		result1 error
	}{result1}
}
func (stub *ExternalEmbeddedInterfaceSupportStub) Method(arg1 int) int {
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
func (stub *ExternalEmbeddedInterfaceSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *ExternalEmbeddedInterfaceSupportStub) MethodArgsForCall(index int) int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *ExternalEmbeddedInterfaceSupportStub) MethodReturns(result1 int) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 int
	}{result1}
}
