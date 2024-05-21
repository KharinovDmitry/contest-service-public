package dto

import "contest/internal/domain/model"

type TaskDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	MemoryLimit int    `json:"memory_limit"`
	TimeLimit   int    `json:"time_limit"`
}

func TasksToTasksDTO(tasks []model.Task) []TaskDTO {
	res := make([]TaskDTO, len(tasks))
	for i, task := range tasks {
		res[i] = TaskDTO(task)
	}
	return res
}
