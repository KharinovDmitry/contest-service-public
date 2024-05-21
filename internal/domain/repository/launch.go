package repository

import (
	"contest/internal/domain/model"
	"context"
)

type LaunchRepository interface {
	AddLaunch(ctx context.Context, launch model.Launch) error
	GetLaunchesByUser(ctx context.Context, userID int) ([]model.Launch, error)
	GetSuccessLaunchesByUser(ctx context.Context, userID int) ([]model.Launch, error)
	GetLaunchesByUserAndContest(ctx context.Context, userID int, contestID int) ([]model.Launch, error)
}
