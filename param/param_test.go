package param_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Irainia/gameoflife-go/io"
	"github.com/Irainia/gameoflife-go/io/file"
	"github.com/Irainia/gameoflife-go/param"
)

func TestNew(t *testing.T) {
	t.Run("should return nil and error for nil args", func(t *testing.T) {
		var args []string = nil
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NilArgsError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for empty args", func(t *testing.T) {
		var args []string = make([]string, 0)
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.EmptyArgsError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for no separator assignment", func(t *testing.T) {
		var args []string = []string{
			"--inputtype",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NoSeparatorError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for unknown input type value", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=unknown",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.UnknownInputTypeValueError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for unknown argument", func(t *testing.T) {
		var args []string = []string{
			"--unknown=unknown",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.UnknownArgumentError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for input type file with no input path", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NoInputPathError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for no output type", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NoOutputTypeError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for unknown output type value", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=unknown",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.UnknownOutputTypeValueError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for output type file with output path", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NoOutputPathError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for no input type", func(t *testing.T) {
		var args []string = []string{
			"--outputtype=file",
			"--outputpath=./output.cell",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NoInputTypeError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for no generation", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NoGenerationError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for invalid generation", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=invalid",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.InvalidGenerationError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for less than one generation", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=0",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.LessThanOneGenerationError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for input type custom and reader is nilt", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=custom",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=1",
		}
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NoCustomReaderError)

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for output type custom and writer nil", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=custom",
			"--outputtype=custom",
			"--generation=1",
		}
		fileStream, _ := file.New("./input.cell")
		var expectedParam *param.Param = nil
		var expectedError error = errors.New(param.NoCustomWriterError)

		actualParam, actualError := param.New(args, fileStream, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
			return
		}
		if actualError.Error() != expectedError.Error() {
			t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
		}
	})

	t.Run("should return nil and error for reader error", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=1",
		}
		var expectedParam *param.Param = nil

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
		}
	})

	t.Run("should return nil and error for writer error", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output",
			"--generation=1",
		}
		var expectedParam *param.Param = nil

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam != expectedParam {
			t.Error("expected: nil -- actual: not nil")
			return
		}
		if actualError == nil {
			t.Errorf("expected: not nil -- actual: nil")
		}
	})

	t.Run("should return param and nil for valid args", func(t *testing.T) {
		var args []string = []string{
			"--inputtype=file",
			"--inputpath=./input.cell",
			"--outputtype=file",
			"--outputpath=./output.cell",
			"--generation=1",
		}
		var expectedError error = nil

		actualParam, actualError := param.New(args, nil, nil)

		if actualParam == nil {
			t.Error("expected: not nil -- actual: nil")
			return
		}
		if actualError != expectedError {
			t.Errorf("expected nil -- actual: %s", actualError.Error())
		}
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

		if actualNumOfGeneration != expectedNumOfGeneration {
			t.Errorf("expected: %d -- actual: %d", expectedNumOfGeneration, actualNumOfGeneration)
		}
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

		if reflect.TypeOf(actualReader) != reflect.TypeOf(expectedReader) {
			t.Errorf("expected: %d -- actual: %d", expectedReader, actualReader)
		}
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

		if reflect.TypeOf(actualReader) != reflect.TypeOf(expectedReader) {
			t.Errorf("expected: %d -- actual: %d", expectedReader, actualReader)
		}
	})
}
