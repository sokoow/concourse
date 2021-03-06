// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db"
)

type FakeConflict struct {
	ConflictStub        func() string
	conflictMutex       sync.RWMutex
	conflictArgsForCall []struct {
	}
	conflictReturns struct {
		result1 string
	}
	conflictReturnsOnCall map[int]struct {
		result1 string
	}
	ErrorStub        func() string
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
	}
	errorReturns struct {
		result1 string
	}
	errorReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConflict) Conflict() string {
	fake.conflictMutex.Lock()
	ret, specificReturn := fake.conflictReturnsOnCall[len(fake.conflictArgsForCall)]
	fake.conflictArgsForCall = append(fake.conflictArgsForCall, struct {
	}{})
	fake.recordInvocation("Conflict", []interface{}{})
	fake.conflictMutex.Unlock()
	if fake.ConflictStub != nil {
		return fake.ConflictStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.conflictReturns
	return fakeReturns.result1
}

func (fake *FakeConflict) ConflictCallCount() int {
	fake.conflictMutex.RLock()
	defer fake.conflictMutex.RUnlock()
	return len(fake.conflictArgsForCall)
}

func (fake *FakeConflict) ConflictCalls(stub func() string) {
	fake.conflictMutex.Lock()
	defer fake.conflictMutex.Unlock()
	fake.ConflictStub = stub
}

func (fake *FakeConflict) ConflictReturns(result1 string) {
	fake.conflictMutex.Lock()
	defer fake.conflictMutex.Unlock()
	fake.ConflictStub = nil
	fake.conflictReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConflict) ConflictReturnsOnCall(i int, result1 string) {
	fake.conflictMutex.Lock()
	defer fake.conflictMutex.Unlock()
	fake.ConflictStub = nil
	if fake.conflictReturnsOnCall == nil {
		fake.conflictReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.conflictReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConflict) Error() string {
	fake.errorMutex.Lock()
	ret, specificReturn := fake.errorReturnsOnCall[len(fake.errorArgsForCall)]
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
	}{})
	fake.recordInvocation("Error", []interface{}{})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		return fake.ErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.errorReturns
	return fakeReturns.result1
}

func (fake *FakeConflict) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeConflict) ErrorCalls(stub func() string) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = stub
}

func (fake *FakeConflict) ErrorReturns(result1 string) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = nil
	fake.errorReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConflict) ErrorReturnsOnCall(i int, result1 string) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = nil
	if fake.errorReturnsOnCall == nil {
		fake.errorReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.errorReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConflict) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.conflictMutex.RLock()
	defer fake.conflictMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConflict) recordInvocation(key string, args []interface{}) {
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

var _ db.Conflict = new(FakeConflict)
