package sudoku

import (
	"bufio"
	"io"
	"math"
	"strconv"

	"github.com/pkg/errors"
	"github.com/pocke/go-minisat"
)

const (
	Unfilled = -1
)

func newSatSolver(N int, seed float64) (*minisat.Solver, [][][]*minisat.Var) {
	s := minisat.NewSolver(seed)
	vars := make([][][]*minisat.Var, 0, N)
	for i := 0; i < N; i++ {
		vars = append(vars, make([][]*minisat.Var, 0, N))
		for j := 0; j < N; j++ {
			vars[i] = append(vars[i], make([]*minisat.Var, 0, N))
			for k := 0; k < N; k++ {
				vars[i][j] = append(vars[i][j], s.NewVar())
			}
		}
	}

	for _, line := range vars {
		for _, cell := range line {
			// 各マスは1..Nのいずれかの数字が入る
			s.AddClause(cell...)
		}
	}

	comb := Combination(N)

	// 各列に対して同じ数字が2回以上現れない
	for _, line := range vars {
		for num := 0; num < N; num++ {
			for _, c := range comb {
				s.AddClause(line[c[0]][num].Not(), line[c[1]][num].Not())
			}
		}
	}

	// 各行に対して同じ数字が2回以上現れない
	for i := 0; i < N; i++ {
		for num := 0; num < N; num++ {
			for _, c := range comb {
				s.AddClause(vars[c[0]][i][num].Not(), vars[c[1]][i][num].Not())
			}
		}
	}

	// 各ブロックに対して同じ数字が2回以上現れない
	m := int(math.Sqrt(float64(N)))
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for _, c := range comb {
				x1 := c[0] % m
				x2 := c[1] % m
				y1 := c[0] / m
				y2 := c[1] / m

				for num := 0; num < N; num++ {
					s.AddClause(vars[i*m+y1][j*m+x1][num].Not(), vars[i*m+y2][j*m+x2][num].Not())
				}
			}
		}
	}
	return s, vars
}

func Solve(board [][]int, N int, seed float64) ([][]int, error) {
	s, vars := newSatSolver(N, seed)

	for i, row := range board {
		for j, v := range row {
			if v != Unfilled {
				s.AddClause(vars[i][j][v])
			}
		}
	}

	res := s.Solve()
	if !res {
		return nil, errors.New("Unsatisfy it.")
	}

	result := make([][]int, N)
	for i := 0; i < N; i++ {
		result[i] = make([]int, N)
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				v := vars[i][j][k]
				b, _ := s.ModelValue(v)
				if b {
					result[i][j] = k
					break
				}
			}
		}
	}

	return result, nil
}

func CheckUnique(board, result [][]int, N int) bool {
	s, vars := newSatSolver(N, 1)
	for i, row := range board {
		for j, v := range row {
			if v != Unfilled {
				s.AddClause(vars[i][j][v])
			}
		}
	}

	uniqChecker := make([]*minisat.Var, 0, N*N)
	for i, row := range result {
		for j, v := range row {
			uniqChecker = append(uniqChecker, vars[i][j][v].Not())
		}
	}
	s.AddClause(uniqChecker...)

	return !s.Solve()
}

func LoadBoard(in io.Reader, N int) [][]int {
	res := make([][]int, N)
	for i := 0; i < N; i++ {
		res[i] = make([]int, N)
	}

	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanBytes)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sc.Scan()
			t := sc.Text()
			n, err := strconv.Atoi(t)
			if err != nil || n == 0 {
				res[i][j] = Unfilled
			} else {
				res[i][j] = n - 1
			}
		}
		sc.Scan() // \n
	}

	return res
}

func Combination(n int) [][]int {
	var res = make([][]int, 0, n*(n-1)/2)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			res = append(res, []int{i, j})
		}
	}

	return res
}
