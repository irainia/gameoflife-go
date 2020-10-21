package param_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/irainia/gameoflife-go/io"
	"github.com/irainia/gameoflife-go/io/file"
	"github.com/irainia/gameoflife-go/param"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("should return nil and error for nil args", func(t *testing.T) {
		var args []string = nil
		var expectedError = param.NilArgsError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for empty args", func(t *testing.T) {
		var args []string = make([]string, 0)
		var expectedError = param.EmptyArgsError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for no separator assignment", func(t *testing.T) {
		var args []string = []string{
			"--inputtype",
		}
		var expectedError = param.NoSeparatorError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for unknown input type value", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=unknown",
		}
		var expectedError = param.UnknownInputTypeValueError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for unknown argument", func(t *testing.T) {
		var args []string = []string{
			"--unknown=unknown",
		}
		var expectedError = param.UnknownArgumentError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for input type file with no input path", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
		}
		var expectedError = param.NoInputPathError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for no output type", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
		}
		var expectedError = param.NoOutputTypeError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for unknown output type value", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=unknown",
		}
		var expectedError = param.UnknownOutputTypeValueError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for output type file with output path", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
		}
		var expectedError = param.NoOutputPathError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for no input type", func(t *testing.T) {
		var args []string = []string{
			"--outputtype=file",
			"--outputpath=./output.cell",
		}
		var expectedError = param.NoInputTypeError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for no generation", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
		}
		var expectedError = param.NoGenerationError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for invalid generation", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=invalid",
		}
		var expectedError = param.InvalidGenerationError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for less than one generation", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=0",
		}
		var expectedError = param.LessThanOneGenerationError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for input type custom and reader is nilt", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=custom",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=1",
		}
		var expectedError = param.NoCustomReaderError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for output type custom and writer nil", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=custom",
			"--outputtype=custom",
			"--generation=1",
		}
		fileStream, _ := file.New("./input.cell")
		var expectedError = param.NoCustomWriterError

		actualParam, actualError := param.New(args, fileStream, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for reader error", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=1",
		}

		var expectedError = file.InvalidExtensionError

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.EqualError(t, actualError, expectedError)
	})

	t.Run("should return nil and error for writer error", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output",
			"--generation=1",
		}

		actualParam, actualError := param.New(args, nil, nil)

		assert.Nil(t, actualParam)
		assert.NotNil(t, actualError)
	})

	t.Run("should return param and nil for valid args", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=1",
		}

		actualParam, actualError := param.New(args, nil, nil)

		assert.NotNil(t, actualParam)
		assert.Nil(t, actualError)
	})
}

func TestGetNumberOfGeneration(t *testing.T) {
	t.Run("should return the same number as parameter", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=10",
		}
		parameter, _ := param.New(args, nil, nil)
		var expectedNumOfGeneration int = 10

		actualNumOfGeneration := parameter.GetNumOfGeneration()

		assert.Equal(t, expectedNumOfGeneration, actualNumOfGeneration)
	})
}

func TestGetReader(t *testing.T) {
	t.Run("should return the same reader as parameter", func(t *testing.T) {
		var path string = "./input.cell"
		var args []string = []string{
			"--inputtype=file",
			fmt.Sprintf("--inputpath=%s", path),
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=10",
		}
		parameter, _ := param.New(args, nil, nil)
		fileStream, _ := file.New(path)
		var expectedReader io.Reader = fileStream

		actualReader := parameter.GetReader()

		assert.Equal(t, reflect.TypeOf(expectedReader), reflect.TypeOf(actualReader))
	})
}

func TestGetWriter(t *testing.T) {
	t.Run("should return the same writer as parameter", func(t *testing.T) {
		var path string = "./output.cell"
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			fmt.Sprintf("--outputpath=%s", path),
			"--generation=10",
		}
		parameter, _ := param.New(args, nil, nil)
		fileStream, _ := file.New(path)
		var expectedReader io.Writer = fileStream

		actualReader := parameter.GetWriter()

		assert.Equal(t, reflect.TypeOf(expectedReader), reflect.TypeOf(actualReader))
	})
}
