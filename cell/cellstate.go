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

func (cellState *CellState) GetGeneration() [][]bool {
	return duplicateGeneration(cellState.currentGeneration)
}

func (cellState *CellState) GetNextState() *CellState {
	currentGeneration := cellState.GetGeneration()
	expandedCurrentGeneration := expandGeneration(currentGeneration, 2)
	nextGeneration := makeNextGeneration(expandedCurrentGeneration)

	nextState := CellState{
		currentGeneration: nextGeneration,
	}
	return &nextState
}

func (cellState *CellState) String() string {
	return "o"
}

func New(initialGeneration [][]bool) (*CellState, error) {
	isValid, err := isGenerationValid(initialGeneration)
	if !isValid || err != nil {
		return nil, err
	}

	cellState := CellState{
		currentGeneration: trimGeneration(initialGeneration),
	}
	return &cellState, nil
}

func isGenerationValid(generation [][]bool) (bool, error) {
	if generation == nil {
		return false, errors.New(ArgumentNilError)
	}
	if len(generation) == 0 {
		return false, errors.New(ArgumentEmptyError)
	}

	colLength := len(generation[0])
	for i := 0; i < len(generation); i++ {
		if len(generation[i]) != colLength {
			return false, errors.New(ArgumentShapeNotRectangleError)
		}
	}

	return true, nil
}

func duplicateGeneration(originalGeneration [][]bool) [][]bool {
	if !isLivingCellExist(originalGeneration) {
		return make([][]bool, 0)
	}

	duplicatedGeneration := make([][]bool, len(originalGeneration))
	for i := 0; i < len(originalGeneration); i++ {
		duplicatedGeneration[i] = make([]bool, len(originalGeneration[i]))
		copy(duplicatedGeneration[i], originalGeneration[i])
	}

	return duplicatedGeneration
}

func isLivingCellExist(generation [][]bool) bool {
	for i := 0; i < len(generation); i++ {
		for j := 0; j < len(generation[i]); j++ {
			if generation[i][j] {
				return true
			}
		}
	}

	return false
}

func trimGeneration(originalGeneration [][]bool) [][]bool {
	if !isLivingCellExist(originalGeneration) {
		return make([][]bool, 0)
	}

	minRowIndex := len(originalGeneration)
	maxRowIndex := 0
	minColIndex := len(originalGeneration[0])
	maxColIndex := 0
	for i := 0; i < len(originalGeneration); i++ {
		for j := 0; j < len(originalGeneration[i]); j++ {
			if originalGeneration[i][j] {
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

	trimmedGeneration := make([][]bool, maxRowIndex-minRowIndex+1)
	for i := minRowIndex; i <= maxRowIndex; i++ {
		trimmedGeneration[i-minRowIndex] = make([]bool, maxColIndex-minColIndex+1)
		for j := minColIndex; j <= maxColIndex; j++ {
			trimmedGeneration[i-minRowIndex][j-minColIndex] = originalGeneration[i][j]
		}
	}

	return trimmedGeneration
}

func expandGeneration(originalGeneration [][]bool, additionalEachSide int) [][]bool {
	expandedGeneration := make([][]bool, len(originalGeneration)+additionalEachSide*2)
	for i := 0; i < len(expandedGeneration); i++ {
		expandedGeneration[i] = make([]bool, len(originalGeneration[0])+additionalEachSide*2)
		if i >= additionalEachSide && i < len(expandedGeneration)-additionalEachSide {
			copy(expandedGeneration[i][additionalEachSide:], originalGeneration[i-additionalEachSide])
		}
	}

	return expandedGeneration
}

func makeEmptyGeneration(row, column int) [][]bool {
	emptyGeneration := make([][]bool, row)
	for i := 0; i < row; i++ {
		emptyGeneration[i] = make([]bool, column)
	}

	return emptyGeneration
}

func makeNextGeneration(currentGeneration [][]bool) [][]bool {
	row := len(currentGeneration)
	column := len(currentGeneration[0])

	newGeneration := makeEmptyGeneration(row, column)

	for i := 1; i < len(currentGeneration)-1; i++ {
		for j := 1; j < len(currentGeneration[i])-1; j++ {
			numOfNeighbors := 0
			for p := i - 1; p <= i+1; p++ {
				for q := j - 1; q <= j+1; q++ {
					if p == i && q == j {
						continue
					}

					if currentGeneration[p][q] {
						numOfNeighbors++
					}
				}
			}

			if numOfNeighbors == 2 && currentGeneration[i][j] || numOfNeighbors == 3 {
				newGeneration[i][j] = true
			}
		}
	}

	return trimGeneration(newGeneration)
}
