package cell_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Irainia/gameoflife-go/cell"
)

func TestNew(t *testing.T) {
	t.Run("should return nil and error for initial generation nil", func(t *testing.T) {
		var expectedCellState *cell.CellState = nil
		var expectedError error = errors.New(cell.GenerationNilError)

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
	})

	t.Run("should return nil and error for initial generation empty", func(t *testing.T) {
		var expectedCellState *cell.CellState = nil
		var expectedError error = errors.New(cell.GenerationEmptyError)
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
	})

	t.Run("should return nil and error for initial generation not rectangle", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{true, true, true},
			{true},
			{true, true},
		}
		var expectedCellState *cell.CellState = nil
		var expectedError = errors.New(cell.GenerationShapeNotRectangleError)

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
	})
}

func TestGetCurrentGeneration(t *testing.T) {
	t.Run("should return initial generation", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{false, true, false},
			{false, false, true},
			{true, true, true},
		}
		cellState, _ := cell.New(initialGeneration)
		var expectedGeneration [][]bool = initialGeneration

		actualGeneration := cellState.GetGeneration()

		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("should be immutable on creation", func(t *testing.T) {
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
		actualGeneration := cellState.GetGeneration()

		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("should be immutable on retrieval", func(t *testing.T) {
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

		temporaryState := cellState.GetGeneration()
		temporaryState[1][1] = !temporaryState[1][1]
		actualGeneration := cellState.GetGeneration()

		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("should trim right side to nearest living cell", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{true, true, false},
			{true, true, false},
		}
		cellState, _ := cell.New(initialGeneration)
		var expectedGeneration [][]bool = [][]bool{
			{true, true},
			{true, true},
		}

		actualGeneration := cellState.GetGeneration()

		if len(actualGeneration) != len(expectedGeneration) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
			return
		}
		if len(actualGeneration[0]) != len(expectedGeneration[0]) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration[0]), len(actualGeneration[0]))
			return
		}
		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("should trim bottom side to nearest living cell", func(t *testing.T) {
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

		actualGeneration := cellState.GetGeneration()

		if len(actualGeneration) != len(expectedGeneration) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
			return
		}
		if len(actualGeneration[0]) != len(expectedGeneration[0]) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration[0]), len(actualGeneration[0]))
			return
		}
		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("should trim left side to nearest living cell", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{false, true, true},
			{false, true, true},
		}
		cellState, _ := cell.New(initialGeneration)
		var expectedGeneration [][]bool = [][]bool{
			{true, true},
			{true, true},
		}

		actualGeneration := cellState.GetGeneration()

		if len(actualGeneration) != len(expectedGeneration) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
			return
		}
		if len(actualGeneration[0]) != len(expectedGeneration[0]) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration[0]), len(actualGeneration[0]))
			return
		}
		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("should trim top side to nearest living cell", func(t *testing.T) {
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

		actualGeneration := cellState.GetGeneration()

		if len(actualGeneration) != len(expectedGeneration) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
			return
		}
		if len(actualGeneration[0]) != len(expectedGeneration[0]) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration[0]), len(actualGeneration[0]))
			return
		}
		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("should return empty no living cell", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}
		cellState, _ := cell.New(initialGeneration)
		var expectedGeneration [][]bool = make([][]bool, 0)

		actualGeneration := cellState.GetGeneration()

		if len(actualGeneration) != len(expectedGeneration) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
		}
	})

	t.Run("should kill living cell with less than two neighbors", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{true, true},
		}
		cellState, _ := cell.New(initialGeneration)
		var expectedGeneration [][]bool = make([][]bool, 0)

		actualGeneration := cellState.GetNextState().GetGeneration()

		if len(actualGeneration) != len(expectedGeneration) {
			t.Errorf("expected: %d -- actual: %d", len(expectedGeneration), len(actualGeneration))
		}
	})

	t.Run("should survive living cell with two neighbors", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{false, true, false},
			{true, false, true},
			{false, true, false},
		}
		cellState, _ := cell.New(initialGeneration)
		var expectedGeneration [][]bool = [][]bool{
			{false, true, false},
			{true, false, true},
			{false, true, false},
		}

		actualGeneration := cellState.GetNextState().GetGeneration()

		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("should survive or reproduce living cell with three neighbors", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{true, true, true},
		}
		cellState, _ := cell.New(initialGeneration)
		var expectedGeneration [][]bool = [][]bool{
			{true},
			{true},
			{true},
		}

		actualGeneration := cellState.GetNextState().GetGeneration()

		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("shoud kill living cell with more than three neighbors", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{false, false, true},
			{false, true, false},
			{true, false, true},
		}
		cellState, _ := cell.New(initialGeneration)
		var expectedGeneration [][]bool = [][]bool{
			{true, true},
			{true, false},
		}

		actualGeneration := cellState.GetNextState().GetGeneration()

		err := elementWiseCheck(expectedGeneration, actualGeneration)
		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestString(t *testing.T) {
	t.Run("convert True -> 0", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{true},
		}
		cellState, _ := cell.New(initialGeneration)
		expectedString := "o"

		actualString := cellState.String()

		if actualString != expectedString {
			t.Errorf("expected: %s -- actual: %s", expectedString, actualString)
		}
	})

	t.Run("convert False -> '-'", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{true, false, true},
		}
		cellState, _ := cell.New(initialGeneration)
		expectedString := "o-o"

		actualString := cellState.String()

		if actualString != expectedString {
			t.Errorf("expected: %s -- actual: %s", expectedString, actualString)
		}
	})

	t.Run("add new line each row except last row", func(t *testing.T) {
		var initialGeneration [][]bool = [][]bool{
			{true, true},
			{true, true},
		}
		cellState, _ := cell.New(initialGeneration)
		expectedString := "oo\noo"

		actualString := cellState.String()

		if actualString != expectedString {
			t.Errorf("expected: %s -- actual: %s", expectedString, actualString)
		}
	})
}

func elementWiseCheck(expectedGeneration, actualGeneration [][]bool) error {
	for i := 0; i < len(actualGeneration); i++ {
		for j := 0; j < len(actualGeneration[i]); j++ {
			if actualGeneration[i][j] != expectedGeneration[i][j] {
				return errors.New(fmt.Sprintf("expected: [%d][%d] [%t] -- actual: [%d][%d] [%t]",
					i, j, expectedGeneration[i][j], i, j, actualGeneration[i][j]))
			}
		}
	}

	return nil
}
