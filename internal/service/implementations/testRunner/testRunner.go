package testRunner

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/model"
	"contest/internal/domain/repository"
	"contest/internal/domain/service/executor"
	executor2 "contest/internal/domain/service/executor"
	"contest/internal/domain/service/testRunner"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type TestRunnerService struct {
	codeRunnerFactory executor.ExecutorFactory
	testRepository    repository.TestRepository
	taskRepository    repository.TaskRepository
	launchRepository  repository.LaunchRepository
}

func NewTestRunnerService(codeRunnerFactory executor.ExecutorFactory, testRepository repository.TestRepository, launchRepository repository.LaunchRepository, taskRepository repository.TaskRepository) *TestRunnerService {
	return &TestRunnerService{
		codeRunnerFactory: codeRunnerFactory,
		testRepository:    testRepository,
		launchRepository:  launchRepository,
		taskRepository:    taskRepository,
	}
}

func (s *TestRunnerService) RunTest(ctx context.Context, taskID int, userId int, language enum.Language, code string) (model.TestsResult, error) {
	program, err := s.codeRunnerFactory.NewExecutor(code, language)
	if err != nil {
		if errors.Is(err, executor2.CompileError) {
			res := model.TestsResult{
				ResultCode:  enum.CompileErrorCode,
				Description: err.Error(),
				Points:      0,
			}
			err = s.launchRepository.AddLaunch(ctx, model.Launch{
				UserID:      userId,
				ContestID:   taskID,
				Date:        time.Now(),
				Code:        code,
				TestsResult: res,
			})
			return res, nil
		}
		return model.TestsResult{}, fmt.Errorf("In RunTestService(RunTest): %w", err)
	}
	defer program.Close()

	task, err := s.taskRepository.FindTaskByID(ctx, taskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.TestsResult{}, fmt.Errorf("%w with id: %d", testRunner.ErrTaskNotFound, taskID)
		}
		return model.TestsResult{}, fmt.Errorf("In RunTestService(RunTest): %w", err)
	}

	tests, err := s.testRepository.FindTestsByTaskID(ctx, taskID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.TestsResult{}, fmt.Errorf("%w for task id: %d", testRunner.ErrTestsNotFound, taskID)
		}
		return model.TestsResult{}, fmt.Errorf("In RunTestService(RunTest): %w", err)
	}

	res := model.TestsResult{
		ResultCode:  enum.SuccessCode,
		Description: "",
		Points:      0,
	}
	for i, test := range tests {
		output, err := program.Execute(test.Input, task.MemoryLimit, task.TimeLimit)
		if err != nil {
			if errors.Is(err, executor2.TimeLimitError) {
				res = model.TestsResult{
					ResultCode:  enum.TimeLimitCode,
					Description: fmt.Sprintf("Test failed: %d. Time limit error", i),
					Points:      res.Points,
				}
				break
			}
			if errors.Is(err, executor2.RuntimeError) {
				res = model.TestsResult{
					ResultCode:  enum.RuntimeErrorCode,
					Description: fmt.Sprintf("Test failed: %d. Description: %s Output: %s", i, err.Error(), output),
					Points:      res.Points,
				}
				break
			}
			return model.TestsResult{}, fmt.Errorf("In TestService(RunTests): %w", err)
		}
		if strings.Replace(output, "\n", "", 1) != test.ExpectedResult {
			res = model.TestsResult{
				ResultCode:  enum.IncorrectAnswerCode,
				Description: fmt.Sprintf("Test failed: %d Excepted: %s Received: %s", i, test.ExpectedResult, output),
				Points:      res.Points,
			}
			break
		}
		res.Points += test.Points
	}

	err = s.launchRepository.AddLaunch(ctx, model.Launch{
		UserID:      userId,
		ContestID:   taskID,
		Date:        time.Now(),
		Code:        code,
		TestsResult: res,
	})
	if err != nil {
		return res, fmt.Errorf("In TestService(RunTests): %w", err)
	}
	return res, nil

}
