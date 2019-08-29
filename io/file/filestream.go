package file

import (
	"errors"
	"fmt"
	"strings"
)

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

	splitPath := strings.Split(path, ".")
	if fmt.Sprintf(".%s", splitPath[len(splitPath)-1]) != FileExtension {
		return nil, errors.New(InvalidExtensionError)
	}

	var fileStream = FileStream{}
	return &fileStream, nil
}
