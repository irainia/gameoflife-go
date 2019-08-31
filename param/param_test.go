package param_test

import (
	"errors"
	"testing"

	"github.com/Irainia/gameoflife-go/param"
)

func TestNewShouldReturnNilAndErrorForNilArgs(t *testing.T) {
	var args []string = nil
	var expectedParam *param.Param = nil
	var expectedError error = errors.New(param.NilArgsError)

	actualParam, actualError := param.New(args)

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

	actualParam, actualError := param.New(args)

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

func TestNewShouldReturnNilAndErrorForNoInputStreamType(t *testing.T) {
	var args []string = []string{
		"",
	}
	var expectedParam *param.Param = nil
	var expectedError error = errors.New(param.NoInputStreamTypeError)

	actualParam, actualError := param.New(args)

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
		"--streamtype",
	}
	var expectedParam *param.Param = nil
	var expectedError error = errors.New(param.NoSeparatorError)

	actualParam, actualError := param.New(args)

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

func TestNewShouldReturnNilAndErrorForNoInputStreamValue(t *testing.T) {
	var args []string = []string{
		"--streamtype=",
	}
	var expectedParam *param.Param = nil
	var expectedError error = errors.New(param.NoInputStreamValueError)

	actualParam, actualError := param.New(args)

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

func TestNewShouldReturnNilAndErrorForUnknownInputStreamValue(t *testing.T) {
	var args []string = []string{
		"--streamtype=unknown",
	}
	var expectedParam *param.Param = nil
	var expectedError error = errors.New(param.UnknownInputStreamValueError)

	actualParam, actualError := param.New(args)

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
