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
	if generateEnable() {
		Gen()
		return
	}

	board := sudoku.LoadBoard(os.Stdin, N)
	result, err := sudoku.Solve(board, N, rand.Float64())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if checkUniqueEnable() {
		ok := sudoku.CheckUnique(board, result, N)
		if !ok {
			fmt.Fprintln(os.Stderr, "It is not unique.")
			os.Exit(1)
		}
	}

	PrintBoard(result)
}

func PrintBoard(board [][]int) {
	for _, row := range board {
		for _, v := range row {
			if v == sudoku.Unfilled {
				fmt.Print(".")
			} else {
				fmt.Print(v + 1)
			}
		}
		fmt.Println()
	}
}

func checkUniqueEnable() bool {
	return len(os.Args) > 1 && os.Args[1] == "check"
}

func generateEnable() bool {
	return len(os.Args) > 1 && os.Args[1] == "gen"
}

func Gen() {
	board, err := sudoku.Generate(N)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	PrintBoard(board)
}
