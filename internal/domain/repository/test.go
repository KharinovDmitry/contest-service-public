package repository

import (
	"contest/internal/domain/model"
	"context"
)

type TestRepository interface {
	AddTest(ctx context.Context, taskID int, input string, expectedResult string, points int) error
	DeleteTest(ctx context.Context, id int) error
	UpdateTest(ctx context.Context, id int, newItem model.Test) error
	GetTests(ctx context.Context) ([]model.Test, error)
	FindTestByID(ctx context.Context, id int) (model.Test, error)
	FindTestsByTaskID(ctx context.Context, taskID int) ([]model.Test, error)
}
