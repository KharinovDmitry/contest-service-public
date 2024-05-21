package service

import (
	executor "contest/internal/domain/service/executor"
	logger "contest/internal/domain/service/logger"
	testRunner "contest/internal/domain/service/testRunner"

	executor2 "contest/internal/service/implementations/executor"
	logger2 "contest/internal/service/implementations/logger"
	testRunner2 "contest/internal/service/implementations/testRunner"

	"contest/internal/storage"
	"contest/lib/adapter/os/linux"
)

type Manager struct {
	Logger          logger.Logger
	ExecutorFactory executor.ExecutorFactory
	TestRunner      testRunner.TestRunner
}

func NewServiceManager() *Manager {
	return &Manager{}
}

func (m *Manager) Init(store *storage.Storage, env string) error {
	log, err := logger2.NewLogger(env)
	if err != nil {
		return err
	}
	m.Logger = log

	adapter := linux.NewLinuxAdapter()
	factory := executor2.NewExecutorFactory(adapter)

	testRunner := testRunner2.NewTestRunnerService(factory, store.TestRepository, store.LaunchRepository, store.TaskRepository)
	m.TestRunner = testRunner

	return nil
}
