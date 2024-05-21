package javaScript

import (
	"context"
	"os"
	"time"
)

type JavaScriptInterpreter struct {
	codeFile *os.File
}

func newJavaScriptInterpreter(codeFile *os.File) *JavaScriptInterpreter {
	return &JavaScriptInterpreter{
		codeFile: codeFile,
	}
}

func (c *JavaScriptInterpreter) Execute(input string, memoryLimitInKb int, timeLimitInMs int) (output string, err error) {
	timeout := time.Duration(timeLimitInMs) * time.Millisecond
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	panic("implement me")
}
