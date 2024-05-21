package linux

import (
	"contest/lib/byteconv"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type LinuxAdapter struct {
}

func NewLinuxAdapter() LinuxAdapter { return LinuxAdapter{} }

func (l LinuxAdapter) CreateTempFileWithText(text string, extension string) (*os.File, error) {
	fileName := time.Now().String() + extension
	file, err := os.Create(fileName)

	if err != nil {
		return nil, fmt.Errorf("")
	}
	if _, err = file.WriteString(text); err != nil {
		return nil, fmt.Errorf("In utils(CreateFileWithText): %w", err)
	}
	return file, nil
}

func (l LinuxAdapter) AddFileExecutablePermission(fileName string) error {
	cmd := exec.Command("chmod", "+x", fileName)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("In utils(AddFileExecutablePermission): %s", byteconv.String(output)+err.Error())
	}
	return nil
}
