package model

import "contest/internal/domain/enum"

type TestsResult struct {
	ResultCode  enum.TestResultCode
	Description string
	Points      int
}
