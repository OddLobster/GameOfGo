package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"GameOfGo/utils"
)

func requestInitialGrid() [][]int {
	response, err := http.Get("http://localhost:8080/requestGrid")
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
	numIterations := 100
	grid := requestInitialGrid()
	utils.PrintGrid(grid)
	for i := 0; i < numIterations; i++ {
		utils.ClearConsole()
		grid = getUpdatedGrid(grid)
		utils.PrintGrid(grid)
		time.Sleep(500 * time.Millisecond)
	}
}
