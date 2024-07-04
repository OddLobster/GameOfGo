package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintGrid(grid [][]int) {
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

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// game of life server client architecture
// client reqeusts initial state
// and client pushes state to server
// server calculates the new state and returns it to the client
// super stupid but seems fun
