package sudoku_test

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"

	sudoku "github.com/pocke/sudoku-solver"
)

func TestSolve(t *testing.T) {
	problem := `..4..5..1
.......72
.9.37..5.
.3..6.249
..64.95..
945.1..3.
.1..46.9.
48.......
5..9..4..
`
	expected := `374625981
658194372
192378654
831567249
726439518
945812736
213746895
489253167
567981423
`

	board := sudoku.LoadBoard(strings.NewReader(problem))
	result, err := sudoku.Solve(board)
	if err != nil {
		t.Fatal(err)
	}
	out := &bytes.Buffer{}
	for _, row := range result {
		for _, v := range row {
			s := strconv.Itoa(v + 1)
			out.WriteString(s)
		}
		out.WriteString("\n")
	}
	if out.String() != expected {
		t.Errorf("Invalid sudoku, expected: %s got: %s", expected, out.String())
	}
}

func TestCombination(t *testing.T) {
	assert := func(expected, got [][]int) {
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, but got %v", expected, got)
		}
	}

	got := sudoku.Combination(2)
	assert([][]int{
		[]int{0, 1},
	}, got)

	got = sudoku.Combination(4)
	assert([][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}, got)
}

func BenchmarkCombination(b *testing.B) {
	N := 9
	for i := 0; i < b.N; i++ {
		sudoku.Combination(N)
	}
}
