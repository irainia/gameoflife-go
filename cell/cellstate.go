package cell

import "errors"

const (
	ArgumentNilError               = "argument passed is nil"
	ArgumentEmptyError             = "argument passed is empty"
	ArgumentShapeNotRectangleError = "argument shape is not rectangle"
)

type CellState struct {
	currentState [][]bool
}

func New(initialState [][]bool) (*CellState, error) {
	if initialState == nil {
		return nil, errors.New(ArgumentNilError)
	}
	if len(initialState) == 0 {
		return nil, errors.New(ArgumentEmptyError)
	}
	return nil, errors.New(ArgumentShapeNotRectangleError)
}
