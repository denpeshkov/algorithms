package sort_test

import (
	"cmp"
	"testing"

	"github.com/denpeshkov/algorithms/sort"
	. "github.com/denpeshkov/algorithms/sort/internal"
)

func TestSelection(t *testing.T) {
	sortingFunc := func(x []int) {
		sort.SelectionCmp(x, cmp.Compare)
	}

	TestEmpty(sortingFunc)
	TestData(sortingFunc, t)
	TestRandom(1_000, sortingFunc, t)
	TestReverseSort(sortingFunc, t)
}

func FuzzSelection(f *testing.F) {
	sortingFunc := func(r []byte) {
		sort.SelectionCmp(r, cmp.Compare)
	}

	FuzzSort(sortingFunc, f)
}
