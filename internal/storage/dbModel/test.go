package dbModel

import "contest/internal/domain/model"

type Test struct {
	ID             int    `db:"id"`
	TaskID         int    `db:"task_id"`
	Input          string `db:"input"`
	ExpectedResult string `db:"expected_result"`
	Points         int    `db:"points"`
}

func DbTestsToModels(tests []Test) []model.Test {
	res := make([]model.Test, len(tests))
	for i, test := range tests {
		res[i] = model.Test(test)
	}
	return res
}
