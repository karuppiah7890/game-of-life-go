package main

import "fmt"

// Grid represents the grid of game of life
type Grid struct {
	AliveCells Cells
	Rows       int
	Columns    int
}

// NextGeneration gives the next generation given a grid of game of life
func (grid Grid) NextGeneration() Grid {
	aliveCells := grid.AliveCells
	nextGenAliveCells := make(Cells, 0, len(aliveCells))

	aliveCellsStayingAlive := grid.NextGenAliveCellsOf(aliveCells)

	nextGenAliveCells = append(nextGenAliveCells, aliveCellsStayingAlive...)

	deadCells := grid.DeadCells()
	deadCellsTurningAlive := grid.NextGenAliveCellsOf(deadCells)

	nextGenAliveCells = append(nextGenAliveCells, deadCellsTurningAlive...)

	return Grid{
		Rows:       grid.Rows,
		Columns:    grid.Columns,
		AliveCells: nextGenAliveCells,
	}
}

// IsCellAlive tells if a cell is alive in a grid given the coordinates
func (grid Grid) IsCellAlive(x, y int) bool {
	for _, aliveCell := range grid.AliveCells {
		if aliveCell.X == x && aliveCell.Y == y {
			return true
		}
	}

	return false
}

// DeadCells gives the list of dead cells in a grid
func (grid Grid) DeadCells() Cells {
	deadCells := make(Cells, 0, grid.Rows*grid.Columns)
	for x := 0; x < grid.Columns; x++ {
		for y := 0; y < grid.Rows; y++ {
			if !grid.IsCellAlive(x, y) {

				deadCell := Cell{
					X:       x,
					Y:       y,
					IsAlive: false,
				}

				deadCells = append(deadCells, deadCell)
			}
		}
	}

	return deadCells
}

// NextGenAliveCellsOf gives the next generation alive cells of the given cells
func (grid Grid) NextGenAliveCellsOf(cells Cells) Cells {
	nextGenAliveCells := make(Cells, 0, len(cells))
	for _, cell := range cells {
		numberOfAliveneighborCells := grid.NumberOfAliveNeighborsOf(cell)

		cellLives := cell.DoILive(numberOfAliveneighborCells)

		if cellLives {
			nextGenCell := Cell{
				X:       cell.X,
				Y:       cell.Y,
				IsAlive: true,
			}
			nextGenAliveCells = append(nextGenAliveCells, nextGenCell)
		}
	}

	return nextGenAliveCells
}

// NumberOfAliveNeighborsOf gives the number of alive neighbor cells for a given cell in a given grid
func (grid Grid) NumberOfAliveNeighborsOf(cell Cell) int {
	numberOfAliveNeighbors := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if grid.IsCellAlive(cell.X+i, cell.Y+j) {
				numberOfAliveNeighbors++
			}
		}
	}

	return numberOfAliveNeighbors
}

// Print prints the grid
func (grid Grid) Print() {
	for y := 0; y < grid.Rows; y++ {
		for x := 0; x < grid.Columns; x++ {
			if grid.IsCellAlive(x, y) {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
