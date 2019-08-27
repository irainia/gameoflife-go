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
	currentState [][]bool
}

func (cellState *CellState) GetCurrentState() [][]bool {
	return duplicateState(cellState.currentState)
}

func New(initialState [][]bool) (*CellState, error) {
	isValid, err := isStateValid(initialState)
	if !isValid || err != nil {
		return nil, err
	}

	minColIndex := len(initialState[0])
	maxColIndex := 0
	maxRowIndex := 0
	for i := 0; i < len(initialState); i++ {
		for j := 0; j < len(initialState[i]); j++ {
			if initialState[i][j] && j > maxColIndex {
				maxColIndex = j
			}
			if initialState[i][j] && i > maxRowIndex {
				maxRowIndex = i
			}
			if initialState[i][j] && j < minColIndex {
				minColIndex = j
			}
		}
	}

	trimmedState := make([][]bool, maxRowIndex+1)
	for i := 0; i <= maxRowIndex; i++ {
		trimmedState[i] = make([]bool, maxColIndex-minColIndex+1)
		for j := minColIndex; j <= maxColIndex; j++ {
			trimmedState[i][j-minColIndex] = initialState[i][j]
		}
	}

	cellState := CellState{
		currentState: trimmedState,
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
	duplicateState := make([][]bool, len(originalState))
	for i := 0; i < len(originalState); i++ {
		duplicateState[i] = make([]bool, len(originalState[i]))
		copy(duplicateState[i], originalState[i])
	}

	return duplicateState
}
