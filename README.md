# Game of Life

## Description

The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970.

The game is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input. One interacts with the Game of Life by creating an initial configuration and observing how it evolves, or, for advanced players, by creating patterns with particular properties.

## Rule

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, alive or dead, (or populated and unpopulated, respectively). Every cell interacts with its eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time, the following transitions occur:

1) Any live cell with fewer than two live neighbours dies, as if by underpopulation.
2) Any live cell with two or three live neighbours lives on to the next generation.
3) Any live cell with more than three live neighbours dies, as if by overpopulation.
4) Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

The initial pattern constitutes the seed of the system. The first generation is created by applying the above rules simultaneously to every cell in the seed; births and deaths occur simultaneously, and the discrete moment at which this happens is sometimes called a tick. Each generation is a pure function of the preceding one. The rules continue to be applied repeatedly to create further generations.

## Dependency

### Go Programming Language

The Go programming language version `go1.12.9` should be installed. Go to [this link](https://golang.org/doc/install) and follow the instructions to install based on the system. To check the installation, we can check its version by running the following command on the terminal:
```
go version
```

Example of the output:
```
go version go1.12.9 darwin/amd64
```

### Go Dep

The Dep, for dependency management tool for Go, version `v0.5.4` should be installed. Go to [this link](https://github.com/golang/dep) and follow the instruction to install based on the system. To pcheck the installation, we can check its version by running the following command on the terminal:
```
dep version
```

Example of the output:
```
dep:
 version     : v0.5.4
 build date  : 2019-06-14
 git hash    : 1f7c19e
 go version  : go1.12.6
 go compiler : gc
 platform    : darwin/amd64
 features    : ImportDuringSolve=false
```

## How to Test

In order to test, go to this project root directory and run the following command:
```
make test
```
