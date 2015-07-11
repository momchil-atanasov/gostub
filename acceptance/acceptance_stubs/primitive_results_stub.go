package acceptance_stubs

import (
	sync "sync"
)

type PrimitiveResultsStub struct {
	UserStub        func() (result1 string, result2 int, result3 float32)
	userMutex       sync.RWMutex
	userArgsForCall []struct {
	}
	userReturns struct {
		result1 string
		result2 int
		result3 float32
	}
}

func (stub *PrimitiveResultsStub) User() (string, int, float32) {
	stub.userMutex.Lock()
	defer stub.userMutex.Unlock()
	stub.userArgsForCall = append(stub.userArgsForCall, struct {
	}{})
	if stub.UserStub != nil {
		return stub.UserStub()
	} else {
		return stub.userReturns.result1, stub.userReturns.result2, stub.userReturns.result3
	}
}
func (stub *PrimitiveResultsStub) UserCallCount() int {
	stub.userMutex.RLock()
	defer stub.userMutex.RUnlock()
	return len(stub.userArgsForCall)
}
func (stub *PrimitiveResultsStub) UserReturns(result1 string, result2 int, result3 float32) {
	stub.userMutex.Lock()
	defer stub.userMutex.Unlock()
	stub.userReturns = struct {
		result1 string
		result2 int
		result3 float32
	}{result1, result2, result3}
}