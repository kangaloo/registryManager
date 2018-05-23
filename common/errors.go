package common

import "fmt"

type ExistError struct {
	file string
}

func (e *ExistError) Error() string {
	return fmt.Sprintf("destnation file %s already exist", e.file)
}

type NotExistError struct {
	file string
}

func (e *NotExistError) Error() string  {
	return fmt.Sprintf("source file %s not exits", e.file)
}

func NewExistErr(s string) *ExistError {
	return &ExistError{file:s}
}