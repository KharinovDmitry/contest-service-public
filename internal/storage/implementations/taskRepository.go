package implementations

import (
	"contest/internal/domain/model"
	"contest/internal/storage/dbModel"
	"contest/lib/adapter/db"
	"context"
	"fmt"
)

type TaskRepository struct {
	db db.DBAdapter
}

func NewTaskRepository(db db.DBAdapter) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (c *TaskRepository) AddTask(ctx context.Context, title string, text string) error {
	sql := `INSERT INTO tasks(title, text) VALUES ($1, $2)`

	err := c.db.Execute(ctx, sql, title, text)

	if err != nil {
		return fmt.Errorf("In ContestRepository(AddTask): %w", err)
	}
	return nil
}

func (c *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	sql := `DELETE FROM tasks WHERE id = $1`

	err := c.db.Execute(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("In ContestRepository(DeleteTask): %w", err)
	}
	return nil
}

func (c *TaskRepository) UpdateTask(ctx context.Context, id int, newItem model.Task) error {
	sql := `UPDATE tasks SET title=$1, text=$2 WHERE id = $3`

	err := c.db.Execute(ctx, sql, newItem.Title, newItem.Text, id)

	if err != nil {
		return fmt.Errorf("In ContestRepository(UpdateItem): %w", err)
	}
	return nil
}

func (c *TaskRepository) GetTasks(ctx context.Context) ([]model.Task, error) {
	sql := `SELECT id, title, text, memory_limit, time_limit  FROM tasks`

	var tasks []dbModel.Task
	err := c.db.Query(ctx, &tasks, sql)

	if err != nil {
		return nil, fmt.Errorf("In ContestRepository(GetTable): %w", err)
	}
	return dbModel.DbTasksToModels(tasks), nil
}

func (c *TaskRepository) FindTaskByID(ctx context.Context, id int) (model.Task, error) {
	query := `SELECT id, title, text, memory_limit, time_limit FROM tasks WHERE id = $1`

	var task dbModel.Task
	err := c.db.QueryRow(ctx, &task, query, id)

	if err != nil {
		return model.Task{}, fmt.Errorf("In ContestRepository(FindItemByID): %w", err)
	}
	return model.Task(task), err
}
