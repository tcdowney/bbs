// This file was generated by counterfeiter
package fake_controllers

import (
	"sync"
	"time"

	"code.cloudfoundry.org/bbs/converger"
	"code.cloudfoundry.org/lager"
)

type FakeTaskController struct {
	ConvergeTasksStub        func(logger lager.Logger, kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration time.Duration) error
	convergeTasksMutex       sync.RWMutex
	convergeTasksArgsForCall []struct {
		logger                      lager.Logger
		kickTaskDuration            time.Duration
		expirePendingTaskDuration   time.Duration
		expireCompletedTaskDuration time.Duration
	}
	convergeTasksReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskController) ConvergeTasks(logger lager.Logger, kickTaskDuration time.Duration, expirePendingTaskDuration time.Duration, expireCompletedTaskDuration time.Duration) error {
	fake.convergeTasksMutex.Lock()
	fake.convergeTasksArgsForCall = append(fake.convergeTasksArgsForCall, struct {
		logger                      lager.Logger
		kickTaskDuration            time.Duration
		expirePendingTaskDuration   time.Duration
		expireCompletedTaskDuration time.Duration
	}{logger, kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration})
	fake.recordInvocation("ConvergeTasks", []interface{}{logger, kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration})
	fake.convergeTasksMutex.Unlock()
	if fake.ConvergeTasksStub != nil {
		return fake.ConvergeTasksStub(logger, kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration)
	} else {
		return fake.convergeTasksReturns.result1
	}
}

func (fake *FakeTaskController) ConvergeTasksCallCount() int {
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	return len(fake.convergeTasksArgsForCall)
}

func (fake *FakeTaskController) ConvergeTasksArgsForCall(i int) (lager.Logger, time.Duration, time.Duration, time.Duration) {
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	return fake.convergeTasksArgsForCall[i].logger, fake.convergeTasksArgsForCall[i].kickTaskDuration, fake.convergeTasksArgsForCall[i].expirePendingTaskDuration, fake.convergeTasksArgsForCall[i].expireCompletedTaskDuration
}

func (fake *FakeTaskController) ConvergeTasksReturns(result1 error) {
	fake.ConvergeTasksStub = nil
	fake.convergeTasksReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTaskController) recordInvocation(key string, args []interface{}) {
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

var _ converger.TaskController = new(FakeTaskController)
