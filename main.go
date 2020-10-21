package main

import (
	"fmt"
	"log"
	"os"

	"github.com/irainia/gameoflife-go/cell"
	"github.com/irainia/gameoflife-go/param"
)

func main() {
	args := os.Args

	parameter, err := param.New(args[1:], nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	reader := parameter.GetReader()
	initialGeneration, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	cellState, err := cell.New(initialGeneration)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i <= parameter.GetNumOfGeneration(); i++ {
		fmt.Println()
		fmt.Printf("genereation %d\n", i)
		fmt.Println(cellState)
		cellState = cellState.GetNextState()
	}

	writer := parameter.GetWriter()
	err = writer.Write(cellState.GetGeneration())
	if err != nil {
		log.Fatalln(err)
	}
}
