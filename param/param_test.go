package param_test

import (
	"errors"
	"testing"

	"github.com/Irainia/gameoflife-go/io/file"
	"github.com/Irainia/gameoflife-go/param"
)

func TestNewShouldReturnNilAndErrorForNilArgs(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForEmptyArgs(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForNoSeparatorAssignment(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForUnknownInputTypeValue(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForUnknownArgument(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForInputTypeFileNoInputPath(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForNoOutputType(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForUnknownOutputTypeValue(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForOutputTypeFileNoOutputPath(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForNoInputType(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForNoGeneration(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForInvalidGeneration(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForLessThanOneGeneration(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForInputTypeCustomAndReaderNil(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForOutputTypeCustomAndWriterNil(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForReaderError(t *testing.T) {
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
}

func TestNewShouldReturnNilAndErrorForWriterError(t *testing.T) {
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
}

func TestNewShouldReturnParamAndNilForValidArgs(t *testing.T) {
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
}
