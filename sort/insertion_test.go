package sort_test

import (
	"cmp"
	"testing"

	"github.com/denpeshkov/algorithms/sort"
	. "github.com/denpeshkov/algorithms/sort/internal"
)

func TestInsertion(t *testing.T) {
	sortingFunc := func(x []int) {
		sort.InsertionCmp(x, cmp.Compare)
	}

	TestEmpty(sortingFunc)
	TestData(sortingFunc, t)
	TestRandom(1_000, sortingFunc, t)
	TestReverseSort(sortingFunc, t)
}

func FuzzInsertion(f *testing.F) {
	sortingFunc := func(s []byte) {
		sort.InsertionCmp(s, cmp.Compare)
	}

	FuzzSort(sortingFunc, f)
}
