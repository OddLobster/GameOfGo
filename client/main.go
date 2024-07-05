package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"GameOfGo/utils"
)

func requestInitialGrid(rows int, cols int) [][]int {
	url := fmt.Sprintf("http://localhost:8080/requestGrid?rows=%d&cols=%d", rows, cols)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error requesting initial grid:", err)
	}
	defer response.Body.Close()

	var grid [][]int
	err = json.NewDecoder(response.Body).Decode(&grid)
	if err != nil {
		fmt.Println("Error decoding response:", err)
	}

	return grid
}

func getUpdatedGrid(grid [][]int) [][]int {
	jsonData, _ := json.Marshal(grid)
	response, _ := http.Post("http://localhost:8080/processGrid", "application/json", bytes.NewBuffer(jsonData))
	defer response.Body.Close()

	var processedGrid [][]int
	json.NewDecoder(response.Body).Decode(&processedGrid)
	return processedGrid
}

func main() {
	numIterations := flag.Int("numIter", 100, "Number of iterations steps for the GOL simulation")
	gridRows := flag.Int("rows", 25, "Number of rows in the grid")
	gridCols := flag.Int("cols", 25, "Number of columns in the grid")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *help {
		flag.Usage()
	}

	grid := requestInitialGrid(*gridRows, *gridCols)
	utils.PrintGrid(grid)
	for i := 0; i < *numIterations; i++ {
		utils.ClearConsole()
		grid = getUpdatedGrid(grid)
		utils.PrintGrid(grid)
		time.Sleep(250 * time.Millisecond)
	}
}
