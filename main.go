package main

import (
	"fmt"
	"log"
)

func readInt() int {
	var value int
	_, err := fmt.Scanf("%d", &value)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return value
}

func readCoordinates() (int, int) {
	var x, y int
	_, err := fmt.Scanf("%d,%d", &y, &x)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return x, y
}

func main() {
	m := readInt()
	n := readInt()
	numberOfGenerations := readInt()
	numberOfLiveCells := readInt()
	cells := make(Cells, 0, numberOfLiveCells)

	for i := 0; i < numberOfLiveCells; i++ {
		x, y := readCoordinates()

		if y >= m {
			log.Fatalf("coordinate (%v,%v) out of bound of %v by %v ", y, x, m, n)
		}

		if x >= n {
			log.Fatalf("coordinate (%v,%v) out of bound of %v by %v ", y, x, m, n)
		}

		cell := Cell{
			X:       x,
			Y:       y,
			IsAlive: true,
		}
		cells = append(cells, cell)
	}

	grid := Grid{
		AliveCells: cells,
		Rows:       m,
		Columns:    n,
	}

	nextGenGrid := grid

	for i := 0; i < numberOfGenerations; i++ {
		nextGenGrid = nextGenGrid.NextGeneration()
	}

	nextGenGrid.Print()
}
