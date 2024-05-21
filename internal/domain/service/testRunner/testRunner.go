package testRunner

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/model"
	"context"
	"errors"
)

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrTestsNotFound = errors.New("tests not found")
)

type TestRunner interface {
	RunTest(ctx context.Context, taskID int, userId int, language enum.Language, code string) (model.TestsResult, error)
}
