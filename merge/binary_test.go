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

		if a := append(e.s, e.t...); !containsAllElements(a, r) {
			t.Error("merged slice contains not all of the elements")
		}

		if !slices.IsSortedFunc(r, cmp.Compare) {
			t.Errorf("merged slice is not sorted")
		}
	}
}

func FuzzBinaryCmp(f *testing.F) {
	f.Fuzz(func(t *testing.T, str1 string, str2 string) {
		s1 := []rune(str1)
		s2 := []rune(str2)

		r := merge.BinaryCmp(s1, s2, cmp.Compare)

		if a := append(s1, s2...); !containsAllElements(a, r) {
			t.Error("merged slice contains not all of the elements")
		}

		if !slices.IsSortedFunc(r, cmp.Compare) {
			t.Errorf("merged slice is not sorted")
		}
	})
}

// Returns true if t contains all of the elements of s; otherwise returns false.
// Meaning t is a permutation of s.
func containsAllElements[T cmp.Ordered](s, t []T) bool {
	slices.Sort(s)
	slices.Sort(t)
	return slices.Equal(s, t)
}
