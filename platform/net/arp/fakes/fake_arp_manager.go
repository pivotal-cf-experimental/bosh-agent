// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/platform/net/arp"
)

type FakeManager struct {
	DeleteStub        func(string)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
}

func (fake *FakeManager) Delete(arg1 string) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		fake.DeleteStub(arg1)
	}
}

func (fake *FakeManager) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeManager) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

var _ arp.Manager = new(FakeManager)
