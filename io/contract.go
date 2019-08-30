package io

type (
	Reader interface {
		Read() ([][]bool, error)
	}

	Writer interface {
		Write(generation [][]bool) error
	}
)
