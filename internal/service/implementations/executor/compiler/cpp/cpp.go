package cpp

import (
	"contest/internal/domain/service/executor"
	os2 "contest/lib/adapter/os"
	"contest/lib/byteconv"
	"fmt"
	"os"
	"os/exec"
)

type CPPCompiler struct {
	osAdapter os2.OSAdapter
}

func NewCPPCompiler(adapter os2.OSAdapter) CPPCompiler {
	return CPPCompiler{
		osAdapter: adapter,
	}
}

func (c CPPCompiler) Compile(code string) (fileName string, err error) {
	file, err := c.osAdapter.CreateTempFileWithText(code, ".cpp")
	defer os.Remove(file.Name())
	defer file.Close()

	if err != nil {
		return "", fmt.Errorf("In CPPCompiler(Compile): %w", err)
	}

	executableName := file.Name() + ".exe"
	cmd := exec.Command("g++", "-o", executableName, "-x", "c++", file.Name())
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("%w: %s", executor.CompileError, byteconv.String(output))
	}

	if err := c.osAdapter.AddFileExecutablePermission(executableName); err != nil {
		return "", fmt.Errorf("In CPPCompiler(Compile): %w", err)
	}
	return executableName, nil
}
