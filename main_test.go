package main

import (
	"reflect"
	"testing"
)

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
