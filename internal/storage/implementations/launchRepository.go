package implementations

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/model"
	"contest/internal/storage/dbModel"
	"contest/lib/adapter/db"
	"context"
	"fmt"
)

type LaunchRepository struct {
	db db.DBAdapter
}

func NewLaunchRepository(db db.DBAdapter) *LaunchRepository {
	return &LaunchRepository{
		db: db,
	}
}

func (l *LaunchRepository) AddLaunch(ctx context.Context, launch model.Launch) error {
	sql := `INSERT INTO launches(user_id, task_id, date, code, test_result_code, description, points) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	err := l.db.Execute(ctx, sql, launch.UserID, launch.ContestID, launch.Date, launch.Code, launch.ResultCode, launch.Description, launch.Points)

	if err != nil {
		return fmt.Errorf("In LaunchRepository(AddLaunch): %w", err)
	}
	return nil
}

func (l *LaunchRepository) GetLaunchesByUser(ctx context.Context, userID int) ([]model.Launch, error) {
	sql := `SELECT id, user_id, task_id, date, code, test_result_code, description, points FROM launches WHERE user_id = $1`

	var launches []dbModel.Launch
	err := l.db.Query(ctx, &launches, sql, userID)

	if err != nil {
		return nil, fmt.Errorf("In LaunchRepository(GetLaunchesByUser): %w", err)
	}
	return dbModel.DbLaunchesToModels(launches), nil
}

func (l *LaunchRepository) GetSuccessLaunchesByUser(ctx context.Context, userID int) ([]model.Launch, error) {
	sql := `SELECT id, user_id, task_id, date, code, test_result_code, description, points FROM launches WHERE user_id = $1 AND test_result_code = $2`

	var launches []dbModel.Launch
	err := l.db.Query(ctx, &launches, sql, userID, enum.SuccessCode)

	if err != nil {
		return nil, fmt.Errorf("In LaunchRepository(GetLaunchesByUser): %w", err)
	}
	return dbModel.DbLaunchesToModels(launches), nil
}

func (l *LaunchRepository) GetLaunchesByUserAndContest(ctx context.Context, userID int, contestID int) ([]model.Launch, error) {
	sql := `SELECT id, user_id, task_id, date, code, test_result_code, description, points FROM launches WHERE user_id = $1 AND task_id = $2`

	var launches []dbModel.Launch
	err := l.db.Query(ctx, &launches, sql, userID, contestID)

	if err != nil {
		return nil, fmt.Errorf("In LaunchRepository(GetLaunchesByUserAndContest): %w", err)
	}
	return dbModel.DbLaunchesToModels(launches), nil
}
