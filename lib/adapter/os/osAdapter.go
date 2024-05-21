package os

import "os"

type OSAdapter interface {
	CreateTempFileWithText(text string, extension string) (*os.File, error)
	AddFileExecutablePermission(fileName string) error
}
