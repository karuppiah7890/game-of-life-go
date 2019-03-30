package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextGeneration(t *testing.T) {
	t.Run("block pattern", func(t *testing.T) {
		cells := Cells{
			Cell{X: 1, Y: 1, IsAlive: true},
			Cell{X: 1, Y: 2, IsAlive: true},
			Cell{X: 2, Y: 1, IsAlive: true},
			Cell{X: 2, Y: 2, IsAlive: true},
		}
		grid := Grid{
			Rows:       4,
			Columns:    4,
			AliveCells: cells,
		}

		expectedNextGenAliveCells := cells

		nextGenGrid := grid.NextGeneration()

		actualNextGenAliveCells := nextGenGrid.AliveCells

		equal := cmp.Equal(expectedNextGenAliveCells, actualNextGenAliveCells)

		if !equal {
			t.Logf("expected: %v\ngot: %v\ndiff: %v",
				expectedNextGenAliveCells,
				actualNextGenAliveCells,
				cmp.Diff(expectedNextGenAliveCells, actualNextGenAliveCells))
			t.Fail()
		}
	})
}
