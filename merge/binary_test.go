package merge_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/denpeshkov/algorithms/merge"
)

func TestBinaryCmp(t *testing.T) {
	tests := []struct {
		s []int
		t []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, []int{1}},
		{[]int{1, 2}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{2}},
		{[]int{3}, []int{1}},
		{[]int{3, 4}, []int{1, 2}},
		{[]int{1, 2}, []int{3, 4}},
		{[]int{1, 4, 10, 23}, []int{2, 2, 5, 8, 13, 22, 25, 30}},
		{[]int{4, 100}, []int{5, 5, 5, 5, 5}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{1, 1, 1, 1}, []int{2, 2, 2, 2}},
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}},
		{[]int{5, 6, 7, 8}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 2, 2, 3, 3}, []int{4, 4, 5, 5, 6, 6}},
		{[]int{2, 4, 6}, []int{1, 3, 5}},
		{[]int{}, []int{}},
		{[]int{1}, []int{}},
		{[]int{}, []int{1}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{}, []int{1, 2, 3}},
	}

	for _, e := range tests {
		r := merge.BinaryCmp(e.s, e.t, cmp.Compare)

		a := append(e.s, e.t...)
		slices.SortFunc(a, cmp.Compare)
		if !slices.Equal(r, a) {
			t.Errorf("BinaryCmp(%v, %v) = %v; want %v", e.s, e.t, r, a)
		}
	}
}

func FuzzBinaryCmp(f *testing.F) {
	f.Fuzz(func(t *testing.T, str1 string, str2 string) {
		s1 := []rune(str1)
		s2 := []rune(str2)

		r := merge.BinaryCmp(s1, s2, cmp.Compare)

		a := append(s1, s2...)
		slices.SortFunc(a, cmp.Compare)
		if !slices.Equal(r, a) {
			t.Errorf("BinaryCmp(%v, %v) = %v; want %v", s1, s2, r, a)
		}
	})
}
