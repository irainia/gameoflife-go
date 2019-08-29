package file

import "errors"

const (
	PathEmptyError = "path passed is empty"
)

type FileStream struct {
}

func New(path string) (*FileStream, error) {
	return nil, errors.New(PathEmptyError)
}
