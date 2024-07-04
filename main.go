package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func countNeighbours(grid [][]int, l int, m int) int {
	countAliveNeighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// check boundaries
			if (i+l >= 0 && i+l < len(grid)) && (m+j >= 0 && m+j < len(grid[0])) {
				countAliveNeighbours += grid[i+l][m+j]
			}
		}
	}
	countAliveNeighbours -= grid[l][m]
	return countAliveNeighbours
}

func updateGridState(grid [][]int) [][]int {
	var newGrid [][]int = make([][]int, len(grid))
	for i, _ := range grid {
		newGrid[i] = make([]int, len(grid[0]))
	}

	for i, row := range grid {
		for j, _ := range row {
			aliveNeighbours := countNeighbours(grid, i, j)
			// lonely cell dies
			if grid[i][j] == 1 && aliveNeighbours < 2 {
				newGrid[i][j] = 0
				// overpopulated cell dies
			} else if grid[i][j] == 1 && aliveNeighbours > 3 {
				newGrid[i][j] = 0
				// happy cell is born
			} else if grid[i][j] == 0 && aliveNeighbours == 3 {
				newGrid[i][j] = 1
			} else {
				newGrid[i][j] = grid[i][j]
			}
		}
	}
	return newGrid
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func initGrid(rows int, cols int, zeroPercent float64) [][]int {
	var grid [][]int = make([][]int, rows)
	for i, _ := range grid {
		grid[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rand.Float64() < zeroPercent {
				grid[i][j] = 0
			} else {
				grid[i][j] = 1
			}
		}
	}
	return grid
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Println("Start Init Game State")

	const rows int = 10
	const cols int = 20
	grid := initGrid(rows, cols, 0.5)
	fmt.Println(grid)
	numIterations := 100
	for i := 0; i < numIterations; i++ {
		clearConsole()
		printGrid(grid)
		grid = updateGridState(grid)
		time.Sleep(500 * time.Millisecond)
	}

}

// game of life server client architecture
// client reqeusts initial state
// and client pushes state to server
// server calculates the new state and returns it to the client
// super stupid but seems fun
