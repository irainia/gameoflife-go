package cell

import "errors"

const (
	ArgumentNilError   = "argument passed is nil"
	ArgumentEmptyError = "argument passed is empty"
)

type CellState struct {
	currentState [][]bool
}

func New(initialState [][]bool) (*CellState, error) {
	if initialState == nil {
		return nil, errors.New(ArgumentNilError)
	}
	return nil, errors.New(ArgumentEmptyError)
}
