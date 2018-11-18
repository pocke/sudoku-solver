package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	sudoku "github.com/pocke/sudoku-solver"
)

const N = 9

func main() {
	rand.Seed(time.Now().UnixNano())
	board := sudoku.LoadBoard(os.Stdin, N)
	result, err := sudoku.Solve(board, N, rand.Float64())
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
