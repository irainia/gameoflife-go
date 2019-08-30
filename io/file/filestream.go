package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	FileExtension = ".cell"
)

const (
	PathEmptyError        = "path passed is empty"
	InvalidExtensionError = "invalid file extension"
	NotFoundFileError     = "file is not found"
	EmptyFileError        = "file is empty"
)

type FileStream struct {
	path string
}

func (fileStream *FileStream) Read() ([][]bool, error) {
	if _, err := os.Stat(fileStream.path); os.IsNotExist(err) {
		return nil, errors.New(NotFoundFileError)
	}

	return nil, errors.New(EmptyFileError)
}

func New(path string) (*FileStream, error) {
	if path == "" {
		return nil, errors.New(PathEmptyError)
	}
	if !isExtensionValid(path) {
		return nil, errors.New(InvalidExtensionError)
	}

	var fileStream = FileStream{
		path: path,
	}
	return &fileStream, nil
}

func isExtensionValid(path string) bool {
	splitPath := strings.Split(path, ".")
	if fmt.Sprintf(".%s", splitPath[len(splitPath)-1]) == FileExtension {
		return true
	}
	return false
}
