package main

// Cell represents a cell in the game of life
type Cell struct {
	X       int
	Y       int
	IsAlive bool
}

// Cells represents a list of cells
type Cells []Cell

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
