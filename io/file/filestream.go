package file

import "errors"

const (
	ArgumentEmptyError = "argument passed is empty"
)

type FileStream struct {
}

func New(path string) (*FileStream, error) {
	return nil, errors.New(ArgumentEmptyError)
}
