package file_test

import (
	"errors"
	"testing"

	"github.com/Irainia/gameoflife-go/io/file"
)

func TestNewShouldReturnNilAndErrorForEmptyPath(t *testing.T) {
	var expectedFileStream *file.FileStream = nil
	var expectedError error = errors.New(file.PathEmptyError)

	actualFileStream, actualError := file.New("")

	if actualFileStream != expectedFileStream {
		t.Errorf("expected: nil -- actual: %s", actualFileStream)
		return
	}
	if actualError == nil {
		t.Errorf("expected: not nil -- actual: nil")
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}

func TestNewShouldReturnNilAndErrorForInvalidFileExtension(t *testing.T) {
	var path string = "input"
	var expectedFileStream *file.FileStream = nil
	var expectedError error = errors.New(file.InvalidExtensionError)

	actualFileStream, actualError := file.New(path)

	if actualFileStream != expectedFileStream {
		t.Errorf("expected: nil -- actual: %s", actualFileStream)
		return
	}
	if actualError == nil {
		t.Errorf("expected: not nil -- actual: nil")
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}
