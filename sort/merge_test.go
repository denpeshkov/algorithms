package sort_test

import (
	"cmp"
	"testing"

	"github.com/denpeshkov/algorithms/sort"
	. "github.com/denpeshkov/algorithms/sort/internal"
)

func TestMerge(t *testing.T) {
	sortingFunc := func(x []int) {
		sort.MergeCmp(x, cmp.Compare)
	}

	TestEmpty(sortingFunc)
	TestData(sortingFunc, t)
	TestRandom(1_000, sortingFunc, t)
	TestReverseSort(sortingFunc, t)
}

func FuzzMerge(f *testing.F) {
	sortingFunc := func(s []byte) {
		sort.MergeCmp(s, cmp.Compare)
	}

	FuzzSort(sortingFunc, f)
}
