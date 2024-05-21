package testRunner_test

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/model"
	mockRepository "contest/internal/domain/repository/mocks"
	mockExecutor2 "contest/internal/domain/service/executor/mocks"
	"contest/internal/service/implementations/testRunner"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initMockTestRepository(cntrl *gomock.Controller) *mockRepository.MockTestRepository {
	repo := mockRepository.NewMockTestRepository(cntrl)
	repo.EXPECT().FindTestsByTaskID(gomock.Any(), 1).Return(
		[]model.Test{
			{
				ID:             1,
				TaskID:         1,
				Input:          "1",
				ExpectedResult: "1",
				Points:         1,
			},
			{
				ID:             2,
				TaskID:         1,
				Input:          "2",
				ExpectedResult: "4",
				Points:         1,
			},
			{
				ID:             3,
				TaskID:         1,
				Input:          "3",
				ExpectedResult: "9",
				Points:         1,
			},
		}, nil)

	return repo
}

func initMockTaskRepository(cntrl *gomock.Controller) *mockRepository.MockTaskRepository {
	repo := mockRepository.NewMockTaskRepository(cntrl)
	repo.EXPECT().FindTaskByID(gomock.Any(), 1).Return(model.Task{
		ID:          1,
		Title:       "Mock Task 1",
		Text:        "",
		MemoryLimit: 1024,
		TimeLimit:   3000,
	}, nil)

	return repo
}

func initMockExecutorFactory(cntrl *gomock.Controller) *mockExecutor2.MockExecutorFactory {
	mockFactory := mockExecutor2.NewMockExecutorFactory(cntrl)
	mockExecutor := mockExecutor2.NewMockExecutor(cntrl)
	mockExecutor.EXPECT().Execute("1", gomock.Any(), gomock.Any()).Return("1", nil)
	mockExecutor.EXPECT().Execute("2", gomock.Any(), gomock.Any()).Return("4", nil)
	mockExecutor.EXPECT().Execute("3", gomock.Any(), gomock.Any()).Return("9", nil)

	mockExecutor.EXPECT().Close()

	mockFactory.EXPECT().NewExecutor(gomock.Any(), gomock.Any()).Return(mockExecutor, nil)

	return mockFactory
}

func initMockLaunchRepository(cntrl *gomock.Controller) *mockRepository.MockLaunchRepository {
	repo := mockRepository.NewMockLaunchRepository(cntrl)
	repo.EXPECT().AddLaunch(gomock.Any(), gomock.Any()).Return(nil)

	return repo
}

func Test(t *testing.T) {
	cntrl := gomock.NewController(t)
	testRepository := initMockTestRepository(cntrl)
	executorFactory := initMockExecutorFactory(cntrl)
	launchRepository := initMockLaunchRepository(cntrl)
	taskRepository := initMockTaskRepository(cntrl)

	service := testRunner.NewTestRunnerService(executorFactory, testRepository, launchRepository, taskRepository)

	actual, err := service.RunTest(context.Background(), 1, 0, enum.CPP, "")
	assert.Nil(t, err)
	expected := model.TestsResult{
		ResultCode:  "SC",
		Description: "",
		Points:      3,
	}
	assert.Equal(t, expected, actual)
}
