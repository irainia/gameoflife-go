package cell_test

import (
	"errors"
	"testing"

	"github.com/Irainia/gameoflife-go/cell"
)

func TestNewShouldReturnNilAndErrorForInitialStateNil(t *testing.T) {
	var expectedCellState *cell.CellState = nil
	var expectedError error = errors.New(cell.ArgumentNilError)

	actualCellState, actualError := cell.New(nil)

	if actualCellState != expectedCellState {
		t.Error("expected: nil -- actual: not nil")
		return
	}
	if actualError == nil {
		t.Errorf("expected: %s -- actual: nil", expectedError.Error())
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}

func TestNewShouldReturnNilAndErrorForInitialStateEmpty(t *testing.T) {
	var expectedCellState *cell.CellState = nil
	var expectedError error = errors.New(cell.ArgumentEmptyError)
	var dim int = 0
	initialState := make([][]bool, dim)

	actualCellState, actualError := cell.New(initialState)

	if actualCellState != expectedCellState {
		t.Error("expected: nil -- actual: not nil")
		return
	}
	if actualError == nil {
		t.Errorf("expected: %s -- actual: nil", expectedError.Error())
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}

func TestNewShouldReturnNilAndErrorForInitialStateNotRectangle(t *testing.T) {
	var initialState [][]bool = [][]bool{
		{true, true, true},
		{true},
		{true, true},
	}
	var expectedCellState *cell.CellState = nil
	var expectedError = errors.New(cell.ArgumentShapeNotRectangleError)

	actualState, actualError := cell.New(initialState)

	if actualState != expectedCellState {
		t.Error("expected: nil -- actual: not nil")
		return
	}
	if actualError == nil {
		t.Errorf("expected: %s -- actual: nil", expectedError.Error())
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}

func TestGetCurrentStateShouldReturnInitialState(t *testing.T) {
	var initialState [][]bool = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	cellState, _ := cell.New(initialState)
	var expectedState [][]bool = initialState

	actualState := cellState.GetCurrentState()

	for i := 0; i < len(expectedState); i++ {
		for j := 0; j < len(expectedState[i]); j++ {
			if actualState[i][j] != expectedState[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedState[i][j], i, j, actualState[i][j])
			}
		}
	}
}

func TestGetCurrentStateShouldBeImmutableOnCreation(t *testing.T) {
	var initialState [][]bool = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	cellState, _ := cell.New(initialState)
	var expectedState [][]bool = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}

	initialState[1][1] = !initialState[1][1]
	actualState := cellState.GetCurrentState()

	for i := 0; i < len(expectedState); i++ {
		for j := 0; j < len(expectedState[i]); j++ {
			if actualState[i][j] != expectedState[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedState[i][j], i, j, actualState[i][j])
			}
		}
	}
}
