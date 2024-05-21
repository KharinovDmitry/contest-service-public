package linux

import (
	"contest/internal/domain/enum"
	executor2 "contest/internal/domain/service/executor"
	"contest/internal/service/implementations/executor/compiler"
	"contest/internal/service/implementations/executor/compiler/cpp"
	"contest/internal/service/implementations/executor/interpreters/python"
	os2 "contest/lib/adapter/os"

	"errors"
	"fmt"
	"os"
)

type ExecutorFactory struct {
	osAdapter os2.OSAdapter

	languageCompilerMap    map[enum.Language]compiler.Compiler
	languageInterpreterMap map[enum.Language]func(codeFile *os.File) executor2.Executor
}

func NewExecutorFactory(adapter os2.OSAdapter) *ExecutorFactory {
	return &ExecutorFactory{
		osAdapter: adapter,

		languageCompilerMap: map[enum.Language]compiler.Compiler{
			enum.CPP: cpp.NewCPPCompiler(adapter),
		},
		languageInterpreterMap: map[enum.Language]func(codeFile *os.File) executor2.Executor{
			enum.Python: python.NewPythonInterpreter,
		},
	}
}

func (c *ExecutorFactory) NewExecutor(code string, language enum.Language) (executor2.Executor, error) {
	if compiler, exist := c.languageCompilerMap[language]; exist {
		fileName, err := compiler.Compile(code)
		if err != nil {
			if errors.Is(err, executor2.CompileError) {
				return nil, err
			}
			return nil, fmt.Errorf("In ExecutorFactory(NewCodeExecutor): %w", err)
		}

		file, err := os.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("In ExecutorFactory(NewCodeExecutor): %w", err)
		}

		return &ExecutorCompiled{executableFile: file}, nil
	}

	if constructor, exist := c.languageInterpreterMap[language]; exist {
		codeFile, err := c.osAdapter.CreateTempFileWithText(code, ".code")
		if err != nil {
			return nil, fmt.Errorf("In ExecutorFactory(NewCodeExecutor): %w", err)
		}

		return constructor(codeFile), nil
	}

	return nil, executor2.ErrUnknownLanguage
}
