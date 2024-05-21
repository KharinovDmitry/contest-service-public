package implementations

import (
	"contest/internal/domain/model"
	"contest/internal/storage/dbModel"
	"contest/lib/adapter/db"
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	defaultExpiration = 5 * time.Minute
	cleanupInterval   = 10 * time.Minute
)

type TestRepository struct {
	db    db.DBAdapter
	cache *cache.Cache
}

func NewTestRepository(db db.DBAdapter) *TestRepository {
	cache := cache.New(defaultExpiration, cleanupInterval)
	return &TestRepository{
		db:    db,
		cache: cache,
	}
}

func (r *TestRepository) AddTest(ctx context.Context, taskID int, input string, expectedResult string, points int) error {
	sql := `INSERT INTO tests(task_id, input, expected_result, points) VALUES ($1, $2, $3, $4)`

	err := r.db.Execute(ctx, sql, taskID, input, expectedResult, points)
	if err != nil {
		err = fmt.Errorf("In TestRepository(AddItem): %w", err)
	}

	cacheKey := fmt.Sprintf("tests_taskID_%d", taskID)
	r.cache.Delete("tests")
	r.cache.Delete(cacheKey)
	return nil
}

func (r *TestRepository) DeleteTest(ctx context.Context, id int) error {
	sql := `DELETE FROM tests WHERE id = $1`

	err := r.db.Execute(ctx, sql, id)
	if err != nil {
		err = fmt.Errorf("In TestRepository(DeleteItem): %w", err)
		return err
	}
	r.cache.DeleteExpired()
	return nil
}

func (r *TestRepository) UpdateTest(ctx context.Context, id int, newItem model.Test) error {
	sql := `UPDATE tests SET id=$1,task_id=$2, input=$3, expected_result=$4, points=$5 WHERE id=$6`

	err := r.db.Execute(ctx, sql, newItem.ID, newItem.TaskID, newItem.Input, newItem.ExpectedResult, newItem.Points, id)

	if err != nil {
		return fmt.Errorf("In TestRepository(UpdateItem): %w", err)
	}

	cacheKey := fmt.Sprintf("tests_taskID_%d", newItem.TaskID)
	r.cache.Delete("tests")
	r.cache.Delete(cacheKey)
	return nil
}

func (r *TestRepository) GetTests(ctx context.Context) ([]model.Test, error) {
	if cachedTests, found := r.cache.Get("tests"); found {
		return cachedTests.([]model.Test), nil
	}

	sql := `SELECT id, task_id, input, expected_result, points FROM tests`

	var tests []dbModel.Test
	err := r.db.Query(ctx, &tests, sql)
	if err != nil {
		return nil, fmt.Errorf("In TestRepository(GetTable): %w", err)
	}

	res := dbModel.DbTestsToModels(tests)
	r.cache.Set("tests", res, cache.DefaultExpiration)
	return res, nil
}

func (r *TestRepository) FindTestByID(ctx context.Context, id int) (model.Test, error) {
	sql := `SELECT id, task_id, input, expected_result, points FROM tests WHERE id = $1`

	var test dbModel.Test
	err := r.db.QueryRow(ctx, &test, sql, id)

	if err != nil {
		return model.Test{}, fmt.Errorf("In TestRepository(FindItemByID): %w", err)
	}
	return model.Test(test), err
}

func (r *TestRepository) FindTestsByTaskID(ctx context.Context, taskID int) ([]model.Test, error) {
	cacheKey := fmt.Sprintf("tests_taskID_%d", taskID)
	if cachedTests, found := r.cache.Get(cacheKey); found {
		return cachedTests.([]model.Test), nil
	}

	sql := `SELECT id, task_id, input, expected_result, points FROM tests WHERE task_id = $1`

	var tests []dbModel.Test
	err := r.db.Query(ctx, &tests, sql, taskID)
	if err != nil {
		return nil, fmt.Errorf("In TestRepository(GetTable): %w", err)
	}

	res := dbModel.DbTestsToModels(tests)
	r.cache.Set(cacheKey, res, cache.DefaultExpiration)
	return res, nil
}
