package cell_test

import (
	"errors"
	"testing"

	"github.com/Irainia/gameoflife-go/cell"
)

func TestNewShouldReturnNilAndErrorForInitialGenerationNil(t *testing.T) {
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

func TestNewShouldReturnNilAndErrorForInitialGenerationEmpty(t *testing.T) {
	var expectedCellState *cell.CellState = nil
	var expectedError error = errors.New(cell.ArgumentEmptyError)
	var dim int = 0
	initialGeneration := make([][]bool, dim)

	actualCellState, actualError := cell.New(initialGeneration)

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

func TestNewShouldReturnNilAndErrorForInitialGenerationNotRectangle(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{true, true, true},
		{true},
		{true, true},
	}
	var expectedCellState *cell.CellState = nil
	var expectedError = errors.New(cell.ArgumentShapeNotRectangleError)

	actualGeneration, actualError := cell.New(initialGeneration)

	if actualGeneration != expectedCellState {
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

func TestGetCurrentGenerationShouldReturnInitialGeneration(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	cellState, _ := cell.New(initialGeneration)
	var expectedGeneration [][]bool = initialGeneration

	actualGeneration := cellState.GetCurrentGeneration()

	for i := 0; i < len(expectedGeneration); i++ {
		for j := 0; j < len(expectedGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j])
			}
		}
	}
}

func TestGetCurrentGenerationShouldBeImmutableOnCreation(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	cellState, _ := cell.New(initialGeneration)
	var expectedGeneration [][]bool = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}

	initialGeneration[1][1] = !initialGeneration[1][1]
	actualGeneration := cellState.GetCurrentGeneration()

	for i := 0; i < len(expectedGeneration); i++ {
		for j := 0; j < len(expectedGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j])
			}
		}
	}
}

func TestGetCurrentGenerationShouldBeImmutableOnRetrieval(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	cellState, _ := cell.New(initialGeneration)
	var expectedGeneration [][]bool = [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}

	temporaryState := cellState.GetCurrentGeneration()
	temporaryState[1][1] = !temporaryState[1][1]
	actualGeneration := cellState.GetCurrentGeneration()

	for i := 0; i < len(expectedGeneration); i++ {
		for j := 0; j < len(expectedGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j])
			}
		}
	}
}

func TestGetCurrentGenerationShouldTrimRightSideToNearestLivingCell(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{true, true, false},
		{true, true, false},
	}
	cellState, _ := cell.New(initialGeneration)
	var expectedGeneration [][]bool = [][]bool{
		{true, true},
		{true, true},
	}

	actualGeneration := cellState.GetCurrentGeneration()

	if len(actualGeneration) != len(expectedGeneration) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
		return
	}
	if len(actualGeneration[0]) != len(expectedGeneration[0]) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration[0]), len(actualGeneration[0]))
		return
	}
	for i := 0; i < len(expectedGeneration); i++ {
		for j := 0; j < len(expectedGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j])
			}
		}
	}
}

func TestGetCurrentGenerationShouldTrimBottomSideToNearestLivingCell(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{true, true},
		{true, true},
		{false, false},
	}
	cellState, _ := cell.New(initialGeneration)
	var expectedGeneration [][]bool = [][]bool{
		{true, true},
		{true, true},
	}

	actualGeneration := cellState.GetCurrentGeneration()

	if len(actualGeneration) != len(expectedGeneration) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
		return
	}
	if len(actualGeneration[0]) != len(expectedGeneration[0]) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration[0]), len(actualGeneration[0]))
		return
	}
	for i := 0; i < len(expectedGeneration); i++ {
		for j := 0; j < len(expectedGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j])
			}
		}
	}
}

func TestGetCurrentGenerationShouldTrimLeftSideToNearestLivingCell(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{false, true, true},
		{false, true, true},
	}
	cellState, _ := cell.New(initialGeneration)
	var expectedGeneration [][]bool = [][]bool{
		{true, true},
		{true, true},
	}

	actualGeneration := cellState.GetCurrentGeneration()

	if len(actualGeneration) != len(expectedGeneration) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
		return
	}
	if len(actualGeneration[0]) != len(expectedGeneration[0]) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration[0]), len(actualGeneration[0]))
		return
	}
	for i := 0; i < len(expectedGeneration); i++ {
		for j := 0; j < len(expectedGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j])
			}
		}
	}
}

func TestGetCurrentGenerationShouldTrimTopSideToNearestLivingCell(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{false, false},
		{true, true},
		{true, true},
	}
	cellState, _ := cell.New(initialGeneration)
	var expectedGeneration [][]bool = [][]bool{
		{true, true},
		{true, true},
	}

	actualGeneration := cellState.GetCurrentGeneration()

	if len(actualGeneration) != len(expectedGeneration) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
		return
	}
	if len(actualGeneration[0]) != len(expectedGeneration[0]) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration[0]), len(actualGeneration[0]))
		return
	}
	for i := 0; i < len(expectedGeneration); i++ {
		for j := 0; j < len(expectedGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				t.Errorf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j])
			}
		}
	}
}

func TestGetCurrentGenerationShouldReturnEmptyNoLivingcell(t *testing.T) {
	var initialGeneration [][]bool = [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}
	cellState, _ := cell.New(initialGeneration)
	var expectedGeneration [][]bool = make([][]bool, 0)

	actualGeneration := cellState.GetCurrentGeneration()

	if len(actualGeneration) != len(expectedGeneration) {
		t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
	}
}
