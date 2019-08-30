package file_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Irainia/gameoflife-go/cell"
	"github.com/Irainia/gameoflife-go/io/file"
)

const (
	cellDirectory = "./"
	tubCell       = "tub.cell"
	blinkerCell   = "blinker.cell"
	emptyCell     = "empty.cell"
	invalidCell   = "invalid.cell"
)

var (
	tubGeneration = [][]bool{
		{false, true, false},
		{true, false, true},
		{false, true, false},
	}
)

const (
	emptyGenerationString   string = ""
	invalidGenerationString string = "o--ox"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	cellState, _ := cell.New(tubGeneration)
	path := fmt.Sprintf("%s%s", cellDirectory, tubCell)
	err := ioutil.WriteFile(path, []byte(cellState.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}

	path = fmt.Sprintf("%s%s", cellDirectory, emptyCell)
	err = ioutil.WriteFile(path, []byte(emptyGenerationString), os.ModePerm)
	if err != nil {
		panic(err)
	}

	path = fmt.Sprintf("%s%s", cellDirectory, invalidCell)
	err = ioutil.WriteFile(path, []byte(invalidGenerationString), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	path := fmt.Sprintf("%s%s", cellDirectory, tubCell)
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}

	path = fmt.Sprintf("%s%s", cellDirectory, emptyCell)
	err = os.Remove(path)
	if err != nil {
		panic(err)
	}

	path = fmt.Sprintf("%s%s", cellDirectory, invalidCell)
	err = os.Remove(path)
	if err != nil {
		panic(err)
	}
}

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

func TestNewShouldReturnFileStreamAndNilForValidFileExtension(t *testing.T) {
	var path string = "input.cell"
	var expectedError error = nil

	actualFileStream, actualError := file.New(path)

	if actualFileStream == nil {
		t.Error("expected: not nil -- actual: nil")
		return
	}
	if actualError != expectedError {
		t.Errorf("expected: nil -- actual: %s", actualError)
	}
}

func TestReadShouldReturnNilAndErrorForNonExistentFile(t *testing.T) {
	var nonExistentFile = "nonexistent.cell"
	fileStream, _ := file.New(nonExistentFile)
	var expectedError error = errors.New(file.NotFoundFileError)

	actualGeneration, actualError := fileStream.Read()

	if actualGeneration != nil {
		t.Error("expected: nil -- actual: not nil")
		return
	}
	if actualError == nil {
		t.Error("expected: not nil -- actual: nil")
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}

func TestReadShouldReturnNilAndErrorForEmptyFile(t *testing.T) {
	path := fmt.Sprintf("%s%s", cellDirectory, emptyCell)
	fileStream, _ := file.New(path)
	var expectedError error = errors.New(file.EmptyFileError)

	actualGeneration, actualError := fileStream.Read()

	if actualGeneration != nil {
		t.Error("expected: nil -- actual: not nil")
		return
	}
	if actualError == nil {
		t.Error("expected: not nil -- actual: nil")
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}

func TestReadShouldReturnNilAndErrorForInvalidFormat(t *testing.T) {
	path := fmt.Sprintf("%s%s", cellDirectory, invalidCell)
	fileStream, _ := file.New(path)
	var expectedError error = errors.New(file.InvalidFormatError)

	actualGeneration, actualError := fileStream.Read()

	if actualGeneration != nil {
		t.Error("expected: nil -- actual: not nil")
		return
	}
	if actualError == nil {
		t.Error("expected: not nil -- actual: nil")
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}

func TestReadShouldReturnGenerationAndNilForValidFile(t *testing.T) {
	fileStream, _ := file.New(tubCell)
	expectedGeneration := tubGeneration
	var expectedError error = nil

	actualGeneration, actualError := fileStream.Read()

	if actualGeneration == nil {
		t.Error("expected: not nil -- actual: nil")
		return
	}
	if actualError != expectedError {
		t.Errorf("expected: nil -- actual: %s", actualError.Error())
		return
	}
	for i := 0; i < len(actualGeneration); i++ {
		for j := 0; j < len(actualGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j])
			}
		}
	}
}
