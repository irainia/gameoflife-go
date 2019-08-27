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

	isLivingCellExist := isLivingCellExist(initialState)
	if !isLivingCellExist {
		cellState := CellState{
			currentState: duplicateState(initialState),
		}
		return &cellState, nil
	}

	minColIndex := len(initialState[0])
	maxColIndex := 0
	minRowIndex := len(initialState)
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
			if initialState[i][j] && i < minRowIndex {
				minRowIndex = i
			}
		}
	}

	trimmedState := make([][]bool, maxRowIndex-minRowIndex+1)
	for i := minRowIndex; i <= maxRowIndex; i++ {
		trimmedState[i-minRowIndex] = make([]bool, maxColIndex-minColIndex+1)
		for j := minColIndex; j <= maxColIndex; j++ {
			trimmedState[i-minRowIndex][j-minColIndex] = initialState[i][j]
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
