package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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

func createJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
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

func generateInitialGrid(w http.ResponseWriter, r *http.Request) {
	rows, _ := strconv.Atoi(r.URL.Query().Get("rows"))
	cols, _ := strconv.Atoi(r.URL.Query().Get("cols"))
	grid := initGrid(rows, cols, 0.5)
	createJsonResponse(w, grid)
}

func processGrid(w http.ResponseWriter, r *http.Request) {
	var grid [][]int
	err := json.NewDecoder(r.Body).Decode(&grid)
	if err != nil {
		fmt.Println("Couldnt decode request", err)
	}

	updatedGrid := updateGridState(grid)
	createJsonResponse(w, updatedGrid)
}

func main() {
	http.HandleFunc("/requestGrid", generateInitialGrid)
	http.HandleFunc("/processGrid", processGrid)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to listen to 8080", err)
	}
}
