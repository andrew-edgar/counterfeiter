// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/locator"
)

type FakeInterfaceLocator struct {
	GetInterfacesFromFilePathStub        func(string) []string
	getInterfacesFromFilePathMutex       sync.RWMutex
	getInterfacesFromFilePathArgsForCall []struct {
		arg1 string
	}
	getInterfacesFromFilePathReturns struct {
		result1 []string
	}
}

func (fake *FakeInterfaceLocator) GetInterfacesFromFilePath(arg1 string) []string {
	fake.getInterfacesFromFilePathMutex.Lock()
	fake.getInterfacesFromFilePathArgsForCall = append(fake.getInterfacesFromFilePathArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getInterfacesFromFilePathMutex.Unlock()
	if fake.GetInterfacesFromFilePathStub != nil {
		return fake.GetInterfacesFromFilePathStub(arg1)
	} else {
		return fake.getInterfacesFromFilePathReturns.result1
	}
}

func (fake *FakeInterfaceLocator) GetInterfacesFromFilePathCallCount() int {
	fake.getInterfacesFromFilePathMutex.RLock()
	defer fake.getInterfacesFromFilePathMutex.RUnlock()
	return len(fake.getInterfacesFromFilePathArgsForCall)
}

func (fake *FakeInterfaceLocator) GetInterfacesFromFilePathArgsForCall(i int) string {
	fake.getInterfacesFromFilePathMutex.RLock()
	defer fake.getInterfacesFromFilePathMutex.RUnlock()
	return fake.getInterfacesFromFilePathArgsForCall[i].arg1
}

func (fake *FakeInterfaceLocator) GetInterfacesFromFilePathReturns(result1 []string) {
	fake.GetInterfacesFromFilePathStub = nil
	fake.getInterfacesFromFilePathReturns = struct {
		result1 []string
	}{result1}
}

var _ locator.InterfaceLocator = new(FakeInterfaceLocator)
