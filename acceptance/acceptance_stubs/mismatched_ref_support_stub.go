package acceptance_stubs

import (
	sync "sync"

	alias1 "github.com/momchil-atanasov/gostub/acceptance/mismatch"
)

type MismatchedRefSupportStub struct {
	MethodStub        func(arg1 alias1.Job) (result1 alias1.Job)
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 alias1.Job
	}
	methodReturns struct {
		result1 alias1.Job
	}
}

func (stub *MismatchedRefSupportStub) Method(arg1 alias1.Job) alias1.Job {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodArgsForCall = append(stub.methodArgsForCall, struct {
		arg1 alias1.Job
	}{arg1})
	if stub.MethodStub != nil {
		return stub.MethodStub(arg1)
	} else {
		return stub.methodReturns.result1
	}
}
func (stub *MismatchedRefSupportStub) MethodCallCount() int {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return len(stub.methodArgsForCall)
}
func (stub *MismatchedRefSupportStub) MethodArgsForCall(index int) alias1.Job {
	stub.methodMutex.RLock()
	defer stub.methodMutex.RUnlock()
	return stub.methodArgsForCall[index].arg1
}
func (stub *MismatchedRefSupportStub) MethodReturns(result1 alias1.Job) {
	stub.methodMutex.Lock()
	defer stub.methodMutex.Unlock()
	stub.methodReturns = struct {
		result1 alias1.Job
	}{result1}
}
