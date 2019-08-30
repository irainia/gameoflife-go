package file

import (
	"errors"
	"fmt"
	"io/ioutil"
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
	InvalidFormatError    = "format is invalid ('o': true and '-': false)"
	NilGenerationError    = "generation is nil"
)

type FileStream struct {
	path string
}

func (fileStream *FileStream) Read() ([][]bool, error) {
	if _, err := os.Stat(fileStream.path); os.IsNotExist(err) {
		return nil, errors.New(NotFoundFileError)
	}

	readGeneration, _ := ioutil.ReadFile(fileStream.path)
	if string(readGeneration) == "" {
		return nil, errors.New(EmptyFileError)
	}

	splitGeneration := strings.Split(string(readGeneration), "\n")
	outputGeneration := make([][]bool, len(splitGeneration))
	for i := 0; i < len(splitGeneration); i++ {
		outputGeneration[i] = make([]bool, len(splitGeneration[i]))
		for j := 0; j < len(splitGeneration[i]); j++ {
			if splitGeneration[i][j] == 'o' {
				outputGeneration[i][j] = true
			} else if splitGeneration[i][j] == '-' {
				outputGeneration[i][j] = false
			} else {
				return nil, errors.New(InvalidFormatError)
			}
		}
	}

	return outputGeneration, nil
}

func (fileStream *FileStream) Write(generation [][]bool) error {
	return errors.New(NilGenerationError)
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
