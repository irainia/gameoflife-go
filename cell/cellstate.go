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

	var currentState [][]bool
	if isLivingCellExist(initialState) {
		currentState = trimState(initialState)
	} else {
		currentState = duplicateState(initialState)
	}

	cellState := CellState{
		currentState: currentState,
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
