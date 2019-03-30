package main

// Cell represents a cell in the game of life
type Cell struct {
	X       int
	Y       int
	IsAlive bool
}

// Cells represents a list of cells
type Cells []Cell

// NumberOfAliveNeighbors gives the number of alive neighbor cells for a given cell in a given grid
func (cell Cell) NumberOfAliveNeighbors(grid Grid) int {
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

// DoILive tells if a given cell lives in the next generation, given the number of alive neighbor cells
func (cell Cell) DoILive(numberOfAliveNeighbors int) bool {
	if cell.IsAlive {
		if numberOfAliveNeighbors < 2 || numberOfAliveNeighbors > 3 {
			return false
		}

		return true
	}

	if numberOfAliveNeighbors == 3 {
		return true
	}

	return false
}
