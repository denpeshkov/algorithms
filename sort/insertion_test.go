package sort_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/denpeshkov/algorithms/sort"
)

var insertionFuncData = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestInsertionFunc_EmptyNil(t *testing.T) {
	testSortFuncEmptyNil(t, sort.InsertionFunc[[]int], cmp.Compare[int])
}

func TestInsertionFunc_Data(t *testing.T) {
	testSortFuncData(t, sort.InsertionFunc[[]int], cmp.Compare[int], insertionFuncData)
}

func TestInsertionFunc_Reverse(t *testing.T) {
	testSortFuncReverse(t, sort.InsertionFunc[[]int], cmp.Compare[int], insertionFuncData)
}

func TestInsertionFunc_RandomInts(t *testing.T) {
	testSortFuncRandomInts(t, sort.InsertionFunc[[]int], cmp.Compare[int])
}

func TestInsertionFunc_Stability(t *testing.T) {
	n, m := 1000, 100

	testSortFuncStability(t, sort.InsertionFunc[intPairs], n, m)
}

func BenchmarkInsertionFunc1K(b *testing.B) {
	benchmarkSortFunc1K(b, sort.InsertionFunc[[]int], cmp.Compare[int])
}

func FuzzInsertionSortFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		sort.InsertionFunc(s, cmp.Compare)

		if !slices.IsSortedFunc(s, cmp.Compare) {
			t.Errorf("slice was not sorted")
		}
	})
}
