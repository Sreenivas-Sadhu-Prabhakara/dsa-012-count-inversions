package main

import (
	"reflect"
	"testing"
)

func TestCountInversions(t *testing.T) {
	cases := []struct {
		in   []int
		want int64
	}{
		{[]int{5, 3, 2, 4, 1}, 8},
		{[]int{2, 4, 1, 3, 5}, 3},
		{[]int{1, 2, 3}, 0},
		{[]int{5, 4, 3, 2, 1}, 10},
		{[]int{1, 1, 1}, 0},
		{[]int{}, 0},
		{[]int{7}, 0},
	}
	for _, c := range cases {
		if got := CountInversions(c.in); got != c.want {
			t.Errorf("CountInversions(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestInputUntouched(t *testing.T) {
	in := []int{3, 1, 2}
	CountInversions(in)
	if !reflect.DeepEqual(in, []int{3, 1, 2}) {
		t.Errorf("input was modified: %v", in)
	}
}
