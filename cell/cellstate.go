package cell

import "errors"

const (
	ArgumentNilError = "argument passed is nil"
)

type CellState struct {
	currentState [][]bool
}

func New(initialState [][]bool) (*CellState, error) {
	return nil, errors.New(ArgumentNilError)
}
