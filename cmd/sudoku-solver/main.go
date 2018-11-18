package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	sudoku "github.com/pocke/sudoku-solver"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	board := sudoku.LoadBoard(os.Stdin)
	result, err := sudoku.Solve(board)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	for _, row := range result {
		for _, v := range row {
			fmt.Print(v + 1)
		}
		fmt.Println()
	}
}
