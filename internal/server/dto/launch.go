package dto

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/model"
	"time"
)

type LaunchDTO struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ContestID int       `json:"contest_id"`
	Date      time.Time `json:"date"`
	Code      string    `json:"code"`
	// @Model domain.TestResultCode
	ResultCode  enum.TestResultCode `json:"result_code"`
	Description string              `json:"description"`
	Points      int                 `json:"points"`
}

func LaunchToLaunchDTO(launch model.Launch) LaunchDTO {
	return LaunchDTO{
		ID:          launch.ID,
		UserID:      launch.UserID,
		ContestID:   launch.ContestID,
		Date:        launch.Date,
		Code:        launch.Code,
		ResultCode:  launch.ResultCode,
		Description: launch.Description,
		Points:      launch.Points,
	}
}

func LaunchesToLaunchesDTO(launches []model.Launch) []LaunchDTO {
	res := make([]LaunchDTO, len(launches))
	for i, launch := range launches {
		res[i] = LaunchToLaunchDTO(launch)
	}
	return res
}
