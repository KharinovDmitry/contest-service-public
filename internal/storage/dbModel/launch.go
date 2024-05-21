package dbModel

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/model"
	"time"
)

type Launch struct {
	ID          int                 `db:"id"`
	UserID      int                 `db:"user_id"`
	ContestID   int                 `db:"contest_id"`
	Date        time.Time           `db:"date"`
	Code        string              `db:"code"`
	ResultCode  enum.TestResultCode `db:"result_code"`
	Description string              `db:"description"`
	Points      int                 `db:"points"`
}

func DbLaunchesToModels(launches []Launch) []model.Launch {
	res := make([]model.Launch, len(launches))
	for i, launch := range launches {
		res[i] = model.Launch{
			ID:        launch.ID,
			UserID:    launch.UserID,
			ContestID: launch.ContestID,
			Date:      launch.Date,
			Code:      launch.Code,
			TestsResult: model.TestsResult{
				ResultCode:  launch.ResultCode,
				Description: launch.Description,
				Points:      launch.Points,
			},
		}
	}
	return res
}
