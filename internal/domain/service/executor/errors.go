package executor

import "errors"

var (
	ErrUnknownLanguage = errors.New("unknown language")
	CompileError       = errors.New("compile error")
	MemoryLimitError   = errors.New("memory limit")
	TimeLimitError     = errors.New("time limit")
	RuntimeError       = errors.New("runtime error")
)
