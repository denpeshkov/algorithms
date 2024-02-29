package sort_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/denpeshkov/algorithms/sort"
)

var selectionFuncData = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestSelectionFunc_EmptyNil(t *testing.T) {
	testSortFuncEmptyNil(t, sort.SelectionFunc[[]int], cmp.Compare[int])
}

func TestSelectionFunc_Data(t *testing.T) {
	testSortFuncData(t, sort.SelectionFunc[[]int], cmp.Compare[int], selectionFuncData)
}

func TestSelectionFunc_Reverse(t *testing.T) {
	testSortFuncReverse(t, sort.SelectionFunc[[]int], cmp.Compare[int], selectionFuncData)
}

func TestSelectionFunc_RandomInts(t *testing.T) {
	testSortFuncRandomInts(t, sort.SelectionFunc[[]int], cmp.Compare[int])
}

func BenchmarkSelectionFunc1K(b *testing.B) {
	benchmarkSortFunc1K(b, sort.SelectionFunc[[]int], cmp.Compare[int])
}

func FuzzSelectionFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		sort.SelectionFunc(s, cmp.Compare)

		if !slices.IsSortedFunc(s, cmp.Compare) {
			t.Errorf("slice was not sorted")
		}
	})
}
