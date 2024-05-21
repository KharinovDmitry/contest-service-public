package repository

import (
	"contest/internal/domain/model"
	"context"
)

type TaskRepository interface {
	AddTask(ctx context.Context, title string, text string) error
	DeleteTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, id int, newItem model.Task) error
	GetTasks(ctx context.Context) ([]model.Task, error)
	FindTaskByID(ctx context.Context, id int) (model.Task, error)
}
