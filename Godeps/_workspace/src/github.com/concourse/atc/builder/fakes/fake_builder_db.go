// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/builder"
	"github.com/concourse/atc/builds"
)

type FakeBuilderDB struct {
	CreateOneOffBuildStub        func() (builds.Build, error)
	createOneOffBuildMutex       sync.RWMutex
	createOneOffBuildArgsForCall []struct{}
	createOneOffBuildReturns struct {
		result1 builds.Build
		result2 error
	}
	StartBuildStub        func(buildID int, abortURL string) (bool, error)
	startBuildMutex       sync.RWMutex
	startBuildArgsForCall []struct {
		buildID  int
		abortURL string
	}
	startBuildReturns struct {
		result1 bool
		result2 error
	}
}

func (fake *FakeBuilderDB) CreateOneOffBuild() (builds.Build, error) {
	fake.createOneOffBuildMutex.Lock()
	fake.createOneOffBuildArgsForCall = append(fake.createOneOffBuildArgsForCall, struct{}{})
	fake.createOneOffBuildMutex.Unlock()
	if fake.CreateOneOffBuildStub != nil {
		return fake.CreateOneOffBuildStub()
	} else {
		return fake.createOneOffBuildReturns.result1, fake.createOneOffBuildReturns.result2
	}
}

func (fake *FakeBuilderDB) CreateOneOffBuildCallCount() int {
	fake.createOneOffBuildMutex.RLock()
	defer fake.createOneOffBuildMutex.RUnlock()
	return len(fake.createOneOffBuildArgsForCall)
}

func (fake *FakeBuilderDB) CreateOneOffBuildReturns(result1 builds.Build, result2 error) {
	fake.CreateOneOffBuildStub = nil
	fake.createOneOffBuildReturns = struct {
		result1 builds.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeBuilderDB) StartBuild(buildID int, abortURL string) (bool, error) {
	fake.startBuildMutex.Lock()
	fake.startBuildArgsForCall = append(fake.startBuildArgsForCall, struct {
		buildID  int
		abortURL string
	}{buildID, abortURL})
	fake.startBuildMutex.Unlock()
	if fake.StartBuildStub != nil {
		return fake.StartBuildStub(buildID, abortURL)
	} else {
		return fake.startBuildReturns.result1, fake.startBuildReturns.result2
	}
}

func (fake *FakeBuilderDB) StartBuildCallCount() int {
	fake.startBuildMutex.RLock()
	defer fake.startBuildMutex.RUnlock()
	return len(fake.startBuildArgsForCall)
}

func (fake *FakeBuilderDB) StartBuildArgsForCall(i int) (int, string) {
	fake.startBuildMutex.RLock()
	defer fake.startBuildMutex.RUnlock()
	return fake.startBuildArgsForCall[i].buildID, fake.startBuildArgsForCall[i].abortURL
}

func (fake *FakeBuilderDB) StartBuildReturns(result1 bool, result2 error) {
	fake.StartBuildStub = nil
	fake.startBuildReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

var _ builder.BuilderDB = new(FakeBuilderDB)