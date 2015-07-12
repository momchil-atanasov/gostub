package acceptance_stubs

import (
	sync "sync"
)

type ReusedResultsStub struct {
	FullNameStub        func() (result1 string, result2 string)
	fullNameMutex       sync.RWMutex
	fullNameArgsForCall []struct {
	}
	fullNameReturns struct {
		result1 string
		result2 string
	}
}

func (stub *ReusedResultsStub) FullName() (string, string) {
	stub.fullNameMutex.Lock()
	defer stub.fullNameMutex.Unlock()
	stub.fullNameArgsForCall = append(stub.fullNameArgsForCall, struct {
	}{})
	if stub.FullNameStub != nil {
		return stub.FullNameStub()
	} else {
		return stub.fullNameReturns.result1, stub.fullNameReturns.result2
	}
}
func (stub *ReusedResultsStub) FullNameCallCount() int {
	stub.fullNameMutex.RLock()
	defer stub.fullNameMutex.RUnlock()
	return len(stub.fullNameArgsForCall)
}
func (stub *ReusedResultsStub) FullNameReturns(result1 string, result2 string) {
	stub.fullNameMutex.Lock()
	defer stub.fullNameMutex.Unlock()
	stub.fullNameReturns = struct {
		result1 string
		result2 string
	}{result1, result2}
}
