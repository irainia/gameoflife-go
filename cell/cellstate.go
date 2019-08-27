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
	if initialState == nil {
		return nil, errors.New(ArgumentNilError)
	}
	if len(initialState) == 0 {
		return nil, errors.New(ArgumentEmptyError)
	}

	colLength := len(initialState[0])
	for i := 0; i < len(initialState); i++ {
		if len(initialState[i]) != colLength {
			return nil, errors.New(ArgumentShapeNotRectangleError)
		}
	}

	cellState := CellState{
		currentState: duplicateState(initialState),
	}
	return &cellState, nil
}

func duplicateState(originalState [][]bool) [][]bool {
	duplicateState := make([][]bool, len(originalState))
	for i := 0; i < len(originalState); i++ {
		duplicateState[i] = make([]bool, len(originalState[i]))
		copy(duplicateState[i], originalState[i])
	}

	return duplicateState
}
