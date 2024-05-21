package compiler

type Compiler interface {
	Compile(code string) (fileName string, err error)
}
