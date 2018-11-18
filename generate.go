package sudoku

import "math/rand"

func Generate(N int) ([][]int, error) {
	emptyBoard := make([][]int, N)
	for i := 0; i < N; i++ {
		emptyBoard[i] = make([]int, N)
		for j := 0; j < N; j++ {
			emptyBoard[i][j] = Unfilled
		}
	}

	seed := rand.Float64()
	solution, err := Solve(emptyBoard, N, seed)
	if err != nil {
		return nil, err
	}

	for {
		result := tryGen(solution)
		if result != nil {
			return result, nil
		}
	}
}

// TODO: Fix magic numbers for N != 9
func tryGen(solution [][]int) [][]int {
	N := len(solution)
	board := copySolution(solution)
	filledCount := int(rand.NormFloat64()*5 + 25)
	if filledCount > 35 {
		filledCount = 35
	} else if filledCount < 17 {
		filledCount = 17
	}

	removeRemaining := N*N - filledCount
	for removeRemaining != 0 {
		x := rand.Intn(N)
		y := rand.Intn(N)
		if board[x][y] != Unfilled {
			board[x][y] = Unfilled
			removeRemaining--
		}
	}

	uniq := CheckUnique(board, solution, N)
	if uniq {
		return board
	}
	return nil
}

func copySolution(sol [][]int) [][]int {
	res := make([][]int, len(sol))
	for i := 0; i < len(sol); i++ {
		res[i] = make([]int, len(sol[i]))
		copy(res[i], sol[i])
	}
	return res
}
