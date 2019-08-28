package cell

import (
	"errors"
)

const (
	ArgumentNilError               = "argument passed is nil"
	ArgumentEmptyError             = "argument passed is empty"
	ArgumentShapeNotRectangleError = "argument shape is not rectangle"
)

type CellState struct {
	currentGeneration [][]bool
}

func (cellState *CellState) GetCurrentGeneration() [][]bool {
	return duplicateState(cellState.currentGeneration)
}

func (cellstate *CellState) GetNextGeneration() [][]bool {
	currentGeneration := cellstate.GetCurrentGeneration()
	expandedGeneration := expandState(currentGeneration, 2)

	nextTempCell := make([][]bool, len(expandedGeneration))
	for i := 0; i < len(expandedGeneration); i++ {
		nextTempCell[i] = make([]bool, len(expandedGeneration[i]))
	}

	for i := 1; i < len(expandedGeneration)-1; i++ {
		for j := 1; j < len(expandedGeneration[i])-1; j++ {
			numOfNeighbors := 0
			for p := i - 1; p <= i+1; p++ {
				for q := j - 1; q <= j+1; q++ {
					if p == i && q == j {
						continue
					}

					if expandedGeneration[p][q] {
						numOfNeighbors++
					}
				}
			}

			if numOfNeighbors < 2 {
				nextTempCell[i][j] = false
			}
			if numOfNeighbors == 2 && expandedGeneration[i][j] || numOfNeighbors == 3 {
				nextTempCell[i][j] = true
			}
			if numOfNeighbors > 3 {
				nextTempCell[i][j] = false
			}
		}
	}

	return trimState(nextTempCell)
}

func New(initialState [][]bool) (*CellState, error) {
	isStateValid, err := isStateValid(initialState)
	if !isStateValid || err != nil {
		return nil, err
	}

	cellState := CellState{
		currentGeneration: trimState(initialState),
	}
	return &cellState, nil
}

func isStateValid(state [][]bool) (bool, error) {
	if state == nil {
		return false, errors.New(ArgumentNilError)
	}
	if len(state) == 0 {
		return false, errors.New(ArgumentEmptyError)
	}

	colLength := len(state[0])
	for i := 0; i < len(state); i++ {
		if len(state[i]) != colLength {
			return false, errors.New(ArgumentShapeNotRectangleError)
		}
	}

	return true, nil
}

func duplicateState(originalState [][]bool) [][]bool {
	if !isLivingCellExist(originalState) {
		return make([][]bool, 0)
	}

	duplicateState := make([][]bool, len(originalState))
	for i := 0; i < len(originalState); i++ {
		duplicateState[i] = make([]bool, len(originalState[i]))
		copy(duplicateState[i], originalState[i])
	}

	return duplicateState
}

func isLivingCellExist(state [][]bool) bool {
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			if state[i][j] {
				return true
			}
		}
	}

	return false
}

func trimState(originalState [][]bool) [][]bool {
	if !isLivingCellExist(originalState) {
		return make([][]bool, 0)
	}

	minRowIndex := len(originalState)
	maxRowIndex := 0
	minColIndex := len(originalState[0])
	maxColIndex := 0
	for i := 0; i < len(originalState); i++ {
		for j := 0; j < len(originalState[i]); j++ {
			if originalState[i][j] {
				if i < minRowIndex {
					minRowIndex = i
				}
				if i > maxRowIndex {
					maxRowIndex = i
				}
				if j < minColIndex {
					minColIndex = j
				}
				if j > maxColIndex {
					maxColIndex = j
				}
			}
		}
	}

	trimmedState := make([][]bool, maxRowIndex-minRowIndex+1)
	for i := minRowIndex; i <= maxRowIndex; i++ {
		trimmedState[i-minRowIndex] = make([]bool, maxColIndex-minColIndex+1)
		for j := minColIndex; j <= maxColIndex; j++ {
			trimmedState[i-minRowIndex][j-minColIndex] = originalState[i][j]
		}
	}

	return trimmedState
}

func expandState(originalState [][]bool, additionalEachSide int) [][]bool {
	expandedState := make([][]bool, len(originalState)+additionalEachSide*2)
	for i := 0; i < len(expandedState); i++ {
		expandedState[i] = make([]bool, len(originalState[0])+additionalEachSide*2)
		if i >= additionalEachSide && i < len(expandedState)-additionalEachSide {
			copy(expandedState[i][additionalEachSide:], originalState[i-additionalEachSide])
		}
	}

	return expandedState
}
