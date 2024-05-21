package dbModel

import "contest/internal/domain/model"

type Task struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Text        string `db:"text"`
	MemoryLimit int    `db:"memory_limit"`
	TimeLimit   int    `db:"time_limit"`
}

func DbTasksToModels(tasks []Task) []model.Task {
	res := make([]model.Task, len(tasks))
	for i, task := range tasks {
		res[i] = model.Task(task)
	}
	return res
}
