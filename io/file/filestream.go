package file

import "errors"

const (
	FileExtension = ".cell"
)

const (
	PathEmptyError        = "path passed is empty"
	InvalidExtensionError = "invalid file extension"
)

type FileStream struct {
}

func New(path string) (*FileStream, error) {
	if path == "" {
		return nil, errors.New(PathEmptyError)
	}
	return nil, errors.New(InvalidExtensionError)
}
