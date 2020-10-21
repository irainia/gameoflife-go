package file_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/irainia/gameoflife-go/cell"
	"github.com/irainia/gameoflife-go/io/file"
	"github.com/stretchr/testify/assert"
)

const (
	cellDirectory = "./"
	tubCell       = "tub.cell"
	gliderCell    = "glider.cell"
	emptyCell     = "empty.cell"
	invalidCell   = "invalid.cell"
)

var (
	tubGeneration = [][]bool{
		{false, true, false},
		{true, false, true},
		{false, true, false},
	}
	gliderGeneration = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
)

const (
	emptyGenerationString   string = ""
	invalidGenerationString string = "o--ox"
	gliderGenerationString  string = "-o-\n--o\nooo"
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

	path = fmt.Sprintf("%s%s", cellDirectory, gliderCell)
	err = os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	t.Run("should return nil and error for empty path", func(t *testing.T) {
		var expectedError = file.PathEmptyError

		actualFileStream, actualError := file.New("")

		assert.Nil(t, actualFileStream)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for invalid file extension", func(t *testing.T) {
		var path string = "input"
		var expectedError = file.InvalidExtensionError

		actualFileStream, actualError := file.New(path)

		assert.Nil(t, actualFileStream)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return file stream and nil for valid file extension", func(t *testing.T) {
		var path string = "input.cell"

		actualFileStream, actualError := file.New(path)

		assert.NotNil(t, actualFileStream)
		assert.Nil(t, actualError)
	})
}

func TestRead(t *testing.T) {
	t.Run("should return nil and error for non existent file", func(t *testing.T) {
		var nonExistentFile = "nonexistent.cell"
		fileStream, _ := file.New(nonExistentFile)
		var expectedError = file.NotFoundFileError

		actualGeneration, actualError := fileStream.Read()

		assert.Nil(t, actualGeneration)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for empty file", func(t *testing.T) {
		path := fmt.Sprintf("%s%s", cellDirectory, emptyCell)
		fileStream, _ := file.New(path)
		var expectedError = file.EmptyFileError

		actualGeneration, actualError := fileStream.Read()

		assert.Nil(t, actualGeneration)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for invalid format", func(t *testing.T) {
		path := fmt.Sprintf("%s%s", cellDirectory, invalidCell)
		fileStream, _ := file.New(path)
		var expectedError = file.InvalidFormatError

		actualGeneration, actualError := fileStream.Read()

		assert.Nil(t, actualGeneration)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return generation and nil for valid file", func(t *testing.T) {
		path := fmt.Sprintf("%s%s", cellDirectory, tubCell)
		fileStream, _ := file.New(path)
		expectedGeneration := tubGeneration

		actualGeneration, actualError := fileStream.Read()

		assert.EqualValues(t, expectedGeneration, actualGeneration)
		assert.Nil(t, actualError)
	})
}

func TestWrite(t *testing.T) {
	t.Run("should return error for nil generation", func(t *testing.T) {
		path := fmt.Sprintf("%s%s", cellDirectory, gliderCell)
		fileStream, _ := file.New(path)
		var nilGeneration [][]bool = nil
		var expectedError = file.NilGenerationError

		actualError := fileStream.Write(nilGeneration)

		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return error for empty generation", func(t *testing.T) {
		path := fmt.Sprintf("%s%s", cellDirectory, gliderCell)
		fileStream, _ := file.New(path)
		var emptyGeneration [][]bool = make([][]bool, 0)
		var expectedError = file.EmptyGenerationError

		actualError := fileStream.Write(emptyGeneration)

		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil for valid generation", func(t *testing.T) {
		path := fmt.Sprintf("%s%s", cellDirectory, gliderCell)
		fileStream, _ := file.New(path)
		var validGeneration [][]bool = gliderGeneration
		var expectedGeneration string = gliderGenerationString

		actualError := fileStream.Write(validGeneration)
		actualGeneration, err := ioutil.ReadFile(path)

		assert.Nil(t, actualError)
		assert.Nil(t, err)
		assert.Equal(t, expectedGeneration, string(actualGeneration))
	})
}
