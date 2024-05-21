package executor

type Executor interface {
	Execute(input string, memoryLimitInKb int, timeLimitInMs int) (output string, err error)
	Close() error
}
