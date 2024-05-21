package executor

import (
	"contest/internal/domain/enum"
)

type ExecutorFactory interface {
	NewExecutor(code string, language enum.Language) (Executor, error)
}
