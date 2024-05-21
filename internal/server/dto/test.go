package dto

import "contest/internal/domain/model"

type TestDTO struct {
	ID             int    `json:"id"`
	TaskID         int    `json:"taskID"`
	Input          string `json:"input"`
	ExpectedResult string `json:"expectedResult"`
	Points         int    `json:"points"`
}

func TestsToTestsDTO(tests []model.Test) []TestDTO {
	res := make([]TestDTO, len(tests))
	for i, test := range tests {
		res[i] = TestDTO(test)
	}
	return res
}
