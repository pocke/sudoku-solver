package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
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

	out := &bytes.Buffer{}
	Solve(strings.NewReader(problem), out)
	if out.String() != expected {
		t.Error("Invalid sudoku")
	}
}

func TestCombination(t *testing.T) {
	assert := func(expected, got [][]int) {
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected: %v, but got %v", expected, got)
		}
	}

	got := Combination(2)
	assert([][]int{
		[]int{0, 1},
	}, got)

	got = Combination(4)
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
		Combination(N)
	}
}
